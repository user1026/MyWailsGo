<template>
    <div id="main">

    </div>
</template>
    <script setup>
    import { ref,reactive,onMounted } from 'vue'
    import   * as echarts from "echarts"
    const props=defineProps(["option"])
    onMounted(()=>{
       // echarts.init(document.getElementById("main")).setOption(props.option)
    })
    const EchartsInit=()=>{
        echarts.init(document.getElementById("main")).setOption(props.option)
    }
    defineExpose({
        EchartsInit
    })
    </script>
    <style lang='scss' scoped>
    #main{
        width: 100px;
        height: 100px;
    }
    </style>