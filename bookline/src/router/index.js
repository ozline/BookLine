import Vue from 'vue'
import VueRouter from 'vue-router'


import index from '../views/index.vue'
import userLogin from '../views/user/login.vue'
import userRegister from '../views/user/register.vue'
import userFavlist from '../views/user/favlist.vue'
import userHistory from '../views/user/history.vue'
import userUpload from '../views/user/upload.vue'
import adminLogin from '../views/admin/login.vue'
import adminReview from '../views/admin/review.vue'
import detail from '../views/detail.vue'
import search from '../views/search.vue'
import notFound from '../views/404.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'index',
    component: index,
    meta: {
      title: 'BookLine - 首页'
    }
  },
  {
    path: '/user/login',
    name: 'userLogin',
    component: userLogin,
    meta: {
      title: 'BookLine - 用户登录'
    }
  },
  {
    path: '/user/register',
    name: 'userRegister',
    component: userRegister,
    meta: {
      title: 'BookLine - 账号注册'
    }
  },
  {
    path: '/user/favlist',
    name: 'userFavlist',
    component: userFavlist,
    meta: {
      title: 'BookLine - 我的收藏'
    }
  },
  {
    path: '/user/history',
    name: 'userHistory',
    component: userHistory,
    meta: {
      title: 'BookLine - 浏览历史'
    }
  },
  {
    path: '/user/upload',
    name: 'userupload',
    component: userUpload,
    meta: {
      title: 'BookLine - 文件上传'
    }
  },
  {
    path: '/admin/login',
    name: 'adminLogin',
    component: adminLogin,
    meta: {
      title: 'BookLine - 管理员登录'
    }
  },
  {
    path: '/admin/review',
    name: 'adminReview',
    component: adminReview,
    meta: {
      title: 'BookLine - 小说审核'
    }
  },
  {
    path: '/detail/:novelID',
    name: 'novelDetail',
    component: detail,
    meta: {
      title: 'BookLine - 小说详情'
    }
  },
  {
    path: '/search/:msg',
    name: 'novelSearch',
    component: search,
    meta: {
      title: 'BookLine - 搜索'
    }
  },
  {
    path: '*',
    component: notFound,
    meta:{
      title: 'BookLine - 404'
    }
  }
]

const router = new VueRouter({
  base: '/',
  mode: 'history',
  routes
})

export default router
