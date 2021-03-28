<template>
  <v-container>
    <v-form
      v-model="valid"
      @submit.prevent="onSubmit"
    >
      <v-row no-gutters>
        <v-col cols="12">
          <div
            class="text-h4 font-weight-bold my-10 mx-auto"
          >
            プロフィール
          </div>
        </v-col>

        <v-col cols="12">
          <v-hover
            v-slot="{ hover }"
          >
            <v-avatar
              :class="{ 'on-hover': hover }"
              class="mb-10"
              color="grey lighten-2"
              size="300"
              @click="onClick"
            >
              <v-img
                :src="cacheImage === '' ? avatar : cacheImage"
              >
                <v-icon
                  :class="{ 'show-btns': hover }"
                  class="mx-auto"
                  color="rgba(255, 255, 255, 0)"
                  x-large
                >
                  mdi-camera
                </v-icon>
              </v-img>
            </v-avatar>
          </v-hover>
          <v-file-input
            ref="avatar"
            prepend-icon=""
            hide-input
            accept="image/*"
            @change="onImagePicked"
          >
          </v-file-input>
        </v-col>

        <v-col cols="12">
          <v-text-field
            v-model="name"
            :rules="nameRules"
            :counter="30"
            label="ユーザー名"
            class="mb-10 px-10"
            required
          />
        </v-col>

        <v-col cols="12">
          <v-text-field
            v-model="email"
            :rules="emailRules"
            label="メールアドレス"
            class="mb-10 px-10"
            required
          />
        </v-col>

        <v-col cols="12">
          <v-btn
            :loading="updating"
            :disabled="updating || !valid"
            class="mx-auto"
            large
            @click="onSubmit"
          >
            <v-icon
              class="mr-2"
            >
              mdi-cog
            </v-icon>
            プロフィールを更新
          </v-btn>
        </v-col>
      </v-row>
    </v-form>
  </v-container>
</template>

<script lang='ts'>
import Vue from 'vue'
import uploadImage from '../plugins/uploadImage'
import { IUser } from '../types/user'

export type DataType = {
  updating: boolean;
  valid: boolean;
  cacheImage: string;
  avatar: IUser['avatar'];
  name: IUser['name'];
  nameRules: [(value: string) => boolean | string, (value: string) => boolean | string];
  email: IUser['email'];
  emailRules: [(value: string) => boolean | string, (value: string) => boolean | string];
}

export default Vue.extend({
  data (): DataType {
    return {
      updating: false,
      valid: true,
      cacheImage: '',
      avatar: '',
      name: '',
      nameRules: [
        (value: string) => !!value || '名前を入力してください',
        (value: string) => value.length <= 30 || '名前は30文字以下で入力してください',
      ],
      email: '',
      emailRules: [
        (value: string) => !!value || 'メールアドレスを入力してください',
        (value: string) => /[a-zA-Z0-9.a-zA-Z0-9.!#$%&'*+-/=?^_`{|}~]+@[a-zA-Z0-9]+\.[a-zA-Z]+/.test(value) || '適切な形のメールアドレスを入力してください',
      ],
    }
  },
  created () {
    this.getUser()
  },
  methods: {
    async getUser () {
      await this.$store.dispatch('user/getUser')
      this.avatar = this.$store.state.user.user.avatar || '/images/flutter.png'
      this.name = this.$store.state.user.user.name
      this.email = this.$store.state.user.user.email
    },
    onClick() {
      // @ts-ignore
      this.$refs.avatar.$refs.input.click()
    },
    onImagePicked(file: File) {
      if (!file) return
      if (file.name.lastIndexOf('.') <= 0) return

      const fr = new FileReader()
      fr.onload = () => this.cacheImage = fr.result as string
      fr.readAsDataURL(file)
    },
    async onSubmit (e: Event) {
      e.preventDefault()

      this.updating = true

      const currentUser = this.$store.state.user.user

      let imageURL: string | null = null
      if (this.cacheImage !== '') imageURL = await uploadImage(this.cacheImage, currentUser.id)

      const user = {
        ...currentUser,
        avatar: imageURL || this.avatar,
        name: this.name,
        email: this.email,
      }

      await this.$store.dispatch('user/putUser', user)
      this.updating = false
    },
  },
})
</script>

<style scoped>
.col.col-12 {
  text-align: center
}

.on-hover {
  opacity: 0.5;
}

.show-btns {
  color: rgba(255, 255, 255, 1) !important;
}
</style>
