<template>
  <v-row justify="center" align="center">
    <v-btn
      block
      outlined
      color="grey darken-3"
      @click="logout"
    >
      ログアウト
    </v-btn>
  </v-row>
</template>

<script lang='ts'>
import Vue from 'vue'

export default Vue.extend({
  methods: {
    async logout (): Promise<void> {
      await this.$store.dispatch('auth/logout')
      this.$router.push({ name: 'login' })
    }
  }
})
</script>
