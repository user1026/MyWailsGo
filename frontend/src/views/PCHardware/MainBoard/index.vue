<template>
  <el-select v-model="FormData.MainBoard" clearable placeholder="请选择主板">
    <el-option v-for="item in MainBoardList" :key="item.value" :label="item.label" :value="item.value" />
  </el-select>
  <template v-if="FormData.MainBoard.Name">
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
const MainBoardList=ref();
</script>
<style scoped>

</style>