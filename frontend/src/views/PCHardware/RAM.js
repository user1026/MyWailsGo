import { ref } from "vue";

const RAM =()=>{
const RAMInfo=ref([
  {label:"内存型号",val:"",key:"Name"},
  {label:"内存容量",val:"",key:"Total"},
  {label:"时序",val:"",key:"CL"},
  {label:"频率",val:"",key:"MHZ"},
  {label:"颗粒",val:"",key:"Type"},
  {label:"首发价格",val:"",key:"Price"},
  {label:"图片",val:"",key:"ImgUrl"},
  {label:"上市日期",val:"",key:"CreateTime"},
])
return {
  RAMInfo
}
}
export default RAM;