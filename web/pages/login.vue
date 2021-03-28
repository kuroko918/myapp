<template>
  <v-container
    class="d-flex flex-column justify-center align-center"
  >
    <v-btn
      v-if='!isAuthenticated()'
      large
      :loading="loading"
      :disabled="loading"
      @click="login"
    >
      <v-icon
        class="mr-2"
      >
        mdi-github
      </v-icon>
      GitHubでログイン
    </v-btn>
  </v-container>
</template>

<script lang='ts'>
import Vue from 'vue'

export default Vue.extend({
  data () {
    return {
      loading: false,
    }
  },
  computed: {
    isAuthenticated(): boolean {
      return this.$store.getters['auth/isAuthenticated'];
    }
  },
  methods: {
    async login () {
      this.loading = true
      await this.$store.dispatch('auth/login')
      this.loading = false
      this.$router.push({ name: 'chat' })
    }
  }
})
</script>

<style scoped>
.container {
  min-height: 100%;
}
</style>
