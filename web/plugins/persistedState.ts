import createPersistedState from 'vuex-persistedstate'
import jsCookie from 'js-cookie'
import cookie from 'cookie'
import { Context } from '@nuxt/types'

const cookieStorage = (req: Context['req'], isDev: Context['isDev']) => ({
  getItem: (key: string) => process.client ? jsCookie.get(key) : cookie.parse(req.headers.cookie || '')[key],
  setItem: (key: string, value: string) => jsCookie.set(key, value, { expires: 365, secure: !isDev }),
  removeItem: (key: string) => jsCookie.remove(key),
})

export default ({ store, req, isDev }: Context) => {
  createPersistedState({
    reducer: (state: any) => ({
      authToken: state.auth.authToken,
      refreshToken: state.auth.refreshToken,
      currentUserId: state.auth.currentUserId,
    }),
    storage: cookieStorage(req, isDev),
    getState: cookieStorage(req, isDev).getItem,
    setState: cookieStorage(req, isDev).setItem,
  })(store)
}
