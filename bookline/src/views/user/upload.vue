<template>
<div style="margin:30px;">
    <h2>递交小说</h2>
    <div class="mb-3 row">
    <label for="novelTitle" class="col-sm-2 col-form-label">小说标题</label>
        <div class="col-sm-10 input-group">
            <input type="text" class="form-control" id="novelTitle">
        </div>
        <div class="form-text">名字应顺应时代发展、严禁反动淫秽色情标题</div>
    </div>
    <div class="mb-3 row">
    <label for="novelFileImg" class="col-sm-2 col-form-label">小说封面</label>
        <div class="col-sm-10 input-group">
            <input type="file" class="form-control" @change="getFileImg($event)" accept=".png" id="novelFileImg">
            <!-- <button class="btn btn-outline-primary" type="button">上传</button> -->
        </div>
        <div class="form-text">暂只允许上传.png图像</div>
    </div>
    <div class="mb-3 row">
    <label for="novelFile" class="col-sm-2 col-form-label">小说文件</label>
        <div class="col-sm-10 input-group">
            <input type="file" class="form-control" @change="getFile($event)" accept=".txt" id="novelFile">
            <!-- <button class="btn btn-outline-primary" type="button">上传</button> -->
        </div>
        <div class="form-text">暂只允许上传txt文件</div>
    </div>
    <button class="btn btn-outline-primary" type="button" @click="uploadFile">上传</button>
</div>
</template>

<script>

export default {
    name: 'UserUpload',
    data(){
        return {
            title: "",
            fileImg: [],
            file: []
        }
    },
    methods: {
        getFileImg(event){
            this.fileImg = event.target.files[0]
            console.log(this.fileImg)
        },
        getFile(event){
            this.file = event.target.files[0]
            console.log(this.file)
        },
        uploadFile(){
            if(this.file.length==0||this.fileImg.length==0){
                alert("请选择需要上传的文件")
                return
            }0
            var data = new FormData()
            data.append('file', this.fileImg)
            data.append('file', this.file)
            this.axios.post("TODO:url待填写",data,{
                headers:{
                    "Content-Type": "multipart/form-data"
                },
                params:{
                    title: this.title,
                    id: 0 //此处应当从cookie读取
                }
            }).then((response)=>{
                console.log(response)
            })
        }
    }
}
</script>


