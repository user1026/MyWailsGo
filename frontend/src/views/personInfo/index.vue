<template>

  <el-row :gutter="4">
    <el-col :span="12">
      <el-form :model="FormData" label-width="120px">
        <el-form-item label="qq">
          <el-input v-model="FormData.qq" clear></el-input>
        </el-form-item>
        <el-form-item label="手机号">
          <el-input v-model="FormData.phone" clear></el-input>
        </el-form-item>
        <el-form-item label="微博ID">
          <el-input v-model="FormData.wbId" clear></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="query">查</el-button>
          <el-button type="primary" @click="clear">清除</el-button>
        </el-form-item>
        </el-form>
    </el-col>
    <el-col :span="12">
      <el-descriptions title="信息">
        <el-descriptions-item label="手机号属地">{{Info.phonediqu}}</el-descriptions-item>
      </el-descriptions>
    </el-col>
  </el-row>
</template>
<script setup>
import {onMounted,ref} from "vue";
import {post,get} from "@/utils/http.js";
const FormData=ref({
  qq:"",
  phone:"",
  wbId:"",
})
const Info=ref({
  qq:"",
  phone:"",
  wbId:"",
  phonediqu:"",
})
const query=()=>{
if(FormData.value.qq){
QueryByQQ();
}else if(FormData.value.phone){
QueryByPhone();
}else if(FormData.value.wbId){
QueryByWb();
}
}
const clear=()=>{
  FormData.clear()
   FormData=ref({
  qq:"",
  phone:"",
  wbId:"",
})
 Info=ref({
  qq:"",
  phone:"",
  wbId:"",
  phonediqu:"",
})
}
const QueryByQQ= async()=>{
let x=await qqToPhone();
 FormData.value.phone=x.phone
 Info.value.phone=x.phone;
 Info.value.phonediqu=x.phonediqu;
 let a=await phoneToWb();
 FormData.value.wbId=a.id;
}
const QueryByPhone=async ()=>{
  let x=await phoneToQQ();
 FormData.value.qq=x.qq
 Info.value.phone=x.phone;
 Info.value.phonediqu=x.phonediqu;
 let a=await phoneToWb();
 FormData.value.wbId=a.id;
}
const QueryByWb=async ()=>{
  let x=await WbToPhone();
 FormData.value.phone=x.phone;
 Info.value.phone=x.phone;
 Info.value.phonediqu=x.phonediqu;
 let a=await phoneToQQ();
 FormData.value.qq=a.qq;
}
const qqToPhone=async ()=>{
 return await get("/cc/qqapi",{qq:FormData.value.qq})
}
const phoneToWb=async ()=>{
  return await get("/cc/wbphone",{phone:FormData.value.phone})
}
const phoneToQQ=async ()=>{
  return await get("/cc/qqphone",{phone:FormData.value.phone})
}
const WbToPhone=async ()=>{
  return await get("/cc/wbapi",{id:FormData.value.wbId})
}
onMounted(()=>{
  // get({url:"/ip"}).then(res=>{
  //  console.log(res)
  // })
})
</script>
<style lang="scss" scoped>

</style>