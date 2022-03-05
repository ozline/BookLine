<template>
<div style="margin:30px;">
    <h2>{{ headtitle }}
        <span class="badge bg-secondary" id="num" style="float:right;">Total:{{ num }}</span>
    </h2>
    <div class="progress" :style="progressShow">
        <div class="progress-bar progress-bar-striped progress-bar-animated" role="progressbar" aria-valuenow="99" aria-valuemin="0" aria-valuemax="100" style="width: 100%"></div>
    </div>
    <div class="accordion" id="accordionAll" style="margin-top:30px;">
        <profile
        v-for="(novel,index) in data"
        v-bind:key="novel.id"
        v-bind:id="novel.id"
        v-bind:data="novel"
        v-on:remove="remove(index)"
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
</div>
</template>

<script>
import novelProfile from '../../components/novelCheck.vue'

export default {
    name: 'AdminReview',
    data() {
        return {
            headtitle: "小说审查",
            num: 0,
            pages: 0,
            page: 1,
            data: [],
            CSSPrevious: "disabled",
            CSSNext: "",
            CSSPage: "",
            CSSPageActive : [],
            progressShow: ""
        }
    },
    components: {
        'profile': novelProfile
    },
    created() {
        this.checkToken().then(res=>{
            if(!res||this.$store.state.isadmin!=1){
                this.$router.push({
                    path: '/user/login'
                })
            }else{
                this.loadData()
                this.CSSPageActive = new Array(this.pages+1)
                this.updatePage()
            }
        })
    },
    methods: {
        async getAxios(){
        var tmp = []
        await this.axios.get('/cors/api/auth/novel/review',{
            params: {
                page : this.page
            },
            headers:{
                AuthToken: this.$store.state.userToken,
            }
        }).then((response) =>{
            tmp = response
            })
        return tmp
        },
        loadData(){
            this.progressShow= ""
            this.CSSPage="disabled"
            this.CSSPrevious="disabled"
            this.CSSNext="disabled"
            this.getAxios().then(res=>{
                var dataReturn = JSON.parse(JSON.stringify(res.data))
                if(dataReturn['success']==false){
                    if(dataReturn['status']==999){
                        // this.$router.push({
                        //     path: '/user/login'
                        // })
                        return
                    }
                    alert("获取失败 ERR:"+dataReturn['error'])
                    return
                }
                this.data = dataReturn['data']['items']
                this.num = dataReturn['data']['total']
                this.pages = Math.ceil(this.num/10)
                this.progressShow= "display:none;"
                this.CSSPage=""
                this.CSSPrevious = this.page==1? "disabled" : ""
                this.CSSNext = this.page>=this.pages? "disabled" : ""
            })
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
            for(var index=1; index<this.pages+1;index++){
                    this.CSSPageActive[index] = ""
            }
            this.CSSPageActive[this.page] = "active "
        }
    }
}
</script>