<template>
  <div>
    <v-list-item two-line>
      <div v-if="getCurrentUserId() === message.userId">
        <WorkerInfo :user="message.user" />
      </div>
      <v-list-item-content>
        <v-list-item-title class="headline mb-2">
          <div v-if="!editing">
            {{ message.content }}
          </div>
          <div v-else>
            <MessageForm
              :message="message"
              :onsubmit="updateMessage"
              :toggle-editing="toggleEditing"
            />
          </div>
        </v-list-item-title>
        <v-list-item-subtitle>
          {{ $dateFns.format(new Date(message.createdAt), 'yyyy-MM-dd HH:mm') + editedMark }}
        </v-list-item-subtitle>
      </v-list-item-content>
      <div v-if="getCurrentUserId() !== message.userId">
        <WorkerInfo :user="message.user" />
      </div>
    </v-list-item>

    <div v-if="getCurrentUserId() === message.userId">
      <v-card-actions>
        <div
          class="text-body-1 font-weight-bold ml-2"
        >
          {{ message.user.name }}
        </div>
        <v-spacer />
        <v-tooltip top>
          <template v-slot:activator="{ on, attrs }">
            <v-icon
              class="mx-3 mb-3"
              color="white darken-2"
              v-bind="attrs"
              v-on="on"
              @click="toggleEditing"
            >
              {{ editing ? 'mdi-undo' : 'mdi-cog' }}
            </v-icon>
          </template>
          <span>{{ editing ? 'キャンセル' : '編集する' }}</span>
        </v-tooltip>

        <v-tooltip top>
          <template v-slot:activator="{ on, attrs }">
            <v-icon
              class="mx-3 mb-3"
              color="white darken-2"
              v-bind="attrs"
              v-on="on"
              @click="onDelete"
            >
              mdi-delete
            </v-icon>
          </template>
          <span>削除する</span>
        </v-tooltip>
      </v-card-actions>
    </div>
  </div>
</template>

<script lang='ts'>
import Vue, { PropType } from 'vue'
import { IMessage } from '../../types/message'
import MessageForm from './MessageForm.vue'
import WorkerInfo from './WorkerInfo.vue'

export type DataType = {
  editing: boolean
}

export default Vue.extend({
  components: {
    MessageForm,
    WorkerInfo,
  },
  props: {
    message: {
      type: Object as PropType<IMessage>,
      required: true,
    }
  },
  data (): DataType {
    return {
      editing: false,
    }
  },
  computed: {
    editedMark (): string {
      if (this.message.createdAt === this.message.updatedAt) return ''

      return '（編集済）'
    },
    getCurrentUserId (): string {
      return this.$store.getters['auth/getCurrentUserId'];
    }
  },
  methods: {
    async updateMessage (content: IMessage['content']) {
      const submitContent = content.trim()
      if (!submitContent) return
      if (submitContent.length > 140) return

      const params = {
        messageId: this.message.id,
        content
      }
      await this.$store.dispatch('message/patchMessage', params)
    },
    async deleteMessage (messageId: IMessage['id']) {
      await this.$store.dispatch('message/deleteMessage', messageId)
    },
    toggleEditing () {
      this.editing = !this.editing
    },
    onDelete () {
      if (!confirm('本当に削除しますか？')) return

      this.deleteMessage(this.message.id)
    },
  }
})
</script>
