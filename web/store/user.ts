import { Store } from 'vuex'
import { Getters, Actions, Mutations, Module } from 'vuex-smart-module'
import { IUser } from '../types/user'

class UserState {
  user!: IUser
}

class UserMutations extends Mutations<UserState> {
  getUser (user: IUser) {
    this.state.user = user
  }

  updateUser (user: IUser) {
    this.state.user = user
  }
}

class UserActions extends Actions<UserState, UserGetters, UserMutations> {
  store!: Store<any>

  $init (store: Store<any>) {
    this.store = store
  }

  async getUser () {
    try {
      const currentUserId = this.store.state.auth.currentUserId

      const response = await this.store.$axios.$get(`${process.env.URL}/user/${currentUserId}`)
      this.commit('getUser', response)
    } catch (error) {
      alert(error)
    }
  }

  async putUser (user: IUser) {
    try {
      const response = await this.store.$axios.$put(`${process.env.URL}/user/${user.id}`, user)
      this.commit('updateUser', response)
    } catch (error) {
      alert(error)
    }
  }
}

class UserGetters extends Getters<UserState> {
  currentUser (): IUser {
    return this.state.user
  }
}

export default new Module({
  state: UserState,
  getters: UserGetters,
  mutations: UserMutations,
  actions: UserActions,
})
