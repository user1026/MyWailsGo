import { ref } from "vue";

const SSD=()=>{
const SSDInfo=ref([
  {label:"固态硬盘名",val:"",key:"Name"},
  {label:"容量",val:"",key:"Total"},
  {label:"颗粒类型",val:"",key:"Type"},
  {label:"规格",val:"",key:""},
  {label:"顺序读写",val:"",key:""},
  {label:"随机读写",val:"",key:""},
])
return {
  SSDInfo
}
}

export default SSD;