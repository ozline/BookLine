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
            <a class="nav-link" aria-current="page" href="/">首页</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="/user/upload">上传</a>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
              用户
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarScrollingDropdown">
              <li><a class="dropdown-item" href="/user/login">登录</a></li>
              <li><a class="dropdown-item" href="/user/register">注册</a></li>
              <li><a class="dropdown-item" href="#">退出</a></li>
              <li><hr class="dropdown-divider"></li>
              <li><a class="dropdown-item" href="/user/favlist">收藏</a></li>
              <li><a class="dropdown-item" href="/user/history">历史</a></li>
            </ul>
          </li>
        </ul>
        <div class="d-flex justify-content-start">
          <input class="form-control me-2" placeholder="小说名/作者/.." v-model="msg">
          <button class="btn btn-outline-success btn-sm" @click="search(msg)">Search</button>
        </div>
      </div>
    </div>
</nav>
</template>

<script>

export default{

  name: 'siteNavbar',
  data(){
    return{
      msg: ''
    }
  },
  methods: {
    search(msg){
      this.msg =''
      this.$router.push({
        path: '/search/'+msg
      })
    }
  }
}

</script>