import { ref } from "vue";

const Power=()=>{
  const PowerInfo=ref([
    {label:"电源型号",val:"",key:"Name"},
    {label:"功率",val:"",key:"Power"},
    {label:"参考价格",val:"",key:"Price"},
    {label:"电源类型",val:"",key:"Type"},
    {label:"接口类型",val:"",key:"InterFaceType"},
    {label:"CPU型号",val:"",key:"Name"},
    {label:"CPU型号",val:"",key:"Name"},
    {label:"CPU型号",val:"",key:"Name"},
    {label:"CPU型号",val:"",key:"Name"},
  ])
  return {
    PowerInfo
  }
}
export default Power;