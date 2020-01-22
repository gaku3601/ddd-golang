<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <div>
      <div>
        <input type="text" v-model="id" />
      </div>
      <div>
        <input type="password" v-model="password" />
      </div>
      <div>
        <button @click="login">login</button>
      </div>
    </div>
    <router-link :to="{name: 'password-reset'}">パスワードリセット</router-link>
    <HelloWorldComponent msg="Welcome to Your Vue.js App" />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import HelloWorldComponent from "@/components/HelloWorld.vue";
import {
  CognitoUser,
  AuthenticationDetails,
  CognitoUserPool
} from 'amazon-cognito-identity-js'

@Component({
  components: {
    HelloWorldComponent
  }
})
export default class Home extends Vue {
  @Prop() private msg!: string;

  id: string = "";
  password: string = "";

  async login(): Promise<void> {
    const config = {
      Region: "ap-northeast-1",
      UserPoolId: "ap-northeast-1_BHmB1pj9e",
      ClientId: "7bmm4mkctc8b6nhbs2kpgs62bt",
      IdentityPoolId: "arn:aws:cognito-idp:ap-northeast-1:982976011432:userpool/ap-northeast-1_BHmB1pj9e"
    };
    const userData = { Username: this.id, Pool: new CognitoUserPool({
      UserPoolId: config.UserPoolId,
      ClientId: config.ClientId
    })};
    const cognitoUser = new CognitoUser(userData);
    const authenticationData = { Username: this.id, Password: this.password };
    const authenticationDetails = new AuthenticationDetails(authenticationData);
    await  cognitoUser.authenticateUser(authenticationDetails, {
        onSuccess: (result) => {
          // 実際にはクレデンシャルなどをここで取得する(今回は省略)
          console.log("success");
          console.log(result);
          console.log(result.getIdToken().getJwtToken());
          localStorage.setItem("token", result.getIdToken().getJwtToken());
          this.$router.push({name: 'dashboard'});
        },
        onFailure: (err) => {
          console.log("error");
          console.log(err);
        },
        newPasswordRequired(userAttributes, requiredAttributes) {
          //Force Change Password状態で、認証に成功した場合に呼ばれる。
          console.log("newPasswordRequired()");
          console.log(userAttributes, requiredAttributes);
          cognitoUser.completeNewPasswordChallenge(authenticationData.Password, {}, {
            onSuccess: (result) => {
              console.log("success");
              console.log(result);
            },
            onFailure: (err) => {
              console.log("error");
              console.log(err);
            },
          });
        },
      })
  }
}
</script>
