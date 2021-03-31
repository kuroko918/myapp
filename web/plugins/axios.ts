import jsCookie from 'js-cookie'
import { NuxtAxiosInstance } from '@nuxtjs/axios'
import { Context } from '@nuxt/types'
import axios from 'axios'

interface IAxios {
  $axios: NuxtAxiosInstance
  store: Context['store']
  redirect: Context['redirect']
}

export default ({ $axios, store, redirect }: IAxios) => {
  $axios.onRequest(config => {
    const vuex = jsCookie.get('vuex')
    if (!vuex) return config

    config.headers.common.Authorization = `bearer ${JSON.parse(vuex).authToken}`
    return config
  })

  $axios.onError(async error => {
    try {
      if (error.response?.data.message.includes('ID token has expired at:')) {
        const vuex = jsCookie.get('vuex')
        if (!vuex) return

        const refreshToken = JSON.parse(vuex).refreshToken

        const params = new URLSearchParams();
        params.append('grant_type', 'refresh_token')
        params.append('refresh_token', refreshToken)
        const response = await axios.post(
          `https://securetoken.googleapis.com/v1/token?key=${process.env.API_KEY}`,
          params, {
            headers: {
              'Content-Type': 'application/x-www-form-urlencoded',
            },
          },
        )

        store.commit('auth/setAuthState', {
          authToken: response.data.access_token,
          currentUserId: response.data.user_id,
          refreshToken,
        })

        axios(error.config)
      }
    } catch {
      store.commit('auth/setAuthState', {
        authToken: null,
        currentUserId: null,
        refreshToken: null,
      })
      redirect('/login')
    }
  })
}
