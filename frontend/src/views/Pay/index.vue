<template>
 
  <div class="descriptions">
  
    <el-descriptions class="margin-top" title="购买" :column="2" size="large" border>
      <el-descriptions-item>
        <template #label>
          <el-tag size="large">贴吧整理的奸商店铺</el-tag>
        </template>
        <el-image style="width: 500px; height: 300px" fit="fill" :src="imgUrl" />
      </el-descriptions-item>
    </el-descriptions>
    <div class="bottom">

    </div>
  </div>
</template>
<script setup>
import {ref} from "vue"
import imgUrl from "@/assets/images/danger.jpg"
</script>
<style scoped>
  .descriptions {}

  .bottom {
    height: 90px;
  }
</style>