<template>
<div style="margin:30px;">
    <h2>递交小说</h2>
    <div class="mb-3 row">
    <label for="novelTitle" class="col-sm-2 col-form-label" >小说标题</label>
        <div class="col-sm-10 input-group">
            <input type="text" class="form-control" id="novelTitle" v-model="title">
        </div>
        <div class="form-text">名字应顺应时代发展、严禁反动淫秽色情标题</div>
    </div>
    <div class="mb-3 row">
    <label for="novelAuthor" class="col-sm-2 col-form-label">小说作者</label>
        <div class="col-sm-10 input-group">
            <input type="text" class="form-control" id="novelAuthor" rows="3" v-model="author">
        </div>
        <div class="form-text">小说作者，不强制真实姓名</div>
    </div>
    <div class="mb-3 row">
    <label for="novelProfile" class="col-sm-2 col-form-label">小说简介</label>
        <div class="col-sm-10 input-group">
            <textarea type="text" class="form-control" id="novelProfile" rows="3" v-model="profile"></textarea>
        </div>
        <div class="form-text">简介内容应当概括小说整体，或者突出小说特色</div>
    </div>
    <div class="mb-3 row">
    <label for="novelCategory" class="col-sm-2 col-form-label">小说分类</label>
        <div class="col-sm-10 input-group">
            <select class="form-select" id="novelCategory" @change="selectChange($event)">
                <option
                    v-for="(item,index) in items"
                    v-bind:value="item.id"
                    v-bind:key="index"
                    >{{ item.name }}</option>
            </select>
        </div>
        <div class="form-text">选择最适合小说的一个分类</div>
    </div>
    <div class="mb-3 row">
    <label for="novelFileImg" class="col-sm-2 col-form-label">小说封面</label>
        <div class="col-sm-10 input-group">
            <input type="file" class="form-control" @change="getFileImg($event)" accept=".png, .jpg, .jpge, .heic, .tif, .dng" id="novelFileImg">
        </div>
        <div class="form-text">支持以下格式图像：.png .jpg .jpge .heic .tif .dng</div>
    </div>
    <div class="mb-3 row">
    <label for="novelFile" class="col-sm-2 col-form-label">小说文件</label>
        <div class="col-sm-10 input-group">
            <input type="file" class="form-control" @change="getFile($event)" accept=".txt, .mobi, .epub, .chm, .prc, .md" id="novelFile">
        </div>
        <div class="form-text">支持以下格式图像：.txt .mobi .epub .chm .prc .md</div>
    </div>
    <button class="btn btn-outline-primary" type="button" @click="uploadFile">上传</button>

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
    name: 'UserUpload',
    data(){
        return {
            title: "",
            profile: "",
            fileImg: [],
            file: [],
            category: 0,
            author: "",
            items: this.$store.state.category,
            msg: "NULL",
            msgTitle: "NULL",
        }
    },
    beforeCreate(){
        this.checkToken().then(res=>{
            if(!res){
                this.$router.push({
                    path: '/user/login'
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
        getFileImg(event){
            this.fileImg = event.target.files[0]
            console.log(this.fileImg)
        },
        getFile(event){
            this.file = event.target.files[0]
            console.log(this.file)
        },
        selectChange(e){
            this.category = e.target.value
            console.log("更改小说分类:"+this.category)
        },
        uploadFile(){
            if(this.file.length==0||this.fileImg.length==0||this.author.length==0||this.title.length==0||this.profile.length==0){
                this.showMsg("出错了!","请填写完整表单内容")
                return
            }
            var data = new FormData()
            data.append('cover', this.fileImg)
            data.append('file', this.file)
            data.append('title',this.title)
            data.append('category',this.category)
            data.append('profile',this.profile)
            data.append('author',this.author)
            data.append('user',this.$store.state.userid)
            data.append('time',new Date().getTime())
            console.log(data)
            this.axios.post("/cors/api/auth/novel/upload",data,{
                headers:{
                    "Content-Type": "multipart/form-data; boundary=--------------------"+ new Date().getTime(),
                    AuthToken: this.$store.state.userToken,
                },
            }).then((res)=>{
                var dataReturn = JSON.parse(JSON.stringify(res.data))
                this.updateToken(dataReturn)
                if(dataReturn['status']!=200){
                    this.showMsg("出错了!","上传文件失败  ERR:"+dataReturn['error'])
                }else{
                    this.showMsg("提交成功","上传文件成功，请等待管理员审核，审核通过后将会上架小说")
                    this.$router.push({
                        path: '/'
                    })
                }
            })
        }
    }
}
</script>


