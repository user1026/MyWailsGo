<template>
  <el-select v-model="FormData.RAM" clearable placeholder="请选择内存">
    <el-option v-for="item in RAMList" :key="item.value" :label="item.label" :value="item.value" />
  </el-select>
  <template v-if="FormData.RAM.Name">
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
const RAMList = ref([]);
const RAMInfo=ref([
  {label:"内存型号",val:"",key:"Name"},
  {label:"内存容量",val:"",key:"Total"},
  {label:"时序",val:"",key:"CL"},
  {label:"频率",val:"",key:"MHZ"},
  {label:"颗粒",val:"",key:"Type"},
  {label:"首发价格",val:"",key:"Price"},
  {label:"图片",val:"",key:"ImgUrl"},
  {label:"上市日期",val:"",key:"CreateTime"},
])
</script>
<style scoped>

</style>