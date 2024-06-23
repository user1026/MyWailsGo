import { toRefs } from "vue";
export const CheckAll = (Form) => {
  let message = [];
  let { CPU, GPU, RAM, Power, Chassis, MainBoard, HDD, SSD, Radiator } = toRefs(Form);
  if (CheckNull(CPU.value.Type) && CheckNull(MainBoard.value.Type)) {
    if (CPU.Type !== MainBoard.Type) {
      message.push(
        `CPU与主板平台不符,CPU为${CPU.Type},主板为${MainBoard.Type}`
      );
    }
  }
  if (CheckNull(CPU.CPUInterface) && CheckNull(MainBoard.CPUInterface)) {
    if (CPU.CPUInterface !== MainBoard.CPUInterface) {
      message.push("CPU接口与主板CPU接口不匹配");
    }
  }
  if (
    CheckNull(Chassis.value.RadiatorHeight) &&
    CheckNull(Radiator.value.Height)
  ) {
    if (StrToNum(Chassis.value.RadiatorHeight) <StrToNum(Radiator.value.Height)) {
      message.push("散热器高度超过机箱支持的散热器高度");
    }
  }
  if (
    CheckNull(Chassis.value.SupportMainBoardTypes) &&
    CheckNull(MainBoard.value.Type)
  ) {
    if (Chassis.value.SupportMainBoardTypes.indexOf(MainBoard.value.Type) < 0) {
      message.push("机箱不支持该类型的主板");
    }
  }
  if(CheckNull(RAM.value.Type)&&CheckNull(MainBoard.value.SupportRAMType)){
    if(RAM.value.Type!==MainBoard.value.SupportRAMType){
      message.push(`主板不支持该类型的内存条,所选主板支持${MainBoard.value.SupportRAMType}类型，所选内存条为${RAM.value.Type}类型`)
    }
  }
  if(CheckNull(Chassis.value.SupportPowerTypes)&&CheckNull(Power.value.Type)){
    if(Chassis.value.SupportPowerTypes.indexOf(Power.value.Type)<0){
        message.push(`所选机箱不支持所选电源，所选机箱支持的电源类型为${Chassis.value.SupportPowerTypes},所选电源类型为${Power.value.Type}`)
    }
  }
  if(CheckNull(Chassis.value.SupportGPULength)&&CheckNull(GPU.value.Length)){
    if(StrToNum(Chassis.value.SupportGPULength)<StrToNum(GPU.value.Length)){
        message.push(`所选显卡长度超过了所选机箱支持的显卡长度`)
    }
  }
  // if(CheckNull()&&CheckNull()){
  //   if(1){

  //   } 
  // }
  // if(CheckNull()&&CheckNull()){
  //   if(1){

  //   }
  // }
};

const CheckNull = (x) => {
  if (x === null || x === undefined || x === "") {
    return false;
  }
  return true;
};

const StrToNum=(s)=>{
      let n=parseInt(s);
      if(Number.isNaN(n)){
        return 0;
      }else{
        return n;
      }
}
