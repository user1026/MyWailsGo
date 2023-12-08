<template>
  <el-select v-model="FormData.CPU" clearable value-key="Name" placeholder="请选择CPU">
    <el-option v-for="(item,i) in CPUList" :key="i" :label="item.Name" :value="item" />
  </el-select>
  <template v-if="FormData.CPU.Name">
    <el-tooltip  class="box-item" effect="light" placement="right">
      <template #content>
        <Descriptions :List="CPUInfo" :Info="FormData.CPU" />
      </template>
      <el-icon size="25" color="#a3b4b5">
        <InfoFilled />
      </el-icon>
    </el-tooltip>
  </template>
</template>
<script setup>
import {ref} from "vue"
import FormData from "../index.js"
import Descriptions from "@/components/Descriptions.vue"
const CPUList = ref([]);
const CPUMap={
  Name:"CPU型号",        
	Cores:"核心数",         
	CPUInterface :"CPU接口型号", 
	Threads:"线程数",       
	GHZ:"基准频率~加速频率",           
	CPUPower:"功耗",      
	CreateTime:"上市时间",    
	Price:"首发价格",
  MainGHZ:"大核（基准频率~加速频率）",
  ACCGHZ:"小核（基准频率~加速频率）"          
}
const GetCPUData=ref({})
const CPUInfo=ref([
  {label:"CPU型号",val:"",key:"Name"},
  {label:"核心数",val:"",key:"Cores"},
  {label:"CPU接口型号",val:"",key:"CPUInterface"},
  {label:"线程数",val:"",key:"Threads"},
  {label:"功耗",val:"",key:"CPUPower"},
  {label:"上市时间",val:"",key:"CreateTime"},
  {label:"首发价格",val:"",key:"Price"},
  {label:"大核（基准频率~加速频率）",val:"",key:"MainGHZ"},
  {label:"小核（基准频率~加速频率）",val:"",key:"AccGHZ"},
  {label:"基准频率~加速频率",val:"",key:"GHZ"},
])
</script>
<style scoped>

</style>