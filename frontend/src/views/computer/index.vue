<template>
    <div>
        <tu :option="option" ref="chart"></tu>
    </div>
</template>
<script setup>
import { ref,reactive,onMounted } from 'vue'
import tu from "@/views/echarts/index.vue"
import {GetCpuInfo,GetUsingCpuInfo} from "../../../wailsjs/go/main/App.js"
import * as echarts from "echarts"
import CpuInfo from "./cpuInfo.js"
const cpuinfo=CpuInfo();
const cpuInfo=ref({})
const UsingCpu=ref(0);
const chart=ref(null)


onMounted(async ()=>{
  cpuInfo.value=await GetCpuInfo().then(res=>res);
  UsingCpu.value=await GetUsingCpuInfo().then(res=>res);
  console.log( cpuInfo.value,UsingCpu.value)
  option.value.title[1].text= UsingCpu.value.toFixed(2)+"";
  chart.value.EchartsInit();
//   setInterval(async ()=>{
//     UsingCpu.value=await GetUsingCpuInfo().then(res=>res);
//     option.value.title[1].text= UsingCpu.value.toFixed(2)+"";
//     option.value.series[0].data=[ UsingCpu.value.toFixed(2)]
//     chart.value.EchartsInit();
//  console.log(11111)
//   },2000)
 
})
</script>
<style lang='scss' scoped>

</style>