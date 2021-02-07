import { Getters, Actions, Mutations, Module } from 'vuex-smart-module'
import { IUser } from '../types/models/user'
import firebase from '../plugins/firebase'

class AuthState {
  currentUser: IUser | null = null;
}

class AuthMutations extends Mutations<AuthState> {
  setCurrentUser (user: any): void {
    if (!user) {
      this.state.currentUser = null
      return
    }

    if (process.server) {
      this.state.currentUser = user
      return
    }

    this.state.currentUser = {
      authToken: user.za,
      id: user.uid,
      name: user.displayName,
      email: user.email,
      avatar: user.photoURL,
    }
  }
}

class AuthActions extends Actions<AuthState, AuthGetters, AuthMutations> {
  async login (): Promise<void> {
    const provider = new firebase.auth.GithubAuthProvider()
    try {
      const result = await firebase.auth().signInWithPopup(provider)
      this.commit('setCurrentUser', result.user)

      const db = firebase.firestore()
      const user = await db.collection('users').doc((result.user as any).uid).get()
      if (user.exists) return

      db.collection('users').doc((result.user as any).uid).set({
        id: (result.user as any).uid,
        name: (result.user as any).displayName,
        email: (result.user as any).email,
        avatar: (result.user as any).photoURL,
        createdAt: firebase.firestore.FieldValue.serverTimestamp(),
        updatedAt: firebase.firestore.FieldValue.serverTimestamp(),
      })
    } catch (error) {
      alert(error)
    }
  }

  async logout (): Promise<void> {
    try {
      await firebase.auth().signOut()
      this.commit('setCurrentUser', null)
    } catch (error) {
      alert(error)
    }
  }
}

class AuthGetters extends Getters<AuthState> {
  isAuthenticated (): boolean {
    return !!this.state.currentUser
  }
}

export default new Module({
  state: AuthState,
  getters: AuthGetters,
  mutations: AuthMutations,
  actions: AuthActions,
})
