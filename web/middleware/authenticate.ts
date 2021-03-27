import { Context } from '@nuxt/types'

export default ({ store, route, redirect }: Context) => {
  if (!store.getters['auth/isAuthenticated']() && route.name !== 'login') {
    redirect('/login')
  }
  if (store.getters['auth/isAuthenticated']() && route.name === 'login') {
    redirect('/chat')
  }
}
