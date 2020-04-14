<template>
  <div>
    <label v-if="label">{{ label }}</label>
    <select
      :value="value"
      @change="updateValue"
      v-bind="$attrs"
      v-on="$listeners"
    >
      <option
        v-for="option in options"
        :value="option"
        :key="option.id"
        :selected="option === value"
        >{{ option }}</option
      >
    </select>
  </div>
</template>

<script>
import { FormFillMixins } from '@/mixins/FormFillMixins.js'
export default {
  mixins: [FormFillMixins],
  props: {
    options: {
      type: Array,
      required: true
    },
    value: [String, Number],
    label: String
  }
}
</script>
