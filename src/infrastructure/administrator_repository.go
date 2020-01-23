package infrastructure

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gaku3601/ddd-golang/src/domain/model"
)

type AdministratorRepository struct{}

func (u *AdministratorRepository) Save(ctx context.Context, administrator *model.Administrator) error {
	// cognitoとregionとuserPoolIDを環境変数に格納しておく
	region := os.Getenv("COGNITO_REGION")
	userPoolID := os.Getenv("COGNITO_USER_POOL_ID")
	// IAMロールでユーザを設定し、以下を設定する
	// export AWS_ACCESS_KEY_ID=your access key id
	// export AWS_SECRET_ACCESS_KEY=your secret key
	creds := credentials.NewEnvCredentials()
	conf := &aws.Config{Region: aws.String(region), Credentials: creds}
	sess, err := session.NewSession(conf)
	if err != nil {
		panic(err)
	}

	user := &cognito.AdminCreateUserInput{
		UserPoolId: aws.String(userPoolID),
		Username:   aws.String(administrator.AdministratorEmail().Value()),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("custom:role"),
				Value: aws.String("admin"),
			},
			{
				Name:  aws.String("nickname"),
				Value: aws.String("gaku"),
			},
		},
	}
	_, err = cognito.New(sess).AdminCreateUser(user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
