<template>
<nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">BookLine</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarContent" aria-controls="navbarContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <router-link @click.native="flushCom" to="/" class="nav-link" aria-current="page">首页</router-link>
          </li>
          <li class="nav-item">
            <router-link @click.native="flushCom" to="/category" class="nav-link" aria-current="page">分类</router-link>
          </li>
          <li class="nav-item">
            <router-link @click.native="flushCom" to="/user/upload" class="nav-link" aria-current="page" v-show="this.$store.state.islogin">上传</router-link>
          </li>
          <li><router-link @click.native="flushCom" to="/user/login" class="nav-link" v-show="!this.$store.state.islogin">登录</router-link></li>
          <li class="nav-item dropdown" v-show="this.$store.state.islogin">
            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
              我的
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarScrollingDropdown">
              <li><a class="dropdown-item disabled">昵称 <b style="float:right;">{{ this.$store.state.username }}</b></a></li>
              <li><a href="#" class="dropdown-item disabled">ID<b style="float:right;">{{ this.$store.state.userid }}</b></a></li>
              <li><a href="#" @click="exitUser()" class="dropdown-item">退出</a></li>
              <li><hr class="dropdown-divider"></li>
              <li><router-link @click.native="flushCom" to="/user/favlist" class="dropdown-item">我的收藏</router-link></li>
              <li><router-link @click.native="flushCom" to="/user/history" class="dropdown-item">浏览历史</router-link></li>
              <li v-show="this.$store.state.isadmin==1"><hr class="dropdown-divider"></li>
              <li><router-link @click.native="flushCom" to="/admin/review" class="dropdown-item" v-show="this.$store.state.isadmin==1">小说审核<span class="badge bg-danger" style="margin-left:10px;">Admin</span></router-link></li>
            </ul>
          </li>
          <li>
          </li>
        </ul>
        <div class="d-flex justify-content-start">
          <input class="form-control me-2" placeholder="小说名/作者/简介.." v-model="msgSearch">
          <button class="btn btn-outline-success btn-sm" @click="search()">Search</button>
        </div>
      </div>
    </div>

    <!-- Modal -->
    <div class="modal fade" id="searchModal" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
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
</nav>
</template>

<script>

const bootstrap = require('bootstrap')


export default{

  name: 'siteNavbar',
  data(){
    return{
      msgSearch: '',
      msg: "NULL",
      msgTitle: "NULL",
    }
  },
  methods: {
    showMsg(title,msg){
        this.msg = msg
        this.msgTitle = title
        var modal = new bootstrap.Modal(document.getElementById('searchModal'),{
            keyboard: false
        })
        modal.show()
    },
    search(){
      if(this.msgSearch==''){
        this.showMsg("出错了!","请输入正确的关键词")
        return
      }
      this.$router.push({path: '/search/'+this.msgSearch})
      if(this.$route.path.substr(1,6)=='search'){ //防止搜索页面重新搜索
        this.$router.go(0)
      }
      this.msgSearch = ''
    },
    exitUser(){
      this.deleteToken()
      this.$router.push({path: '/user/login'})
    },
    flushCom(){
      this.$router.go(0)
    }
  }
}

</script>