<template>
<div style="margin:30px;">
    <h2>用户登录</h2>
    <div class="form-floating mb-3" style="margin-top:30px;">
        <input class="form-control" placeholder="username" v-model="username">
        <label for="username">用户名</label>
        <div class="form-text">注册时使用的用户名</div>
    </div>
    <div class="form-floating mb-3">
        <input type="password" class="form-control" placeholder="passwd" v-model="password">
        <label for="passwd">密码</label>
        <div class="form-text">用作登录的密码</div>
    </div>
    <button type="submit" class="btn btn-outline-primary btn-lg" @click="login">登录</button>
    <router-link to="/user/register" class="btn btn-outline-danger btn-lg" style="margin-left:20px;">注册</router-link>

    <!-- Modal -->
    <div class="modal fade" id="modal" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="staticBackdropLabel">{{ msgTitle }}</h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                {{ msg }}
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-outline-primary" data-bs-dismiss="modal">我知道了</button>
            </div>
            </div>
        </div>
    </div>
</div>
</template>

<script>

const bootstrap = require('bootstrap')

export default {
    name: 'UserLogin',
    data(){
        return {
            username: '',
            password: '',
            msg: "NULL",
            msgTitle: "NULL",
        }
    },
    created(){
        this.checkToken().then(res=>{
            if(res){
                this.$router.push({
                    path: '/'
                })
            }
        })
    },
    methods: {
        showMsg(title,msg){
            this.msg = msg
            this.msgTitle = title
            var modal = new bootstrap.Modal(document.getElementById('modal'),{
                keyboard: false
            })
            modal.show()
        },
        login(){
            if(this.username.length==0||this.password.length==0){
                this.showMsg("出错了!","用户名或密码填写不正确")
                return
            }
            var data = new FormData()
            data.append('username', this.username)
            data.append('password',this.password)
            this.axios.post("/cors/api/user/login",data
            ).then((res)=>{
                var dataReturn = JSON.parse(JSON.stringify(res.data))
                if(dataReturn['status']!=200){
                    this.showMsg("出错了!","登录账号失败  ERR:"+dataReturn['error'])
                }else{
                    this.saveToken(dataReturn['token'])
                    this.$store.state.isadmin = dataReturn['isadmin']
                    this.$router.push({
                        path: '/'
                    })
                }
            })
        }
    }
}
</script>


