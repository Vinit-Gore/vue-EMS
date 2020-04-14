import Vue from 'vue'
import upperFirst from 'lodash/upperFirst'
import camelCase from 'lodash/camelCase'
import App from './App.vue'
import router from './router'
import store from './store/store'
import 'nprogress/nprogress.css'
import Vuelidate from 'vuelidate'

Vue.use(Vuelidate)

Vue.config.productionTip = false

const requireComponent = require.context(
  './components',
  false,
  /Base[A-Z]\w+\.(vue|js)$/
)

requireComponent.keys().forEach(fileName => {
  const componentConfig = requireComponent(fileName)

  const componentName = upperFirst(
    camelCase(fileName.replace(/^\.\/(.*)\.\w+$/, '$1'))
  )

  Vue.component(componentName, componentConfig.default || componentConfig)
})

Vue.filter('toStdDate', value => {
  if (!value) return ''
  const d = new Date(value)
  return d.toLocaleString(['en-US'], {
    month: 'short',
    day: '2-digit',
    year: 'numeric'
  })
})

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
