<template>
  <div>
    <div>
      <p>ユーザを作成</p>
      <div>
        <input type="text" v-model="email" />
      </div>
      <div>
        <button @click="userCreate">ユーザを作成</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from 'vue-property-decorator'
import axios from '@/util/axios'

@Component({
  components: {}
})
export default class Dashboard extends Vue {
  email: string = '';

  async userCreate(): Promise<void> {
    console.log('create');
    await axios.post('/users', {
      email: this.email,
    });
  }
}
</script>

<style scoped lang="scss">
</style>