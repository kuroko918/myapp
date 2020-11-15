import createPersistedState from 'vuex-persistedstate'
import * as jsCookie from 'js-cookie'
import Cookie from 'cookie'
import { Context } from '@nuxt/types'

const cookieStorage = (req: any, isDev: boolean) => ({
  getItem: (key: string) => (process.client) ? jsCookie.get(key) : Cookie.parse(req.headers?.cookie || '')[key],
  setItem: (key: string, value: string) => jsCookie.set(key, value, { expires: 7, secure: !isDev }),
  removeItem: (key: string) => jsCookie.remove(key)
})

export default ({ store, req, isDev }: Context) => {
  createPersistedState({
    reducer: (state: any) => ({
      currentUser: state.auth.currentUser
    }),
    storage: cookieStorage(req, isDev),
    getState: cookieStorage(req, isDev).getItem,
    setState: cookieStorage(req, isDev).setItem
  })(store)
}
