<template>
  <el-select v-model="FormData.SSD" clearable placeholder="请选择固态硬盘">
    <el-option v-for="item in SSDList" :key="item.value" :label="item.label" :value="item.value" />
  </el-select>
  <el-button @click="add">
    +
  </el-button>
  <template v-if="FormData.SSD.Name">
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
const SSDList = ref([]);
const SSDInfo=ref([
  {label:"固态硬盘名",val:"",key:"Name"},
  {label:"容量",val:"",key:"Total"},
  {label:"颗粒类型",val:"",key:"Type"},
  {label:"规格",val:"",key:""},
  {label:"顺序读写",val:"",key:""},
  {label:"随机读写",val:"",key:""},
])
const add=()=>{}
</script>
<style scoped>

</style>