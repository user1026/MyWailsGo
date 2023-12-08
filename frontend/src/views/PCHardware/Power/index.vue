<template>
  <el-select v-model="FormData.Power" clearable placeholder="请选择电源">
    <el-option v-for="item in PowerList" :key="item.value" :label="item.label" :value="item.value" />
  </el-select>
  <template  v-if="FormData.Power.Name">
    <el-tooltip class="box-item" effect="light" placement="right">
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
const PowerList = ref([]);
  const PowerInfo=ref([
    {label:"电源型号",val:"",key:"Name"},
    {label:"功率",val:"",key:"Power"},
    {label:"参考价格",val:"",key:"Price"},
    {label:"电源类型",val:"",key:"Type"},
    {label:"接口类型",val:"",key:"InterFaceType"},
    {label:"CPU型号",val:"",key:"Name"},
    {label:"CPU型号",val:"",key:"Name"},
    {label:"CPU型号",val:"",key:"Name"},
    {label:"CPU型号",val:"",key:"Name"},
  ])
</script>
<style scoped>

</style>