<template>
  <div>
    <MessageBody />
    <MessageForm
      :onsubmit="createMessage"
    />
  </div>
</template>

<script lang='ts'>
import Vue from 'vue'
import MessageBody from '../components/chat/MessageBody.vue'
import MessageForm from '../components/chat/MessageForm.vue'

export default Vue.extend({
  components: {
    MessageBody,
    MessageForm,
  },
  created (): void {
    this.indexMessage()
  },
  methods: {
    async indexMessage (): Promise<void> {
      await this.$store.dispatch('message/getMessages')
    },
    async createMessage (content: string): Promise<void> {
      if (!content) { return }

      await this.$store.dispatch('message/postMessage', content)
    }
  }
})
</script>
