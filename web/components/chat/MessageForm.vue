<template>
  <v-form
    class="d-flex align-center"
    @submit.prevent="onSubmit"
  >
    <v-text-field v-model="inputValue" value="message.content" class="pa-10" rows="1" />
    <v-btn :disabled="!innerInputValue" @click="onSubmit">
      送信
    </v-btn>
  </v-form>
</template>

<script lang='ts'>
import Vue, { PropType } from 'vue'
import { IMessage } from '../../types/models/message'

type DataType = {
  innerInputValue: string
}

type PropObjType = (content: string) => Promise<void>

export default Vue.extend({
  props: {
    message: {
      type: Object as PropType<IMessage>,
      default: () => {}
    },
    onsubmit: {
      type: Function as PropType<PropObjType>,
      required: true
    },
    toggleEditing: {
      type: Function as PropType<() => {}>,
      default: null
    }
  },
  data (): DataType {
    return {
      innerInputValue: ''
    }
  },
  computed: {
    inputValue: {
      get (): string {
        return this.innerInputValue
      },
      set (value: string): void {
        this.innerInputValue = value
      }
    }
  },
  methods: {
    async onSubmit (e: Event): Promise<void> {
      e.preventDefault()

      await this.onsubmit(this.innerInputValue)
      this.innerInputValue = ''
      if (this.toggleEditing) { this.toggleEditing() }
    }
  }
})
</script>
