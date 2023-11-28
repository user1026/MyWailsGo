<script setup>
  import { reactive, ref } from 'vue'
  import { SetExportFileUrl } from '../../../wailsjs/go/main/App'
  const FormData = ref({
    ExportFileType: "",
    ExportFileUrl: ""
  })
  const ExportFileTypeList = ref([{
    label: "txt",
    value: "txt",
  }, {
    label: "excel",
    value: "excel",
  }])
  const FileUrl = async () => {
    let filename = await SetExportFileUrl();
    FormData.value.ExportFileUrl = filename
  }
</script>

<template>
  <el-form :model="FormData" label-width="120px">
    <el-row :gutter="1">
      <el-col :span="8">
        <el-form-item label="导出位置">
          <el-input v-model="FormData.ExportFileUrl"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="1">
        <el-button type="primary" @click="FileUrl">选择</el-button>
      </el-col>
    </el-row>
    <el-row :gutter="1">
      <el-col :span="8">
        <el-form-item label="导出类型">
          <el-select v-model="FormData.ExportFileType" clearable placeholder="请选择">
            <el-option v-for="item in ExportFileTypeList" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row :gutter="1">
      <el-form-item label="版本">
        0.0.1
      </el-form-item>
    </el-row>
    <el-row :gutter="1">
      <el-form-item label="更新时间">
        2023.10.16
      </el-form-item>
    </el-row>
  </el-form>
</template>

<style scoped>

</style>