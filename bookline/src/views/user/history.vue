<template>
<div style="margin:30px;">
    <h2>{{ headtitle }}
        <div style="float:right;">
            <button class="btn btn-outline-danger btn-sm" @click="clearHistory()">清空浏览记录</button>
        <span class="badge bg-secondary" id="num" style="margin-left:20px;">Total:{{ num }}</span>
        </div>
    </h2>
    <div class="accordion" id="accordionAll" style="margin-top:30px;">
        <profile
        v-for="novel in data"
        v-bind:key="novel[0].id"
        v-bind:id="novel[0].id"
        v-bind:data="novel[0]"
        v-bind:time="novel[1]"
        v-bind:islocal=1
        ></profile>
    </div>
    <nav aria-label="pageNav" style="margin-top:30px">
        <ul class="pagination justify-content-center">
            <li :class="'page-item '+CSSPrevious">
                <a class="page-link" href="#" @click="movePage(-1)">Previous</a>
            </li>
            <template
            v-for="n in pages"
            >
            <li :class="'page-item '+ CSSPageActive[n] + CSSPage" :key="n" aria-current="page"><a class="page-link" @click="changePage(n)" href="#">{{ n }}</a></li>
            </template>
            <li :class="'page-item '+CSSNext">
                <a class="page-link" href="#" @click="movePage(1)">Next</a>
            </li>
        </ul>
    </nav>

    <!-- Modal -->
    <div class="modal fade" id="modalHistory" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
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
                <button type="button" class="btn btn-outline-danger" data-bs-dismiss="modal" @click="deleteHistory()">清空浏览记录</button>
                <button type="button" class="btn btn-outline-primary" data-bs-dismiss="modal">关闭</button>
            </div>
            </div>
        </div>
    </div>
</div>
</template>

<script>
import novelProfile from '../../components/novel.vue'

const bootstrap = require('bootstrap')

export default {
    name: 'UserHistory',
    data() {
        return {
            headtitle: "",
            num: 0,
            pages: 0,
            page: 1,
            dataAll: [],
            data: [],
            CSSPrevious: "disabled",
            CSSNext: "",
            CSSPage: "",
            CSSPageActive : [],
            msg: "NULL",
            msgTitle: "NULL",
        }
    },
    components: {
        'profile': novelProfile
    },
    created() {
        this.checkToken().then(res=>{
            if(!res){
                this.$router.push({
                    path: '/user/login'
                })
            }else{
                this.headtitle="'"+this.$store.state.username+"'的浏览历史"
                this.loadData()
                this.CSSPageActive = new Array(this.pages+1)
                this.updatePage()
            }
        })
    },
    methods: {
        showMsg(title,msg){
            this.msg = msg
            this.msgTitle = title
            var modal = new bootstrap.Modal(document.getElementById('modalHistory'),{
                keyboard: false
            })
            modal.show()
        },
        loadData(){
            this.CSSPage="disabled"
            this.CSSPrevious="disabled"
            this.CSSNext="disabled"

            var localdata = JSON.parse(localStorage.getItem(this.$store.state.userid))
            if(localdata!=null){
                localdata.sort((a,b)=>{
                    return (b[1]-a[1])
                })
                this.dataAll = localdata
                this.num = localdata.length
                this.pages = Math.ceil(this.num/10)
                this.CSSPage=""
                this.CSSPrevious = this.page==1? "disabled" : ""
                this.CSSNext = this.page>=this.pages? "disabled" : ""
            }else{
                this.dataAll = []
                this.num = 0
            }
        },
        changePage(n){
            this.page = n
            this.loadData()
            this.updatePage()
        },
        movePage(n){
            this.page += n
            this.loadData()
            this.updatePage()
        },
        updatePage(){
            if(this.num==0) return
            var endNum = this.page*10>this.num? this.num : this.page*10 //防止超标
            this.data = this.dataAll.slice((this.page-1)*10,endNum)
            for(var index=1; index<this.pages+1;index++){
                    this.CSSPageActive[index] = ""
            }
            this.CSSPageActive[this.page] = "active "
        },
        clearHistory(){
            this.showMsg("确认","操作不可恢复，确认清空浏览记录吗？")
        },
        deleteHistory(){
            localStorage.removeItem(this.$store.state.userid)
            this.data=[]
            this.loadData()
            this.CSSPageActive = new Array(this.pages+1)
            this.updatePage()
            this.CSSPrevious="disabled"
            this.CSSNext= "disabled"
            this.pages = 0
        }
    }
}
</script>