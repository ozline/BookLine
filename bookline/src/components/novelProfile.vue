<template>
    <div class="accordion-item">
        <h2 class="accordion-header" :id="'header'+id">
            <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" :data-bs-target="'#collapse'+id" aria-expanded="true" :aria-controls="'collapse'+id">
                {{ title }}
            </button>
        </h2>
        <div :id="'collapse'+id" class="accordion-collapse collapse" :aria-labelledby="'header'+id" data-bs-parent="#accordionAll">
            <div class="accordion-body">
                {{ introduce }}
                <button type="button" class="btn btn-primary" @click="detail(id)">测试vueRouter跳转传值功能 id:{{ id }}</button>
            </div>
        </div>
        </div>
</template>

<script>

export default{
    name: "novelProfile",
    props: ["title","introduce","id"],
    methods:{
        detail(id){
            this.$router.push({
                path: '/detail/'+id,
                query: {
                    novelName: 'test'
                }
            })
        }
    }
}

</script>