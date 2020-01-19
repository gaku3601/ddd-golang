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
    <HelloWorldComponent msg="Welcome to Your Vue.js App" />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import HelloWorldComponent from "@/components/HelloWorld.vue";

@Component({
  components: {
    HelloWorldComponent
  }
})
export default class HelloWorld extends Vue {
  @Prop() private msg!: string;

  id: string = "";
  password: string = "";

  login(): void {
    alert("id:" + this.id + " password:" + this.password);
  }
}
</script>
