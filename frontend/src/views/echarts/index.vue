<template>
    <div id="computer">

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
        echarts.init(document.getElementById("computer")).setOption(props.option)
    }
    defineExpose({
        EchartsInit
    })
    </script>
    <style lang='scss' scoped>
    #computer{
        width: 300px;
        height: 300px;
    }
    </style>