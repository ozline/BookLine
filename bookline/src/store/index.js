import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    userToken: 'userToken',
    username: 'username',
    userid: 'userid',
    isadmin: 0,
    islogin: false,
    userdata: [],
    category:[
      {id:0, name: '未分类'},
      {id:1, name: '言情'},
      {id:2, name: '校园'},
      {id:3, name: '恐怖'},
      {id:4, name: '惊悚'},
      {id:5, name: '推理'}
  ]
  },
  getters: {
  },
  mutations: {
  },
  actions: {
    updatetoken(token){
      this.userToken = token
    }
  },
  modules: {
  }
})
