package middleware

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// https://github.com/yoskeoka/gognito
// cognito user poolまたはgoogle+の認証したJWTのチェックを行う
func AuthMiddleware() (gin.HandlerFunc, error) {
	// この部分はサーバー起動時に1度だけ実行される
	region := os.Getenv("COGNITO_REGION")
	userPoolID := os.Getenv("COGNITO_USER_POOL_ID")

	// 1. Download and store the JSON Web Key (JWK) for your user pool.
	jwkURL := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v/.well-known/jwks.json", region, userPoolID)
	fmt.Println(jwkURL)
	jwk, err := getJWK(jwkURL)
	if err != nil {
		return nil, errors.New("jwk error")
	}

	// googleのJWKも一括で管理して混ぜてしまう。どうせkidは被らない
	googleJwk, err := getJWK("https://www.googleapis.com/oauth2/v3/certs")
	if err != nil {
		return nil, errors.New("jwk error")
	}
	for key, val := range googleJwk {
		jwk[key] = val
	}

	fmt.Println(jwk)

	// この部分はクライアント接続毎に実行される
	return func(c *gin.Context) {
		tokenString, ok := getBearer(c.Request.Header["Authorization"])

		if !ok {
			// jwtがHeaderに添付されていない
			c.AbortWithStatusJSON(401, gin.H{"text": "Authorization Bearer Header is missing"})
			return
		}

		token, err := validateToken(tokenString, region, userPoolID, jwk)
		if err != nil || !token.Valid {
			// jwtの検証に失敗
			fmt.Printf("token is not valid\n%v", err)
			c.AbortWithStatusJSON(401, gin.H{"text": fmt.Sprintf("token is not valid%v", err)})
		} else {
			// 認証したtokenを渡してやると、そこに含まれるユーザー情報を各リソースパスで利用できる
			c.Set("token", token)
			c.Next()
		}
	}, nil
}

// Cognito User PoolのJWT検証は下記の手順を参照
//https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-identity-user-pools-using-id-and-access-tokens-in-web-api

// Google のJWT検証は下記の手順を参照
// https://developers.google.com/identity/sign-in/web/backend-auth

// validateAWSJwtClaims validates AWS Cognito User Pool JWT
func validateAWSJwtClaims(claims jwt.MapClaims, region, userPoolID string) error {
	var err error
	// 3. Check the iss claim. It should match your user pool.
	issShoudBe := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v", region, userPoolID)
	err = validateClaimItem("iss", []string{issShoudBe}, claims)
	if err != nil {
		return err
	}

	// 4. Check the token_use claim.
	validateTokenUse := func() error {
		if tokenUse, ok := claims["token_use"]; ok {
			if tokenUseStr, ok := tokenUse.(string); ok {
				if tokenUseStr == "id" || tokenUseStr == "access" {
					return nil
				}
			}
		}
		return errors.New("token_use should be id or access")
	}

	err = validateTokenUse()
	if err != nil {
		return err
	}

	// 7. Check the exp claim and make sure the token is not expired.
	err = validateExpired(claims)
	if err != nil {
		return err
	}

	return nil
}

// validateAWSJwtClaims validates Google JWT
func validateGoogleJwtClaims(claims jwt.MapClaims) error {
	var err error
	issShoudBe := []string{"accounts.google.com", "https://accounts.google.com"}
	err = validateClaimItem("iss", issShoudBe, claims)
	if err != nil {
		return err
	}
	aud := []string{os.Getenv("GOOGLE_CLIENT_ID")}
	err = validateClaimItem("aud", aud, claims)
	if err != nil {
		return err
	}
	err = validateExpired(claims)
	if err != nil {
		return err
	}

	return nil
}

func validateClaimItem(key string, keyShouldBe []string, claims jwt.MapClaims) error {
	if val, ok := claims[key]; ok {
		if valStr, ok := val.(string); ok {
			for _, shouldbe := range keyShouldBe {
				if valStr == shouldbe {
					return nil
				}
			}
		}
	}
	return fmt.Errorf("%v does not match any of valid values: %v", key, keyShouldBe)
}

