<template>
  <div>
    <v-card
      v-for="(message) in this.$store.state.message.messages"
      :key="message.id"
      class="mx-auto mb-5"
      max-width="800"
      outlined
    >
      <MessageRow
        :message="message"
      />
    </v-card>
  </div>
</template>

<script lang='ts'>
import Vue from 'vue'
import MessageRow from './MessageRow.vue'
import WorkerInfo from './WorkerInfo.vue'

export default Vue.extend({
  components: {
    MessageRow,
    WorkerInfo,
  },
})
</script>
