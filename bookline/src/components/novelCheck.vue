<template>
    <div class="accordion-item" v-show="show">
        <h2 class="accordion-header" :id="'header'+id">
            <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" :data-bs-target="'#collapse'+id" aria-expanded="true" :aria-controls="'collapse'+id">
                {{ data.title }}
            </button>
        </h2>
        <div :id="'collapse'+id" class="accordion-collapse collapse show" :aria-labelledby="'header'+id" data-bs-parent="#accordionAll">
            <div class="accordion-body">
                <div class="card mb-3 text-dark bg-light">
                    <div class="row g-0">
                        <div class="col-md-4">
                            <img :src="data.coverPath" class="img-fluid rounded-start" />
                        </div>
                        <div class="col-md-8">
                            <div class="card-body">
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
                                    <div class="btn-group" role="group" style="float:right;">
                                        <button type="button" class="btn btn-outline-warning" @click="passNovel()">通过</button>
                                        <button type="button" class="btn btn-outline-danger" @click="refuseNovel()">拒绝并删除</button>
                                    </div>
                                </a>
                            </div>
                        </div>
                    </div>
                </div>
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
                    <button type="button" class="btn btn-outline-primary" data-bs-dismiss="modal">关闭</button>
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
    name: "novelCheck",
    props: ["id","data"],
    data(){
        return{
            date: "",
            show: true,
            msg: "NULL",
            msgTitle: "NULL",
        }
    },
    components:{
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
        refuseNovel(){
            this.axios.delete("/cors/api/auth/novel/"+this.id,{
                headers:{
                    AuthToken: this.getToken(),
                },
            }).then((res)=>{
                var dataReturn = JSON.parse(JSON.stringify(res.data))
                this.updateToken(dataReturn)
                if(dataReturn['status']==200){
                    this.show = false
                }else{
                    this.showMsg("出错了!","删除失败  ERR:"+dataReturn['error'])
                }
            })
        },
        passNovel(){
            console.log(this.getToken())
            var data = new FormData();
            this.axios.patch("/cors/api/auth/novel/"+this.id+"/1",data,{
                headers:{
                    AuthToken: this.getToken(),
                },
            }).then((res)=>{
                var dataReturn = JSON.parse(JSON.stringify(res.data))
                if(dataReturn['status']==200){
                    this.show = false
                }else{
                    this.showMsg("出错了!","审核通过失败 ERR:"+dataReturn['error'])
                }
            })
        }
    }
}

</script>