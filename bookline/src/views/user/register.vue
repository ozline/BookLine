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
    <button type="submit" class="btn btn-outline-primary btn-lg" @click="register">注册</button>
    <router-link to="/user/login" class="btn btn-outline-success btn-lg" style="margin-left:20px;">登录</router-link>
</div>
</template>

<script>

export default {
    name: 'UserRegister',
    data(){
        return {
            username: '',
            password: '',
            passwordConfirm: ''
        }
    },
    beforeCreate(){
        this.checkToken().then(res=>{
            if(res){
                this.$router.push({
                    path: '/'
                })
            }
        })
    },
    methods: {
        register(){
            if(this.username.length==0||this.password.length==0||this.passwordConfirm.length==0){
                alert("请检查填写内容")
                return
            }
            if(this.password!=this.passwordConfirm){
                alert("两次密码输入不正确")
                return
            }
            var data = new FormData()
            data.append('username', this.username)
            data.append('password',this.password)
            this.axios.post("/cors/api/user/register",data
            ).then((res)=>{
                var dataReturn = JSON.parse(JSON.stringify(res.data))
                if(dataReturn['status']!=200){
                    alert("注册账号失败  ERR:"+dataReturn['error'])
                }else{
                    alert("注册成功，即将跳转登录页面")
                    console.log(dataReturn)
                    this.$router.push({
                        path: '/user/login'
                    })
                }
            })
        }
    }
}
</script>


