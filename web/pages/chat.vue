<template>
  <div>
    <v-container>
      <MessageBody />
    </v-container>
    <v-footer
      class="px-16 py-5"
      padless
      fixed
    >
      <MessageForm
        :onsubmit="createMessage"
      />
    </v-footer>
  </div>
</template>

<script lang='ts'>
import Vue from 'vue'
import MessageBody from '../components/chat/MessageBody.vue'
import MessageForm from '../components/chat/MessageForm.vue'
import { IMessage } from '../types/message'

export default Vue.extend({
  components: {
    MessageBody,
    MessageForm,
  },
  created () {
    this.getMessages()
  },
  methods: {
    async getMessages () {
      await this.$store.dispatch('message/getMessages')
    },
    async createMessage (content: IMessage['content']) {
      const submitContent = content.trim()
      if (!submitContent) return
      if (submitContent.length > 140) return

      await this.$store.dispatch('message/postMessage', content)

      this.scrollToBottom()
    },
    scrollToBottom () {
      const cardElements = document.querySelectorAll('.v-card')
      // @ts-ignore
      this.$scrollTo(cardElements[cardElements.length -1])
    }
  }
})
</script>

<style scoped>
.container {
  margin-bottom: 100px;
}

footer {
  display: unset;
}
</style>
