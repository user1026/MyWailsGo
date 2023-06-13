<template>
<div>
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
      <el-button type="primary" @click="select">查</el-button>
    </el-form-item>
    </el-form>
</div>
<div>
</div>
</template>
<script setup>
import {onMounted,ref} from "vue";
import {post,get} from "@/utils/http.js";
const FormData=ref({
  qq:"",
  phone:"",
  wbId:"",
})
const select=async ()=>{
 FormData.phone=phone();
}
const phone=async ()=>{
  console.log(FormData.qq,FormData.value.qq,"----")
 return await post({url:"/qqapi",data:{qq:FormData.value.qq}}).then(res=>{
    return res.phone
  })
}
onMounted(()=>{
  get({url:"/ip"}).then(res=>{
   console.log(res)
  })
})
</script>
<style lang="scss" scoped>

</style>