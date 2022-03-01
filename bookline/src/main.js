import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import router from './router/index.js'


Vue.config.productionTip = false

router.beforeEach((to,from,next)=>{
  if(to.meta.title){
    document.title = to.meta.title
  }
  next()
})

new Vue({
  render: h => h(App),
  router
}).$mount('#app')