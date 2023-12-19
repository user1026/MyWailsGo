<template>
  <template v-for="(v,i) in HDDvalues" :key="v.key">
    <el-select  @change="changeS" v-model="FormData.HDD[i]" clearable placeholder="请选择机械硬盘">
      <el-option v-for="item in HDDList" :key="item.value" :label="item.label" :value="item.value" />
    </el-select>
    <el-button v-if="i==0" @click="addHDD">
      +
    </el-button>
    <el-button @click="delHDD(i)">
      -
    </el-button>
    <template v-if="FormData.HDD.Name">
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
</template>
<script setup>
import {ref} from "vue";
import { ElMessage } from 'element-plus';
import FormData from "../index.js";
import Descriptions from "@/components/Descriptions.vue";
const HDDList = ref([{
  label:"1",
  value:1,
},{
  label:"2",
  value:2,
},{
  label:"3",
  value:3,
}]);
let n=0;
const HDDvalues=ref([{
  key:0,
  value:""
}])
  const HDDInfo=ref([
    {label:"硬盘名",val:"",key:"Name"},
    {label:"容量",val:"",key:"Total"},
    {label:"转速",val:"",key:"Name"},
    {label:"参考价格",val:"",key:"Price"},
    {label:"上市日期",val:"",key:"CreateTime"},
    {label:"硬盘类型",val:"",key:"Type"},
    {label:"图片",val:"",key:"ImgUrl"},
  ]);
  const changeS=()=>{
    console.log(FormData.value.HDD);
  }
  const addHDD=()=>{
    if(n==3){
      ElMessage({
    message: '目前只限制4个',
    type: 'warning',
  })
      return ;
    }
    HDDvalues.value.push({key:n++,value:""});
  }
  const delHDD=(i)=>{
    if(n==0){
        return 
    }
    HDDvalues.value.splice(i,1);
    FormData.value.HDD.splice(i,1);
    n--;
  }
</script>
<style scoped>

</style>