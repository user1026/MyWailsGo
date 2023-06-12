<template>
  <div id="computer">
  <div>
    <el-descriptions title="cpu信息" :column="3" border class="elFont">
      <el-descriptions-item
        label="cpu型号"
        label-align="right"
        align="center"
        label-class-name="my-label"
        class-name="my-content"
        >{{info.modelName}}</el-descriptions-item
      >
      <el-descriptions-item label="频率" label-align="right" align="center"
        >{{info.mhz+'mhz'}}</el-descriptions-item
      >
      <el-descriptions-item label="cpu核心" label-align="right" align="center"
        >{{info.cores}}</el-descriptions-item
      >
     
    </el-descriptions>
  </div>
    <div>
        <tu :option="option" style="width: 300px;height: 300px;"  ref="chart"></tu>
        <!-- :width="width" :height="height" -->
    </div>
    
  </div>
</template>
<script setup>
import { ref,reactive,onMounted,onBeforeMount,watch} from 'vue'
import tu from "@/views/echarts/index.vue"
import {GetCpuInfo,GetUsingCpuInfo} from "../../../wailsjs/go/main/App.js"
import CpuInfo from "./cpuInfo.js"
const UsingCpu=ref("0")
const chart=ref(null)
const option=ref({title:[{},{text:""}]})
const width=ref("200px");
const height=ref("200px")
const info=ref({modelName:"",mhz:0,cores:0})
  onBeforeMount(async ()=>{
  const cpuinfo=await CpuInfo();
  option.value=cpuinfo.option.value;
  info.value=cpuinfo.cpuInfo.value;
  console.log(option.value,info.value,"option")
})

onMounted(()=>{
    GetUsingCpuInfo().then(res=>{
      option.value.title[1].text=res.toFixed(2)+"";
      chart.value.EchartsInit();
  });
})
</script>
<style lang='scss' scoped>
#computer{
  width: 100%;
  height: 100%;
  position: relative;
  background-color: black;
  overflow: hidden;
  .elFont{
    color: white;
  }
}
</style>