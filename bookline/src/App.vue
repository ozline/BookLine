<template>
  <div>
    <navbar></navbar>
    <router-view></router-view>
  </div>
</template>

<script>

import navBar from './components/navBar.vue'

export default{
  name: 'indexPage',
  components: {
    'navbar': navBar
  }
}
</script>