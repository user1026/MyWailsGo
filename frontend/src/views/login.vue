<template>
     <el-dialog
    v-model="isShow"
    width="30%"
    center
    :show-close="false"
    :close-on-click-modal="false"
  >
    <span class="txt">欢迎使用本客户端</span>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="to">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>
<script setup>
import { ref, reactive, onMounted ,watch} from 'vue'
import {useRouter} from "vue-router"
const count=ref(1);
const isShow=ref(true);
const router=useRouter();
const to=()=>{
router.push("/index")
}
window.runtime.EventsOn("openFile",(data)=>{
    count.value=JSON.stringify(data)
    console.log(data);
})

</script>
<style lang='scss' scoped>
.txt{
    font: normal 40px/40px "微软雅黑";
    text-align: center;
    width: 100%;
    display: inline-block;
}
</style>