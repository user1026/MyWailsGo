<script setup>
  import { reactive, ref,onMounted } from 'vue'
  import { SetExportFileUrl,SetExportFileType } from '../../../wailsjs/go/main/App'
  import Global from "../../Global/index.js"
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
    if(filename!==""){
      FormData.value.ExportFileUrl = filename
      Global.ExportFileUrl=filename
    }
  }
  const FileType=async ()=>{
  let flag= await SetExportFileType(FormData.value.ExportFileType)
  if(flag==false){

  }
  }
 onMounted(async ()=>{
  FormData.value.ExportFileType=Global.ExportFileType
  FormData.value.ExportFileUrl=Global.ExportFileUrl
 })
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
          <el-select v-model="FormData.ExportFileType" @change="FileType" clearable placeholder="请选择">
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