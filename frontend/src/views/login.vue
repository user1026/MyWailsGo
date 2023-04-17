<template>
    <div>
        欢迎使用客户端 
        <el-button @click="add">{{ count }}</el-button>
    </div>
</template>
<script setup>
import { ref, reactive, onMounted ,watch} from 'vue'
const count=ref(1)
const add=()=>{
    console.log(window.backend)
window.runtime.openFile.then(list=>{
    count.value=list
console.log(list)
})
}
window.runtime.EventsOn("openFile",(data)=>{
    count.value=data;
    console.log(data);
})
watch(count,(n,o)=>{

    console.log(n,o)
})
</script>
<style lang='scss' scoped></style>