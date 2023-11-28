import { ref } from "vue";

const GPU=()=>{
const GPUInfo=ref([
  {label:"图片",val:"",key:"ImgUrl"},
  {label:"显卡型号",val:"",key:"Name"},
  {label:"流处理器数",val:"",key:"LGPUNumber"},
  {label:"核心型号",val:"",key:"Name"},
  {label:"功耗",val:"",key:"GPUPower"},
  {label:"位宽",val:"",key:"Bit"},
  {label:"显存容量",val:"",key:"Name"},
  {label:"显存频率",val:"",key:"RAMGhz"},
  {label:"厂商",val:"",key:"OEM"},
  {label:"显存型号",val:"",key:"RAMType"},
  {label:"长度",val:"",key:"Length"},
  {label:"宽度",val:"",key:"Width"},
  {label:"高度",val:"",key:"Height"},
  {label:"首发价格",val:"",key:"Price"},
  {label:"上市日期",val:"",key:"CreateTime"},
])
return {
  GPUInfo
}
}
export default GPU;