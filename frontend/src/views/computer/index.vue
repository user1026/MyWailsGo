<template>
  <div>
    <el-descriptions title="Customized style list" :column="3" border>
      <el-descriptions-item
        label="Username"
        label-align="right"
        align="center"
        label-class-name="my-label"
        class-name="my-content"
        width="150px"
        >kooriookami</el-descriptions-item
      >
      <el-descriptions-item label="Telephone" label-align="right" align="center"
        >18100000000</el-descriptions-item
      >
      <el-descriptions-item label="Place" label-align="right" align="center"
        >Suzhou</el-descriptions-item
      >
     
    </el-descriptions>
  </div>
    <div>
        <tu :option="option" ref="chart"></tu>
    </div>
</template>
<script setup>
import { ref,reactive,onMounted,onBeforeMount} from 'vue'
import tu from "@/views/echarts/index.vue"
//import {GetCpuInfo,GetUsingCpuInfo} from "../../../wailsjs/go/main/App.js"
import CpuInfo from "./cpuInfo.js"
const UsingCpu=ref("0")
const chart=ref(null)
const option=ref(null)

  onBeforeMount(async ()=>{
  const cpuinfo=await CpuInfo();
  option.value=cpuinfo.option.value;
  console.log(option,"option")
})

onMounted(()=>{
  try {
    chart.value.EchartsInit();
  } catch (error) {
    console.log(error)
  }
  
})
</script>
<style lang='scss' scoped>

</style>