func validateExpired(claims jwt.MapClaims) error {
	if tokenExp, ok := claims["exp"]; ok {
		if exp, ok := tokenExp.(float64); ok {
			now := time.Now().Unix()
			fmt.Printf("current unixtime : %v\n", now)
			fmt.Printf("expire unixtime  : %v\n", int64(exp))
			if int64(exp) > now {
				return nil
			}
		}
		return errors.New("cannot parse token exp")
	}
	return errors.New("token is expired")
}

func validateToken(tokenStr, region, userPoolID string, jwk map[string]JWKKey) (*jwt.Token, error) {

	// 2. Decode the token string into JWT format.
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		// cognito user pool, googleは : RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// 5. Get the kid from the JWT token header and retrieve the corresponding JSON Web Key that was stored
		if kid, ok := token.Header["kid"]; ok {
			if kidStr, ok := kid.(string); ok {
				key := jwk[kidStr]
				// 6. Verify the signature of the decoded JWT token.
				rsaPublicKey := convertKey(key.E, key.N)
				return rsaPublicKey, nil
			}
		}

		// rsa public key取得できず
		return "", nil
	})

	if err != nil {
		return token, err
	}

	claims := token.Claims.(jwt.MapClaims)

	iss, ok := claims["iss"]
	if !ok {
		return token, fmt.Errorf("token does not contain issuer")
	}
	issStr := iss.(string)
	if strings.Contains(issStr, "cognito-idp") {
		// 3. 4. 7.のチェックをまとめて
		err = validateAWSJwtClaims(claims, region, userPoolID)
		if err != nil {
			return token, err
		}
	} else if strings.Contains(issStr, "accounts.google.com") {
		err = validateGoogleJwtClaims(claims)
		if err != nil {
			return token, err
		}
	}

	if token.Valid {
		return token, nil
	}
	return token, err
}

func getBearer(auth []string) (jwt string, ok bool) {
	for _, v := range auth {
		ret := strings.Split(v, " ")
		if len(ret) > 1 && ret[0] == "Bearer" {
			return ret[1], true
		}
	}
	return "", false
}

// JWK is json data struct for JSON Web Key
type JWK struct {
	Keys []JWKKey
}

// JWKKey is json data struct for cognito jwk key
type JWKKey struct {
	Alg string
	E   string
	Kid string
	Kty string
	N   string
	Use string
}

func getJSON(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getJWK(jwkURL string) (map[string]JWKKey, error) {

	jwk := &JWK{}

	if err := getJSON(jwkURL, jwk); err != nil {
		return nil, err
	}

	jwkMap := make(map[string]JWKKey)
	for _, jwk := range jwk.Keys {
		jwkMap[jwk.Kid] = jwk
	}
	return jwkMap, nil
}

// https://gist.github.com/MathieuMailhos/361f24316d2de29e8d41e808e0071b13
func convertKey(rawE, rawN string) *rsa.PublicKey {
	decodedE, err := base64.RawURLEncoding.DecodeString(rawE)
	if err != nil {
		panic(err)
	}
	if len(decodedE) < 4 {
		ndata := make([]byte, 4)
		copy(ndata[4-len(decodedE):], decodedE)
		decodedE = ndata
	}
	pubKey := &rsa.PublicKey{
		N: &big.Int{},
		E: int(binary.BigEndian.Uint32(decodedE[:])),
	}
	decodedN, err := base64.RawURLEncoding.DecodeString(rawN)
	if err != nil {
		panic(err)
	}
	pubKey.N.SetBytes(decodedN)
	// fmt.Println(decodedN)
	// fmt.Println(decodedE)
	// fmt.Printf("%#v\n", *pubKey)
	return pubKey
}
