import jsCookie from 'js-cookie'
import { NuxtAxiosInstance } from '@nuxtjs/axios'
import axios from 'axios'

interface IAxios {
  $axios: NuxtAxiosInstance
  redirect: any
}

export default ({ $axios, redirect }: IAxios) => {
  $axios.onRequest(config => {
    const vuex = jsCookie.get('vuex')
    if (!vuex) return config

    config.headers.common.Authorization = `bearer ${JSON.parse(vuex).authToken}`
    return config
  })

  $axios.onError(async error => {
    console.log('error', error)
    console.log('error.response', error.response)
    try {
      if (error.response?.data.message.includes('ID token has expired')) {
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

        jsCookie.set('authToken', response.data.id_token, { expires: 365 })

        axios(error.config)
      }
    } catch {
      jsCookie.remove('authToken')
      jsCookie.remove('currentUserId')
      jsCookie.remove('refreshToken')
      redirect('/login')
    }
  })
}
