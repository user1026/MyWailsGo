<template>
  <el-row :gutter="4">
    <el-col :span="12">
      <el-form :model="FormData" label-width="120px">
        <el-form-item label="CPU">
          <el-select v-model="FormData.CPU" clearable value-key="Name" placeholder="请选择CPU">
            <el-option v-for="(item,i) in CPUList" :key="i" :label="item.Name" :value="item" />
          </el-select>
          <el-tooltip class="box-item" effect="light" placement="right">
            <template #content>
              <Descriptions :List="CPUInfo" :Info="FormData.CPU" />
            </template>
            <el-icon size="25" color="#a3b4b5">
              <InfoFilled />
            </el-icon>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="主板">
          <el-select v-model="FormData.MainBoard" clearable placeholder="请选择主板">
            <el-option v-for="item in MainBoardList" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
          <el-tooltip class="box-item" effect="light" placement="right">
            <template #content>
              <Descriptions />
            </template>
            <el-icon size="25" color="#a3b4b5">
              <InfoFilled />
            </el-icon>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="内存">
          <el-select v-model="FormData.RAM" clearable placeholder="请选择内存">
            <el-option v-for="item in RAMList" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
          <el-tooltip class="box-item" effect="light" placement="right">
            <template #content>
              <Descriptions />
            </template>
            <el-icon size="25" color="#a3b4b5">
              <InfoFilled />
            </el-icon>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="显卡">
          <el-select v-model="FormData.GPU" clearable placeholder="请选择显卡">
            <el-option v-for="item in GPUList" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
          <el-tooltip class="box-item" effect="light" placement="right">
            <template #content>
              <Descriptions />
            </template>
            <el-icon size="25" color="#a3b4b5">
              <InfoFilled />
            </el-icon>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="固态硬盘">
          <el-select v-model="FormData.SSD" clearable placeholder="请选择固态硬盘">
            <el-option v-for="item in SSDList" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
          <el-tooltip class="box-item" effect="light" placement="right">
            <template #content>
              <Descriptions />
            </template>
            <el-icon size="25" color="#a3b4b5">
              <InfoFilled />
            </el-icon>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="散热器">
          <el-select v-model="FormData.Radiator" clearable placeholder="请选择散热器">
            <el-option v-for="item in RadiatorList" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
          <el-tooltip class="box-item" effect="light" placement="right">
            <template #content>
              <Descriptions />
            </template>
            <el-icon size="25" color="#a3b4b5">
              <InfoFilled />
            </el-icon>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="电源">
          <el-select v-model="FormData.Power" clearable placeholder="请选择电源">
            <el-option v-for="item in PowerList" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
          <el-tooltip class="box-item" effect="light" placement="right">
            <template #content>
              <Descriptions />
            </template>
            <el-icon size="25" color="#a3b4b5">
              <InfoFilled />
            </el-icon>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="机箱">
          <el-select v-model="FormData.Chassis" clearable placeholder="请选择机箱">
            <el-option v-for="item in ChassisList" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
          <el-tooltip class="box-item" effect="light" placement="right">
            <template #content>
              <Descriptions />
            </template>
            <el-icon size="25" color="#a3b4b5">
              <InfoFilled />
            </el-icon>
          </el-tooltip>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="12">
      <el-button type="success">导出</el-button>
    </el-col>
  </el-row>
</template>
<script setup>
  import {
    onMounted,
    ref
  } from "vue";
  import {
    post,
    get
  } from "@/utils/http.js";
  import Descriptions from "@/components/Descriptions.vue"
  import CPU from "./CPU.js"
  import {
    GetCpuJSONData
  } from "../../../wailsjs/go/main/App.js"
  const cpu = CPU();
  const FormData = ref({
    CPU: {},
    GPU: {},
    RAM: {},
    Power: {},
    Chassis: {},
    MainBoard: {},
    HDD: {},
    SSD: {},
    Radiator: {}
  })
  const RadiatorList = ref([]);
  const CPUList = ref([]);
  const GPUList = ref([]);
  const RAMList = ref([]);
  const PowerList = ref([]);
  const HDDList = ref([]);
  const SSDList = ref([]);
  const ChassisList = ref([]);
  const MainBoardList = ref([]);
  const CPUInfo = ref({});
  onMounted(async () => {
    CPUInfo.value = cpu.CPUInfo.value;
    console.log(CPUInfo, "-----------")
    CPUList.value = await GetCpuJSONData();

  })
</script>
<style lang="scss" scoped>

</style>