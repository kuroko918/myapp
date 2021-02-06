<template>
  <div>
    <img class="workerIcon" :src="userAvatar">
    <div>{{ user.name }}</div>
  </div>
</template>

<script lang='ts'>
import Vue, { PropType } from 'vue'
import { IUser } from '../../types/models/user'

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

<style>
.workerIcon {
  width: 25px;
  height: 25px;
}
</style>
