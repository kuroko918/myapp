<template>
  <div class="container">
    <button class="button is-primary is-rounded" @click="login">
      GitHubでログイン
    </button>
  </div>
</template>

<script lang='ts'>
import Vue from 'vue'

export default Vue.extend({
  methods: {
    async login (): Promise<void> {
      await this.$store.dispatch('auth/login')
      this.$router.push({ name: 'index' })
    }
  }
})
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>
