const { defineConfig } = require('@vue/cli-service')
const proxy = require('http-proxy-middleware')
module.exports = defineConfig({
  transpileDependencies: true
})


//设置代理，测试的时候使用
module.exports = {
  devServer:{
    host : 'localhost',
    port : 8080,
    proxy: {
      '/api': {
        target: 'https://apis.juhe.cn/',
        changeOrigin: true,
        ws: true,
        pathRewrite: {
          '^/api': '/'
        }
      }
    }
  }
}