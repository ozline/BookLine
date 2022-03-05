import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import router from './router/index.js'
import store from './store'
import cookie from 'vue-cookie'


Vue.config.productionTip = false

Vue.prototype.cookie = cookie

Vue.prototype.getToken = function(){
  var cookietoken = this.cookie.get("AuthToken")
  if(cookietoken!=''){
    this.$store.state.userToken = this.cookie.get("AuthToken")
  }
  return this.$store.state.userToken
}

Vue.prototype.saveToken = function(token){
  this.$store.state.userToken = token
  this.cookie.set('AuthToken',token)
  this.$store.state.islogin = true
}

Vue.prototype.deleteToken = function(){
  this.cookie.delete("AuthToken")
  this.$store.state.userToken = ''
  this.$store.state.islogin = false
}

Vue.prototype.updateToken = function(data){
  if("token" in data && data['token']!=''){
    // console.log("token: ["+data['token']+"]")
    this.saveToken(data['token'])
  }
}

Vue.prototype.checkToken = async function(){
  this.$store.state.isadmin = 0
  var dataReturn =[]
  await this.axios.get("/cors/api/auth/checkToken",{
    headers:{
        AuthToken: this.getToken(),
    },
  }).then((res)=>{
    dataReturn = JSON.parse(JSON.stringify(res.data))
  })
  var tmp = dataReturn['status']==200
  if(tmp){
    this.$store.state.userdata = dataReturn
    this.$store.state.username = dataReturn['data']['username']
    this.$store.state.userid = dataReturn['data']['id']
    this.$store.state.isadmin = dataReturn['data']['isadmin']
    this.$store.state.islogin = true
  }
  return tmp
}

router.beforeEach((to,from,next)=>{
  // console.log('组件跳转 从'+from.path+" 至"+to.name)
  if(to.meta.title){
    document.title = to.meta.title
  }
  next()
})

new Vue({
  render: h => h(App),
  router,
  cookie,
  store
}).$mount('#app')