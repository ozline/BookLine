<template>
<a class="list-group-item list-group-item-action" aria-current="true">
    <div class="d-flex w-100 justify-content-between">
        <h5 class="mb-1">{{ title }}</h5>
            <small>{{ small }}</small>
    </div>
    <h4><span class="badge bg-secondary">{{ data }}</span></h4>
</a>
</template>
<script>

export default{
    name: "listItem",
    props: ["title","small","data"],
}
</script>