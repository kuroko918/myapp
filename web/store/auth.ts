import { Getters, Actions, Mutations, Module } from 'vuex-smart-module'
import firebase from '../plugins/firebase'

class AuthState {
  authToken: string | null = null;
  currentUserId: string | null = null;
}

class AuthMutations extends Mutations<AuthState> {
  setAuthState ({ authToken, currentUserId }: AuthState) {
    if (!authToken || !currentUserId) {
      this.state.authToken = null
      this.state.currentUserId = null
      return
    }

    if (process.server) {
      this.state.authToken = authToken
      this.state.currentUserId = currentUserId
      return
    }

    this.state.authToken = authToken
    this.state.currentUserId = currentUserId
  }
}

class AuthActions extends Actions<AuthState, AuthGetters, AuthMutations> {
  async login (){
    try {
      const provider = new firebase.auth.GithubAuthProvider()
      const result = await firebase.auth().signInWithPopup(provider)
      const user = result.user
      if (!user) throw 'ログインに失敗しました'

      const authToken = await user.getIdToken()
      this.commit('setAuthState', { authToken, currentUserId: user.uid })

      const db = firebase.firestore()
      const userSnap = await db.collection('users').doc(user.uid).get()
      if (userSnap.exists) return

      db.collection('users').doc(user.uid).set({
        id: user.uid,
        name: user.displayName,
        email: user.email,
        avatar: user.photoURL,
        createdAt: firebase.firestore.FieldValue.serverTimestamp(),
        updatedAt: firebase.firestore.FieldValue.serverTimestamp(),
      })
    } catch (error) {
      alert(error)
    }
  }

  async logout () {
    try {
      await firebase.auth().signOut()
      this.commit('setAuthState', { authToken: null, currentUserId: null })
    } catch (error) {
      alert(error)
    }
  }
}

class AuthGetters extends Getters<AuthState> {
  getCurrentUserId (): AuthState['currentUserId'] {
    return this.state.currentUserId
  }

  isAuthenticated (): boolean {
    return !!this.state.authToken
  }
}

export default new Module({
  state: AuthState,
  getters: AuthGetters,
  mutations: AuthMutations,
  actions: AuthActions,
})
