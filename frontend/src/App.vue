<script setup>
import { onMounted } from 'vue';
import Global from "./Global/index.js"
import { GetSettingsJson } from '../wailsjs/go/main/App'
onMounted(async () => {
  const settings=await GetSettingsJson()
  console.log(settings)
  Global.ExportFileType=settings.ExportFileType
  Global.ExportFileUrl=settings.ExportFileUrl
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