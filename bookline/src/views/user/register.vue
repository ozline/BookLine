<template>
<div style="margin:30px;">
    <h2>用户注册</h2>
    <div class="form-floating mb-3" style="margin-top:30px;">
        <input class="form-control" placeholder="username" v-model="username">
        <label for="username">用户名</label>
        <div class="form-text">用作登录时的账号</div>
    </div>
    <div class="form-floating mb-3">
        <input type="password" class="form-control" placeholder="passwd" v-model="password">
        <label for="passwd">密码</label>
        <div class="form-text">用作登录的密码</div>
    </div>
    <div class="form-floating mb-3">
        <input type="password" class="form-control" placeholder="passwdConfirm" v-model="passwordConfirm">
        <label for="passwordConfirm">确认密码</label>
    </div>
    <div class="form-floating mb-3">
        <input class="form-control" placeholder="phone" v-model="phone">
        <label for="phone">手机号码</label>
        <div class="form-text">手机号码</div>
    </div>
    <button type="submit" class="btn btn-outline-primary btn-lg" @click="register">注册</button>
</div>
</template>

<script>

export default {
    name: 'UserRegister',
    data(){
        return {
            username: '',
            password: '',
            passwordConfirm: '',
            phone: ''
        }
    },
    methods: {
        register(){
            console.log(this.username)
            console.log(this.password)
            console.log(this.passwordConfirm)
            // todo things...
        }
    }
}
</script>


