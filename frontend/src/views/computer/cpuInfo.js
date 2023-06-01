import {ref} from "vue"
import   * as echarts from "echarts"
import {GetCpuInfo,GetUsingCpuInfo,GetMemInfo} from "../../../wailsjs/go/main/App.js"
let cpuInfo=async ()=>{
  const UsingCpu=ref("0")
  const option =ref({
    title: [
      {
        text: "cpu使用率",
        x: "center",
        top: "55%",
        textStyle: {
          color: "#FFFFFF",
          fontSize: 16,
          fontWeight: "100",
        },
      },
      {
        text: "0",
        x: "center",
        y: "center",
        textStyle: {
        //  fontSize: "60",
          color: "#FFFFFF",
          fontFamily: "DINAlternate-Bold, DINAlternate",
          foontWeight: "600",
        },
      },
    ],
    backgroundColor: "#111",
    polar: {
      radius: ["42%", "52%"],
      center: ["50%", "50%"],
    },
    angleAxis: {
      max: 100,
      show: false,
    },
    radiusAxis: {
      type: "category",
      show: true,
      axisLabel: {
        show: false,
      },
      axisLine: {
        show: false,
      },
      axisTick: {
        show: false,
      },
    },
    series: [
      {
        name: "",
        type: "bar",
        roundCap: true,
        barWidth: 30,
        showBackground: true,
        backgroundStyle: {
          color: "rgba(66, 66, 66, .3)",
        },
        data: [60],
        coordinateSystem: "polar",
        itemStyle: {
          normal: {
            color: new echarts.graphic.LinearGradient(0, 1, 0, 0, [
              {
                offset: 0,
                color: "#16CEB9",
              },
              {
                offset: 1,
                color: "#6648FF",
              },
            ]),
          },
        },
      },
      {
        name: "",
        type: "pie",
        startAngle: 80,
        radius: ["56%"],
        hoverAnimation: false,
        center: ["50%", "50%"],
        itemStyle: {
          color: "rgba(66, 66, 66, .1)",
          borderWidth: 1,
          borderColor: "#5269EE",
        },
        data: [100],
      },
      {
        name: "",
        type: "pie",
        startAngle: 80,
        radius: ["38%"],
        hoverAnimation: false,
        center: ["50%", "50%"],
        itemStyle: {
          color: "rgba(66, 66, 66, .1)",
          borderWidth: 1,
          borderColor: "#5269EE",
        },
        data: [100],
      },
    ],
  });
  cpuInfo.value=await GetCpuInfo().then(res=>res);
  UsingCpu.value=await GetUsingCpuInfo().then(res=>res);
  option.value.title[1].text=UsingCpu.value.toFixed(2)+"";
  return {
    cpuInfo,
    option,
  }
}



export default  cpuInfo