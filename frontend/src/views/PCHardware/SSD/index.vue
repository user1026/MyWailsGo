<template>
  <template v-for="(v,i) in SSDvalues">
    <el-select v-model="FormData.SSD[i]" clearable placeholder="请选择固态硬盘">
      <el-option v-for="item in SSDList" :key="item.value" :label="item.label" :value="item.value" />
    </el-select>
    <el-button v-if="i==0" @click="addSSD">
      +
    </el-button>
    <el-button @click="delSSD(i)">
      -
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
</template>
<script setup>
import {ref} from "vue";
import FormData from "../index.js";
import Descriptions from "@/components/Descriptions.vue";
const SSDList = ref([]);
const SSDvalues=ref([{}]);
let n=0;
const SSDInfo=ref([
  {label:"固态硬盘名",val:"",key:"Name"},
  {label:"容量",val:"",key:"Total"},
  {label:"颗粒类型",val:"",key:"Type"},
  {label:"规格",val:"",key:""},
  {label:"顺序读写",val:"",key:""},
  {label:"随机读写",val:"",key:""},
])
const addSSD=()=>{
    if(n==3){
      ElMessage({
    message: '目前只限制4个',
    type: 'warning',
  })
      return;
    }
    SSDvalues.value.push({key:n++,value:""});
  }
  const delSSD=(i)=>{
    if(n==0){
        return ;
    }
    SSDvalues.value.splice(i,1);
    FormData.value.SSD.splice(i,1);
    n--;
  }
</script>
<style scoped>

</style>