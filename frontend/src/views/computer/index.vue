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
        >{{cpuInfo.Name}}</el-descriptions-item
      >
      <el-descriptions-item label="cpu当前频率" label-align="right" align="center"
        >{{cpuInfo.GHZ+'GHZ'}}</el-descriptions-item
      >
      <el-descriptions-item label="cpu核心" label-align="right" align="center"
        >{{cpuInfo.Cores}}核{{cpuInfo.Threads}}线程</el-descriptions-item
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
import tu from "@/components/echarts.vue"
import {GetCpuInfo,GetUsingCpuInfo} from "../../../wailsjs/go/main/App.js"
import ComputerInfo from "./ComputerInfo.js"
const UsingCpu=ref("0")
const chart=ref(null)
const option=ref({title:[{},{text:""}]})
const width=ref("200px");
const height=ref("200px")
const cpuInfo=ref({Name:"",GHZ:0,Cores:0,Threads:0})
const ramInfo=ref({})
  onBeforeMount(async ()=>{
  const computerInfo=await ComputerInfo();
  option.value=computerInfo.option.value;
  cpuInfo.value=computerInfo.cpuInfo.value;
  ramInfo.value=computerInfo.ramInfo.value;
  //console.log(option.value,cpuInfo.value,ramInfo.value,"option")
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
  background-color: white;
  overflow: hidden;
  .elFont{
    color: white;
  }
}
</style>