const { defineConfig } = require('@vue/cli-service')
const proxy = require('http-proxy-middleware')
module.exports = defineConfig({
  transpileDependencies: true
})


//设置代理，测试的时候使 用
module.exports = {
  devServer:{
    host : 'localhost',
    port : 8080,
    proxy: {
      '/cors': {
        target: 'http://127.0.0.1:6666/',
        changeOrigin: true,
        // ws: true,
        pathRewrite: {
          '^/cors': ''
        }
      }
    }
  }
}