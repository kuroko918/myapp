<template>
  <v-app dark>
    <v-navigation-drawer
      v-model="drawer"
      v-if="isAuthenticated()"
      :clipped="clipped"
      app
    >
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar
      :clipped-left="clipped"
      app
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title v-text="title" />
      <v-spacer />
      <v-icon
        v-if="isAuthenticated()"
        @click="logout"
      >
        mdi-logout
      </v-icon>
    </v-app-bar>

    <v-main>
      <nuxt />
    </v-main>
  </v-app>
</template>

<script lang='ts'>
import Vue from 'vue'

export default Vue.extend({
  data () {
    return {
      clipped: true,
      drawer: false,
      items: [
        {
          icon: 'mdi-comment',
          title: 'Chat',
          to: '/chat'
        },
        {
          icon: 'mdi-account',
          title: 'Profile',
          to: '/profile'
        },
      ],
      title: 'My App',
    }
  },
  computed: {
    isAuthenticated(): boolean {
      return this.$store.getters['auth/isAuthenticated'];
    }
  },
  methods: {
    async logout () {
      if (!confirm('ログアウトしますか？')) return

      await this.$store.dispatch('auth/logout')
      this.$router.push({ name: 'login' })
    }
  }
})
</script>
