<template>
  <div>
    <v-chip-group
      v-for="(message) in this.$store.state.message.messages"
      :key="message.id"
    >
      <MessageRow
        :chip-position="() => chipPosition(message)"
        :message="message"
      />
    </v-chip-group>
  </div>
</template>

<script lang='ts'>
import Vue from 'vue'
import { IMessage } from '../../types/models/message'
import MessageRow from './MessageRow.vue'

export default Vue.extend({
  components: {
    MessageRow
  },
  computed: {
    // @ts-ignore
    chipPosition (message: IMessage): string {
      if (this.$store.state.auth.currentUser.id === message.userId) return 'pullRight'

      return 'pullLeft'
    }
  }
})
</script>

<style scoped>
.pullLeft {
  justify-content: flex-start;
}

.pullRight {
  justify-content: flex-end;
}
</style>
