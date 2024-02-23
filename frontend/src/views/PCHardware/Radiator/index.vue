<template>
  <el-select v-model="FormData.Radiator" clearable placeholder="请选择散热器">
    <el-option v-for="item in RadiatorList" :key="item.value" :label="item.label" :value="item.value" />
  </el-select>
  <template  v-if="FormData.Radiator.Name">
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
const RadiatorList = ref([]);
</script>
<style scoped>

</style>