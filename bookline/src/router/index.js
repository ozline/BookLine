import Vue from 'vue'
import VueRouter from 'vue-router'


import index from '../views/index.vue'
import userLogin from '../views/user/login.vue'
import userRegister from '../views/user/register.vue'
import userFavlist from '../views/user/favlist.vue'
import userUpload from '../views/user/upload.vue'
import userHistory from '../views/user/history.vue'
import search from '../views/search.vue'
import notFound from '../views/404.vue'
import review from '../views/admin/review.vue'
import category from '../views/category.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'index',
    component: index,
    meta: {
      title: 'BookLine - 首页',
      check: true,
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
      title: 'BookLine - 我的收藏',
    }
  },
  {
    path: '/user/history',
    name: 'userHistory',
    component: userHistory,
    meta: {
      title: 'BookLine - 浏览历史',
    }
  },
  {
    path: '/user/upload',
    name: 'userupload',
    component: userUpload,
    meta: {
      title: 'BookLine - 文件上传',
    }
  },
  {
    path: '/admin/review',
    name: 'adminReview',
    component: review,
    meta: {
      title: 'BookLine - 小说审查',
    }
  },
  {
    path: '/search/:msg',
    name: 'novelSearch',
    component: search,
    meta: {
      title: 'BookLine - 搜索',
    }
  },
  {
    path: '/category',
    name: 'category',
    component: category,
    meta: {
      title: 'BookLine - 按分类查找',
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
