package infrastructure

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gaku3601/ddd-golang/src/domain/model"
)

type AdministratorRepository struct{}

func (u *AdministratorRepository) Save(ctx context.Context, administrator *model.Administrator) error {
	conf := &aws.Config{Region: aws.String("ap-northeast-1")}
	sess, err := session.NewSession(conf)
	if err != nil {
		panic(err)
	}

	user := &cognito.AdminCreateUserInput{
		UserPoolId: aws.String("ap-northeast-1_BHmB1pj9e"),
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
