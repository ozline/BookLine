<template>
<div style="margin:30px;">
    <h2>管理员登录</h2>
    <div class="form-floating mb-3" style="margin-top:30px;">
        <input class="form-control" placeholder="phone" v-model="phone">
        <label for="phone">手机号码</label>
        <div class="form-text">注册时使用的手机号码,不支持使用用户名登录</div>
    </div>
    <div class="form-floating mb-3">
        <input type="password" class="form-control" placeholder="passwd" v-model="password">
        <label for="passwd">密码</label>
        <div class="form-text">用作登录的密码</div>
    </div>
    <button type="submit" class="btn btn-outline-primary btn-lg" @click="login">登录</button>
</div>
</template>

<script>

export default {
    name: 'AdminLogin',
    data(){
        return {
            phone: '',
            password: '',
        }
    },
    methods: {
        login(){
            console.log(this.username)
            console.log(this.password)
            // todo things...
            // this.router.push('/')
            // this.router.go()
        }
    }
}
</script>


