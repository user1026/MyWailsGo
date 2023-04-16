<template>
    <div>
        欢迎使用客户端 
        <el-button @click="add">{{ count }}</el-button>
    </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
const count=ref(1)
const add=()=>{
window.runtime.openFile().then(list=>{
    count.value=list
console.log(list)
})
}
</script>
<style lang='scss' scoped></style>