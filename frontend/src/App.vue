<script setup>
import { onMounted } from 'vue';
onMounted(() => {

})
</script>

<template>
  <div id="main">
    <router-view></router-view>
  </div>
</template>
 
<style lang="scss" scoped>
#main {
  position: relative;
  width: 100vw;
  height: 100vh;
  background-color: white;
}
</style>