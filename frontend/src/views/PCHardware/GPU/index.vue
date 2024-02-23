<template>
  <el-select v-model="FormData.GPU" clearable placeholder="请选择显卡">
    <el-option v-for="item in GPUList" :key="item.value" :label="item.label" :value="item.value" />
  </el-select>
  <template v-if="FormData.GPU.Name">
    <el-tooltip  class="box-item" effect="light" placement="right">
      <template #content>
        <Descriptions />
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
const GPUList = ref([]);
const GPUInfo=ref([
  {label:"图片",val:"",key:"ImgUrl"},
  {label:"显卡型号",val:"",key:"Name"},
  {label:"流处理器数",val:"",key:"LGPUNumber"},
  {label:"核心型号",val:"",key:"Name"},
  {label:"功耗",val:"",key:"GPUPower"},
  {label:"位宽",val:"",key:"Bit"},
  {label:"显存容量",val:"",key:"Name"},
  {label:"显存频率",val:"",key:"RAMGhz"},
  {label:"厂商",val:"",key:"OEM"},
  {label:"显存型号",val:"",key:"RAMType"},
  {label:"长度",val:"",key:"Length"},
  {label:"宽度",val:"",key:"Width"},
  {label:"高度",val:"",key:"Height"},
  {label:"首发价格",val:"",key:"Price"},
  {label:"上市日期",val:"",key:"CreateTime"},
])
</script>
<style scoped>

</style>