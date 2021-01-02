<template>
  <div v-if="editing">
    <v-chip
      close
      close-icon="mdi-delete"
      label
      outlined
      @click:close="onDelete"
    >
      <v-icon
        color="white darken-2"
        @click="toggleEditing"
      >
        mdi-close-thick
      </v-icon>
      <MessageForm
        :message="message"
        :onsubmit="updateMessage"
        :toggle-editing="toggleEditing"
      />
    </v-chip>
  </div>
  <div v-else>
    <v-chip
      label
      outlined
    >
      <WorkerInfo :user="message.user" />
      <div>{{ displayedTime }}</div>
      <div>{{ message.content }}</div>
      <v-icon
        color="white darken-2"
        @click="toggleEditing"
      >
        mdi-cog
      </v-icon>
    </v-chip>
  </div>
</template>

<script lang='ts'>
import Vue, { PropType } from 'vue'
import { IMessage } from '../../types/models/message'
import MessageForm from './MessageForm.vue'
import WorkerInfo from './WorkerInfo.vue'

export type DataType = {
  editing: boolean
}

export default Vue.extend({
  components: {
    MessageForm,
    WorkerInfo
  },
  props: {
    chipPosition: {
      type: Function as PropType<() => {}>,
      required: true
    },
    message: {
      type: Object as PropType<IMessage>,
      required: true
    }
  },
  data (): DataType {
    return {
      editing: false
    }
  },
  computed: {
    displayedTime (): string {
      if (this.message.createdAt === this.message.updatedAt) { return this.message.createdAt }

      return this.message.updatedAt + '(編集済)'
    }
  },
  methods: {
    async updateMessage (content: string): Promise<void> {
      const params = {
        messageId: this.message.id,
        content
      }
      await this.$store.dispatch('message/putMessage', params)
    },
    async deleteMessage (messageId: number): Promise<void> {
      await this.$store.dispatch('message/deleteMessage', messageId)
    },
    toggleEditing (): void {
      this.editing = !this.editing
    },
    onDelete (): void {
      if (!confirm('本当に削除しますか？')) { return }

      this.deleteMessage(this.message.id)
    }
  }
})
</script>

<style scoped>
.v-chip {
  display: inline-block;
  height: 70px;
}
</style>
