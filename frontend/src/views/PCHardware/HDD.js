import { ref } from "vue";

const HDD=()=>{
  const HDDInfo=ref([
    {label:"硬盘名",val:"",key:"Name"},
    {label:"容量",val:"",key:"Total"},
    {label:"转速",val:"",key:"Name"},
    {label:"参考价格",val:"",key:"Price"},
    {label:"上市日期",val:"",key:"CreateTime"},
    {label:"硬盘类型",val:"",key:"Type"},
    {label:"图片",val:"",key:"ImgUrl"},
  ])
  return {
    HDDInfo
  }
}
export default HDD;