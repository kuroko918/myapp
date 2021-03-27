import { Store } from 'vuex'
import { Getters, Actions, Mutations, Module } from 'vuex-smart-module'
import { IMessage } from '../types/message'

class MessageState {
  message!: IMessage
}

class MessagesState extends MessageState {
  messages: IMessage[] = []
}

class MessageMutations extends Mutations<MessagesState> {
  getMessages (messages: IMessage[]) {
    this.state.messages = messages
  }

  addMessage (message: IMessage) {
    this.state.messages.push(message)
  }

  updateMessage (message: IMessage) {
    this.state.messages = this.state.messages.map((m) => {
      if (m.id === message.id) return message
      return m
    })
  }

  deleteMessage (messageId: IMessage['id']) {
    this.state.messages = this.state.messages.filter(message => {
      return message.id !== messageId
    })
  }
}

class MessageActions extends Actions<MessagesState, MessageGetters, MessageMutations> {
  store!: Store<any>

  $init (store: Store<any>) {
    this.store = store
  }

  async getMessages () {
    try {
      const response = await this.store.$axios.$get(`${process.env.URL}/messages`)
      this.commit('getMessages', response.messages)
    } catch (error) {
      alert(error)
    }
  }

  async postMessage (content: IMessage['content']) {
    try {
      const params = {
        content
      }
      const response = await this.store.$axios.$post(`${process.env.URL}/message`, params)
      this.commit('addMessage', response)
    } catch (error) {
      alert(error)
    }
  }

  async patchMessage ({ messageId, content }: { messageId: IMessage['id'], content: IMessage['content'] }) {
    try {
      const params = {
        content
      }
      const response = await this.store.$axios.$patch(`${process.env.URL}/message/${messageId}`, params)
      this.commit('updateMessage', response)
    } catch (error) {
      alert(error)
    }
  }

  async deleteMessage (messageId: IMessage['id']) {
    try {
      await this.store.$axios.$delete(`${process.env.URL}/message/${messageId}`)
      this.commit('deleteMessage', messageId)
    } catch (error) {
      alert(error)
    }
  }
}

class MessageGetters extends Getters<MessagesState> {
}

export default new Module({
  state: MessagesState,
  getters: MessageGetters,
  mutations: MessageMutations,
  actions: MessageActions,
})
