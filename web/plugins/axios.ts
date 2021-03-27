import * as jsCookie from 'js-cookie'
import { NuxtAxiosInstance } from '@nuxtjs/axios'

interface IAxios {
  $axios: NuxtAxiosInstance
}

export default ({ $axios }: IAxios) => {
  $axios.onRequest(config => {
    const vuex = jsCookie.get('vuex')
    if (!vuex) return

    config.headers.common.Authorization = `bearer ${JSON.parse(vuex).authToken}`
    return config
  })
}
