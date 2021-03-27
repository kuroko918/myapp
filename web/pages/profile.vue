<template>
  <div>
  </div>
</template>

<script lang='ts'>
import Vue from 'vue'

export default Vue.extend({
  created () {
    this.getUser()
  },
  methods: {
    async getUser () {
      await this.$store.dispatch('user/getUser')
    },
  }
})
</script>
