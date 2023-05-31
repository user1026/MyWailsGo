<template>
  <div id="context">
    <el-container>
      <!-- <el-row :gutter="0">
        <el-col :span="leftSpan"> -->
      <el-aside :width="leftWidth">
        <div id="left">
            <el-icon @click="changeCollapse">
              <Menu />
            </el-icon>
          <el-menu
          ref="leftElMenu"
            default-active="/index/echart"
            class="leftMenu el-menu-vertical-demo"
            :collapse="isCollapse"
            router 
          >
            <el-menu-item index="/index/echarts">
              <el-icon>
                <document />
              </el-icon>
              <template #title>电脑信息</template>
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
import { ref, reactive, onMounted } from "vue";
import {GetCpuInfo,GetUsingCpuInfo,GetMemInfo} from "../../wailsjs/go/main/App.js"
import computer from "@/views/computer/index.vue"
const leftSpan = ref(4);
const rightSpan = ref(20);
const isCollapse = ref(false);
const leftElMenu=ref(null)
const leftWidth = ref("20vh");
const t = () => {
  leftSpan.value = 8;
  rightSpan.value = 16;
};
onMounted(()=>{
    leftWidth.value=leftElMenu.width+"px"
  //   GetCpuInfo().then(res=>{console.log(res)})
  // GetUsingCpuInfo().then(res=>{console.log(res,"using")})
  GetMemInfo().then(res=>{console.log(res,"men")})

})
const changeCollapse = () => {
  isCollapse.value = !isCollapse.value;
  // isCollapse.value?leftSpan.value=1:leftSpan.value=4;
   leftWidth.value=leftElMenu.width+"px";
};
</script>
<style lang="scss" scoped>
#context {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}

#left {
  
  height: 100vh;
  background-color: #c6cbd6;

  .leftMenu {
    background-color: #e9ecf3;
    border-right: none;
  }
}

#right {
  width: 100%;
  height: 100vh;
  background-color: blue;
}
.el-main{
  padding: 0;
}
</style>
