import { Module } from 'vuex-smart-module'

const root = new Module({
  modules: {}
})

export const {
  state,
  getters,
  mutations,
  actions,
  modules,
  plugins
} = root.getStoreOptions()
