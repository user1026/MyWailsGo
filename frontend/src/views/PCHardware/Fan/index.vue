<template>
  <el-select v-model="FormData.Fan" clearable placeholder="请选择内存">
    <el-option v-for="item in FanList" :key="item.value" :label="item.label" :value="item.value" />
  </el-select>
  <el-input  v-model="FormData.Fan.Num"></el-input>
  <template v-if="FormData.Fan.Name">
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
const FanList = ref([]);
const FanInfo=ref([
  {label:"风扇名",val:"",key:"Name"},
  {label:"转速",val:"",key:"Total"},
  {label:"神光同步",val:"",key:"CL"},
  {label:"支持RGB",val:"",key:"SupportRGB"},
  {label:"规格",val:"",key:"Type"},
  {label:"首发价格",val:"",key:"Price"},
  {label:"图片",val:"",key:"ImgUrl"},
  {label:"上市日期",val:"",key:"CreateTime"},
])
</script>
<style scoped>
.el-input{
  width: 40px;
}
</style>