<template>
  <v-list-item-avatar
    tile
    size="80"
    color="grey"
  >
    <v-img :src="userAvatar" />
  </v-list-item-avatar>
</template>

<script lang='ts'>
import Vue, { PropType } from 'vue'
import { IUser } from '../../types/user'

export default Vue.extend({
  props: {
    user: {
      type: Object as PropType<IUser>,
      required: true,
    }
  },
  computed: {
    userAvatar (): string {
      if (!this.user.avatar) return '/images/flutter.png'

      return this.user.avatar
    }
  }
})
</script>
