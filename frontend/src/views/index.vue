<template>
    <div id="context">
        <el-row :gutter="0">
        <el-col :span="leftSpan">
            <div id="left">
                <div class="leftHeader">
                    <el-icon @click="changeCollapse">
                                <Menu />
                    </el-icon>
                </div>
                    <el-menu
                    default-active="2"
                    class="leftMenu"
                    :collapse="isCollapse"
                    >
                        <el-menu-item>图表</el-menu-item>
                        <el-menu-item>硬件选择</el-menu-item>
                        <el-menu-item></el-menu-item>
                        <el-menu-item></el-menu-item>
                        <el-menu-item></el-menu-item>
                    </el-menu>
            </div>
        </el-col>
        <el-col :span="rightSpan">
            <div id="right">
                    <el-button @click="t">+</el-button>
            </div>
        </el-col>
    </el-row>
    </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
const leftSpan=ref(4);
const rightSpan=ref(20);
const isCollapse=ref(true);
const t=()=>{
    leftSpan.value=8;
    rightSpan.value=16
}
const changeCollapse=()=>{
    isCollapse.value=!isCollapse.value
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
    background-color: #203561;
    .leftMenu{
        width: 100%;
        background-color: #203561;
        border-right: none;
    }
   
}
#right{
    width: 100%;
    height: 100vh;
    background-color: blue;
}
</style>