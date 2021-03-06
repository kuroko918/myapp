import { Actions, Module } from 'vuex-smart-module'
import Cookie from 'cookie'
import { Context } from '@nuxt/types'
import auth from './auth'
import message from './message'
import user from './user'

class RootActions extends Actions {
  nuxtServerInit ({ req, store }: Context): void {
    if (!req.headers.cookie) return

    const cookie = Cookie.parse(req.headers.cookie)
    const vuex = cookie.vuex && JSON.parse(cookie.vuex)
    const { authToken, currentUserId, refreshToken } = vuex
    if (vuex) store.commit('auth/setAuthState', { authToken, currentUserId, refreshToken })
  }
}

const root = new Module({
  actions: RootActions,
  modules: {
    auth,
    message,
    user,
  }
})

export const {
  state,
  getters,
  mutations,
  actions,
  modules,
  plugins,
} = root.getStoreOptions()
