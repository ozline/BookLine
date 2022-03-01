<template>
<div style="margin:30px;">
    <h2>'{{ this.$route.params.msg }}' 的搜索结果 <span class="badge bg-secondary" id="num" style="float:right;">Total:{{ num }}</span></h2>
    <button type="button" class="btn btn-primary" @click="test()">临时API测试代码运行</button>
    <div class="accordion" id="accordionAll" style="margin-top:30px;">
        <profile
        v-for="(novel,index) in data"
        v-bind:key="novel.province_id"
        v-bind:introduce="novel.province_id"
        v-bind:title="novel.province"
        v-bind:id="novel.province_id"
        v-on:remove="remove(index)"
        ></profile>
    </div>
    <nav aria-label="pageNav" style="margin-top:30px">
        <!-- 此处分页功能将会由后端实现 -->
        <ul class="pagination justify-content-end">
            <template
            v-for="n in pages"
            >
            <li class="page-item" :key="n"><a class="page-link" @click="changePage(n)">{{ n }}</a></li>
            </template>
        </ul>
    </nav>
</div>
</template>

<script>
// let key = "c33e230f80318034bce399982d3187eb"; // API所需的用户key
let key = "004cce07c2dab96288d743150e76188d"

import novelProfile from '../components/novelProfile.vue'

export default {
    name: 'novelDetail',
    data() {
        return {
        num: 0,
        pages: 0,
        data: [],
        }
    },
    components: {
        'profile': novelProfile
    },
    methods: {
        async getAxios(){
        let tmp = []
        await this.axios.get('/api/springTravel/citys',{
            params: {
            key : key
            }
        }).then((response) =>{
            tmp = response
            })
        return tmp
        },
        test(){
        this.getAxios().then(res=>{
            this.data = JSON.parse(JSON.stringify(res.data))['result']
            console.log(this.data)
            this.num = this.data.length
            this.pages = Math.ceil(this.num/10)
        })
        console.log(1)
        },
        changePage(n){
        console.log(this.pages)
        }
    }
}
</script>


