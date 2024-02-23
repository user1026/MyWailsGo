<template>
  <el-select v-model="FormData.Chassis" clearable placeholder="请选择机箱">
    <el-option v-for="item in ChassisList" :key="item.value" :label="item.label" :value="item.value" />
  </el-select>
  <template v-if="FormData.Chassis.Name">
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
  import Descriptions from "@/components/Descriptions.vue"
  import {ref} from "vue"
  import FormData from "../index.js"
  const ChassisList = ref([]);
</script>
<style scoped>

</style>