<template>
<div style="margin:30px;">
    <h2>'{{ this.$route.query.novelName }}' 小说详情 <span class="badge bg-secondary" id="num" style="float:right;">ID:{{ this.$route.params.novelID }}</span></h2>
    <span>测试项目,测试传值: {{ this.$route.params.novelID }}</span>
    <div class="row">
        <div class="col-4" style="max-width: 200px;">
        <div class="list-group" id="list-tab" role="tablist">
            <a class="list-group-item list-group-item-action active" id="list-cover-list" data-bs-toggle="list" href="#list-cover" role="tab" aria-controls="list-cover">封面</a>
            <a class="list-group-item list-group-item-action" id="list-profile-list" data-bs-toggle="list" href="#list-profile" role="tab" aria-controls="list-profile">简介</a>
            <a class="list-group-item list-group-item-action" id="list-download-list" data-bs-toggle="list" href="#list-download" role="tab" aria-controls="list-download">下载</a>
        </div>
        </div>
        <div class="col-8">
        <div class="tab-content" id="nav-tabContent">
            <div class="tab-pane fade show active" id="list-cover" role="tabpanel" aria-labelledby="list-cover-list">这是封面</div>
            <div class="tab-pane fade" id="list-profile" role="tabpanel" aria-labelledby="list-profile-list">小说：穿越之重生我是狼王</div>
            <div class="tab-pane fade" id="list-download" role="tabpanel" aria-labelledby="list-download-list">本地下载</div>
        </div>
        </div>
    </div>
</div>
</template>

<script>

// var store = require('store')

// let key = "c33e230f80318034bce399982d3187eb"; // API所需的用户key
let key = "004cce07c2dab96288d743150e76188d"

export default {
    name: 'novelDetail',
    data() {
        return {
        num: 0,
        pages: 0,
        data: [],
        novelID : 0
        }
    },
}
</script>


