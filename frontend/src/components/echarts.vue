<template>
    <!-- :style="{'width':props.width,'height':props.height}" -->
    <div id="echarts" >

    </div>
</template>
<script setup>
    import { ref,reactive,onMounted,watch } from 'vue'
    import   * as echarts from "echarts"
    const props=defineProps(["option","width","height"])
    let echartsDiv;
    onMounted(()=>{
       // echarts.init(document.getElementById("main")).setOption(props.option)
    })
    const EchartsInit=()=>{
        echartsDiv=echarts.init(document.getElementById("echarts")).setOption(props.option)
    }
    const setOption=(option)=>{
        echartsDiv.setOption(option);
    }
    defineExpose({
        EchartsInit,
        setOption
    })
    // watch([()=>props.width,props.height],([width,height],[newV,oldV])=>{
    //     echarts.init(document.getElementById("echarts")).setOption(props.option)
    // })
   
</script>
<style lang='scss' scoped>
  
</style>