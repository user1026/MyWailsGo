<template>
  <div id="context">
    <el-container>
      <!-- <el-row :gutter="0">
        <el-col :span="leftSpan"> -->
      <el-aside :width="leftWidth">
        <div id="left">
          <el-menu
          ref="leftElMenu"
            :default-active="computerUrl"
            class="leftMenu el-menu-vertical-demo"
            router 
          >
            <el-menu-item :index="computerUrl">
              <el-icon v-if="activeUrl==computerUrl">
              <img  src="../assets/images/computer_active.png" alt="" srcset="">
              </el-icon>
              <el-icon v-else>
                <img  src="../assets/images/computer_noactive2.png" alt="" srcset="">
                </el-icon> 
              <template #title>电脑信息</template>
            </el-menu-item>
            <el-menu-item :index="PCHardwareUrl">
              <el-icon v-if="activeUrl==PCHardwareUrl">
               <img  src="../assets/images/PCHardware_active.png" alt="" srcset="">
              </el-icon>
              <el-icon v-else>
               <img  src="../assets/images/PCHardware_noactive.png" alt="" srcset="">
              </el-icon>
              <template #title>电脑装机</template>
            </el-menu-item>
            <el-menu-item :index="SettingsUrl">
              <el-icon v-if="activeUrl==SettingsUrl">
               <img  src="../assets/images/settings_active.png" alt="" srcset="">
              </el-icon>
              <el-icon v-else>
               <img  src="../assets/images/settings_noactive.png" alt="" srcset="">
              </el-icon>
              <template #title>设置</template>
            </el-menu-item>
          </el-menu>
        </div>
      </el-aside>
      <!-- </el-col>
        <el-col :span="rightSpan"> -->
      <el-container>
        <el-header></el-header>
        <el-main>
          <div id="right">
            <router-view></router-view>
          </div>
        </el-main>
      </el-container>
      <!-- </el-col>
    </el-row> -->
    </el-container>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted,watch } from "vue";
import {GetCpuInfo,GetUsingCpuInfo,GetRamInfo} from "../../wailsjs/go/main/App.js"
import Global from "../Global/index.js"
import computer from "@/views/computer/index.vue"
const leftElMenu=ref(null)
const leftWidth = ref("20vh");
const activeUrl=ref("")
const computerUrl=ref("/computerInfo")
const PCHardwareUrl=ref("/PCHardware")
const SettingsUrl=ref("/settings")
 onMounted(()=>{

 })

watch(()=>Global.activeUrl,(n,o)=>{
  activeUrl.value=n;
})
</script>
<style lang="scss" scoped>
#context {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}

#left {
  
  height: 100vh;
  background-color: white;

  .leftMenu {
    background-color: white;
    border-right: none;
  }
}

#right {
  width: 100%;
  height: 100vh;
  background-color: white;
}
.el-main{
  padding: 0;
}
img{
  width: 20px;
  height: 20px;
}
</style>
