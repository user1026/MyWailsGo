<template>
    <div id="context">
        <el-container>
        <!-- <el-row :gutter="0">
        <el-col :span="leftSpan"> -->
            <el-aside :width="leftWidth">
            <div id="left">
                <div class="leftHeader">
                    <el-icon @click="changeCollapse">
                                <Menu />
                    </el-icon>
                </div>
                    <el-menu
                    :collapse-transition="false"
                    default-active="2"
                    class="leftMenu el-menu-vertical-demo"
                    :collapse="isCollapse"
                    >
                        <el-menu-item index="1">  <el-icon><document /></el-icon>图表</el-menu-item>
                        <el-menu-item index="2">  <el-icon><document /></el-icon>硬件选择</el-menu-item>
                        <el-menu-item></el-menu-item>
                        <el-menu-item></el-menu-item>
                        <el-menu-item></el-menu-item>
                    </el-menu>
            </div>
        </el-aside>
        <!-- </el-col>
   
        <el-col :span="rightSpan"> -->
            <el-container>
                <el-header></el-header>
                <el-main>
                    <div id="right">
                    <el-button @click="t">+</el-button>
            </div>
                </el-main>
        </el-container>
        <!-- </el-col>
    </el-row> -->
</el-container>
    </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
const leftSpan=ref(4);
const rightSpan=ref(20);
const isCollapse=ref(false);
const leftWidth=ref("20vh")
const t=()=>{
    leftSpan.value=8;
    rightSpan.value=16
}
const changeCollapse=()=>{
    isCollapse.value=!isCollapse.value;
   // isCollapse.value?leftSpan.value=1:leftSpan.value=4; 
   isCollapse.value?leftWidth.value="5vw":leftWidth.value="10vw"; 
}
</script>
<style lang='scss' scoped>
#context{
    width: 100vw;
    height: 100vh;
    overflow: hidden;
}
#left{
    width: 100%;
    height: 100vh;
    background-color: #c6cbd6;
    .leftMenu{
        width: 100%;
        background-color: #789be6;
        border-right: none;
    }
   
}
#right{
    width: 100%;
    height: 100vh;
    background-color: blue;
}
</style>