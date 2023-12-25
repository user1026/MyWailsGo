<template>
  <el-row :gutter="4">
    <el-col :span="12">
      <el-form :model="FormData" label-width="120px">
        <el-form-item label="CPU">
          <CPU />
        </el-form-item>
        <el-form-item label="主板">
          <MainBoard />
        </el-form-item>
        <el-form-item label="内存">
          <RAM />
        </el-form-item>
        <el-form-item label="显卡">
          <GPU />
        </el-form-item>
        <el-form-item label="散热器">
          <Radiator />
        </el-form-item>
        <el-form-item label="电源">
          <Power />
        </el-form-item>
        <el-form-item label="机箱">
          <Chassis />
        </el-form-item>
        <el-form-item label="固态硬盘">
          <SSD />
        </el-form-item>
        <el-form-item label="机械硬盘">
          <HDD />
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="12">
      <el-button @click="Export" type="success">导出</el-button>
    </el-col>
  </el-row>
</template>
<script setup>
  import {
    onMounted,
    ref,
    reactive,
    onDeactivated
  } from "vue";
  import {
    post,
    get
  } from "@/utils/http.js";
  import CPU from "./CPU/index.vue";
  import GPU from "./GPU/index.vue";
  import RAM from "./RAM/index.vue";
  import Power from "./Power/index.vue";
  import MainBoard from "./MainBoard/index.vue";
  import SSD from "./SSD/index.vue";
  import Chassis from "./Chassis/index.vue";
  import Radiator from "./Radiator/index.vue";
  import HDD from "./HDD/index.vue"
  import FormData from "./index.js";
  import Global from "../../Global/index.js"
  import { ExportFile } from "../../../wailsjs/go/main/App.js"
  import { ElMessage } from 'element-plus'
  // const RadiatorList = ref([]);
  // const CPUList = ref([]);
  // const GPUList = ref([]);
  // const RAMList = ref([]);
  // const PowerList = ref([]);
  // const HDDList = ref([]);
  // const SSDList = ref([]);
  // const ChassisList = ref([]);
  // const MainBoardList = ref([]);
  // const CPUInfo = ref({});
  onMounted(async () => {


  })
  onDeactivated(() => {

  })
  const Export=async ()=>{
         let flag=await ExportFile("txt","1234567890")
         if(flag){
          ElMessage({
            type:"success",
            message:"导出成功"
          })
         }
  }
</script>
<style lang="scss" scoped>

</style>