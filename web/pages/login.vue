<template>
  <button class="button is-primary is-rounded" @click="login">
    GitHubでログイン
  </button>
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
