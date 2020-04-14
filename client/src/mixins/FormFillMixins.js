export const FormFillMixins = {
  inheritAttrs: false,
  methods: {
    updateValue(event) {
      this.$emit('input', event.target.value)
    }
  }
}
