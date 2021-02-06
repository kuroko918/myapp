import { Actions, Module } from 'vuex-smart-module'
import Cookie from 'cookie'
import { Context } from '@nuxt/types'
import auth from './auth'
import message from './message'

class RootActions extends Actions {
  nuxtServerInit (context: Context): void {
    if (!context.req.headers.cookie) return

    const cookie = Cookie.parse(context.req.headers.cookie)
    const vuex = cookie.vuex && JSON.parse(cookie.vuex)
    // @ts-ignore
    if (vuex) this.commit('auth/setCurrentUser', vuex.currentUser)
  }
}

const root = new Module({
  actions: RootActions,
  modules: {
    auth,
    message,
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
