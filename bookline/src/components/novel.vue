<template>
    <div class="accordion-item" v-show="show">
        <h2 class="accordion-header" :id="'header'+id">
            <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" :data-bs-target="'#collapse'+id" aria-expanded="true" :aria-controls="'collapse'+id">
                <span class="badge bg-success" v-show="islocal==1||isfav==1">
                    {{ transferTime(time) }}
                </span>
                <b>{{ data.title }} </b>
            </button>
        </h2>
        <div :id="'collapse'+id" class="accordion-collapse collapse" :aria-labelledby="'header'+id" data-bs-parent="#accordionAll">
            <div class="accordion-body">
                <div class="card mb-3 text-dark bg-light">
                    <div class="row g-0">
                        <div class="col-md-4">
                            <img :src="data.coverPath" class="img-fluid rounded-start" style="max-height=200px;" />
                        </div>
                        <div class="col-md-8">
                            <div class="card-body">
                                <h5 calss="card-title"> {{ data.title }}</h5>
                                <p class="card-text">{{ data.profile }}</p>
                                <p class="card-text"><small class="text-muted">更新时间: {{ transferTime(data.update) }}</small></p>
                                <button type="button" class="btn btn-outline-primary" data-bs-toggle="offcanvas" :data-bs-target="'#offcanvas'+id" :aria-controls="'offcanvas'+id" @click="detailClick()">查看详情</button>
                                <button type="button" class="btn btn-outline-dark" style="margin-left:10px;" v-show="this.$store.state.islogin&&isfav!=1" @click="addFav()">加入收藏</button>
                                <button type="button" class="btn btn-outline-danger" style="margin-left:10px;" v-show="this.$store.state.islogin&&isfav==1" @click="deleteFav()">取消收藏</button>
                                <button type="button" class="btn btn-outline-danger" style="margin-left:10px;" @click="deleteHistory()" v-show="this.$store.state.islogin&&islocal==1? true : false">删除浏览记录</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="offcanvas offcanvas-end" tabindex="-1" :id="'offcanvas'+id" :aria-labelledby="'offcanvasLabel'+id">
            <div class="offcanvas-header">
                <h5 class="offcanvas-title" id="'offcanvasLabel'+id">{{ data.title }}</h5>
                <button type="button" class="btn-close" data-bs-dismiss="offcanvas" aria-label="Close"></button>
            </div>
            <div class="offcanvas-body">
                <listitem title="编号" small="DBNUMBER" :data="id"></listitem>
                <listitem title="分类" small="CATEGORY" :data="this.$store.state.category[data.category].name"></listitem>
                <listitem title="作者" small="AUTHOR" :data="data.author"></listitem>
                <listitem title="最后更新" small="UPDATE" :data="transferTime(data.update)"></listitem>
                <a class="list-group-item list-group-item-action" aria-current="true">
                    <div class="d-flex w-100 justify-content-between">
                        <h5 class="mb-1">简介</h5>
                        <small>PROFILE</small>
                    </div>
                    <p class="mb-1">{{ data.profile }}</p>
                </a>
                <a class="list-group-item list-group-item-action" aria-current="true">
                    <div class="d-flex w-100 justify-content-between">
                        <h5 class="mb-1">下载</h5>
                        <small>DOWNLOAD</small>
                    </div>
                    <p>文件由用户 <b>ID {{ data.user }}</b> 提供</p>
                    <a type="button" class="btn btn-outline-primary" :href="data.coverPath" style="margin-right:10px;" :download="data.title" target="_blank">封面下载</a>
                    <a type="button" class="btn btn-outline-success" :href="data.filePath" :download="data.title" target="_blank">小说下载</a>
                    <button type="button" class="btn btn-outline-danger" @click="deleteNovel()" style="float:right;" v-show="this.$store.state.isadmin==1">删除小说</button>
                </a>
            </div>
        </div>

        <!-- Modal -->
        <div class="modal fade" :id="'modal'+id" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
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

