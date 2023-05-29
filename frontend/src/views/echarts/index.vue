<template>
    <div id="main">

    </div>
    </template>
    <script setup>
    import { ref,reactive,onMounted } from 'vue'
    import  * as echarts from 'echarts'
    const props=defineProps(["option"])
    onMounted(()=>{
        echarts.init(document.getElementById("main")).setOption(props.option)
    })
    const EchartsInit=()=>{
         
    }
    </script>
    <style lang='scss' scoped>
    #id{
        width: 400px;
        height: 400px;
    }
    </style>