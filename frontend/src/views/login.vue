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

}
window.runtime.EventsOn("openFile",(data)=>{
    count.value=JSON.stringify(data)
    console.log(data);
})

</script>
<style lang='scss' scoped></style>