import listitem from './listitem.vue'

const bootstrap = require('bootstrap')

export default{
    name: "novelProfile",
    props: ["id","data","islocal","time","isfav"],
    data(){
        return{
            date: "",
            show: true,
            msg: "NULL",
            msgTitle: "NULL",
        }
    },
    components: {
        'listitem' : listitem
    },
    methods: {
        showMsg(title,msg){
            this.msg = msg
            this.msgTitle = title
            var modal = new bootstrap.Modal(document.getElementById('modal'+this.id),{
                keyboard: false
            })
            modal.show()
        },
        transferTime(cTime){
			var jsonDate = new Date(parseInt(cTime));
			Date.prototype.format = function(format) {
				var o = {
					"y+": this.getFullYear(),
					"M+": this.getMonth() + 1,
					"d+": this.getDate(),
					"h+": this.getHours(),
					"m+": this.getMinutes(),
					// "s+": this.getSeconds()
				};
				if (/(y+)/.test(format)) {
					format = format.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
				}
				for (var k in o) {
					if (new RegExp("(" + k + ")").test(format)) {
						format = format.replace(RegExp.$1, RegExp.$1.length == 1 ? o[k] : ("" + o[k]).substr("" + o[k].length));
					}
				}
				return format;
			};
			var newDate = jsonDate.format("yyyy-MM-dd hh:mm");
			return newDate
		},
        deleteNovel(){
            this.axios.delete("/cors/api/auth/novel/delete/"+this.id,{
                headers:{
                    AuthToken: this.getToken(),
                },
            }).then((res)=>{
                var dataReturn = JSON.parse(JSON.stringify(res.data))
                if(dataReturn['status']==200){
                    this.show = false
                }else{
                    this.showMsg("出错了!","删除小说失败,错误:"+dataReturn['error'])
                }
            })
        },
        detailClick(){
            if(!this.$store.state.islogin){
                this.$router.push({path: '/user/login'})
                return
            }
            var history = JSON.parse(localStorage.getItem(this.$store.state.userid))
            if(history == null) history = []
            for(var i in history){
                if(history[i][0].id == this.id){
                    history[i][1] = new Date().getTime() //刷新浏览历史
                    localStorage.setItem(this.$store.state.userid, JSON.stringify(history))
                    return
                }
            }
            var item = [this.data,new Date().getTime()]
            history.push(item)
            localStorage.setItem(this.$store.state.userid, JSON.stringify(history))
        },
        deleteHistory(){
            var history = JSON.parse(localStorage.getItem(this.$store.state.userid))
            if(history == null) history = []
            var ans = null
            for(var i in history){
                if(history[i][0].id == this.id){
                    ans = i
                    break
                }
            }
            history.splice(ans,1) //删除这条记录
            localStorage.setItem(this.$store.state.userid, JSON.stringify(history))
            this.$router.go(0)
        },
        addFav(){
            this.axios.post("/cors/api/auth/novel/fav/"+this.id,null,{
                headers:{
                    AuthToken: this.getToken(),
                },
            }).then((res)=>{
                var dataReturn = JSON.parse(JSON.stringify(res.data))
                this.updateToken(dataReturn)
                if(dataReturn['status']==200){
                    this.showMsg("收藏小说","加入收藏成功")
                }else{
                    this.showMsg("出错了!","收藏失败,错误:"+dataReturn['error'])
                }
            })
        },
        deleteFav(){
            this.axios.delete("/cors/api/auth/novel/fav/"+this.id,{
                headers:{
                    AuthToken: this.getToken(),
                },
            }).then((res)=>{
                var dataReturn = JSON.parse(JSON.stringify(res.data))
                this.updateToken(dataReturn)
                if(dataReturn['status']==200){
                    this.show = false
                }else{
                    this.showMsg("出错了!","取消收藏失败,错误:"+dataReturn['error'])
                }
            })
        },
    }
}

</script>