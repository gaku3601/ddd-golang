# cognitoへのinitユーザの登録
cognito - init-importUser.csvを作成したcognitoへimportすればok(メールアドレスは適宜変える)

# credential設定
IAMロールでユーザを設定し、以下を設定する  
export AWS_ACCESS_KEY_ID=your access key id  
export AWS_SECRET_ACCESS_KEY=your secret key

# cognitoの環境変数設定
export COGNITO_REGION=cognito region  
export COGNITO_USER_POOL_ID=cognito user pool id
