import { Store } from 'vuex'
import { Getters, Actions, Mutations, Module } from 'vuex-smart-module'
import { IMessage } from '../types/models/message'

export class MessageState {
  message!: IMessage
}

class MessagesState extends MessageState {
  messages: IMessage[] = []
}

class MessageMutations extends Mutations<MessagesState> {
  setMessages (messages: IMessage[]): void {
    this.state.messages = messages
  }

  addMessage (message: IMessage): void {
    this.state.messages.push(message)
  }

  updateMessage (message: IMessage): void {
    this.state.messages = this.state.messages.map((m) => {
      if (m.id === message.id) return message
      return m
    })
  }

  deleteMessage (messageId: number): void {
    this.state.messages = this.state.messages.filter(message => {
      return message.id !== messageId
    })
  }
}

class MessageActions extends Actions<MessagesState, MessageGetters, MessageMutations> {
  store!: Store<any>

  $init (store: Store<any>): void {
    this.store = store
  }

  async getMessages (): Promise<void> {
    try {
      const response = await this.store.$axios.$get(`${process.env.URL}/messages`)
      this.commit('setMessages', response.messages)
    } catch (error) {
      alert(error)
    }
  }

  async postMessage (content: string): Promise<void> {
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

  async putMessage ({ messageId, content }: { messageId: number, content: string }): Promise<void> {
    try {
      const params = {
        content
      }
      const response = await this.store.$axios.$put(`${process.env.URL}/message/${messageId}`, params)
      this.commit('updateMessage', response)
    } catch (error) {
      alert(error)
    }
  }

  async deleteMessage (messageId: number): Promise<void> {
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
