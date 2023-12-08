package Computer

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"io/ioutil"
	"time"
)

type Computer struct {
}
type CPUInfo struct {
	Name    string
	Cores   int
	Threads int
	GHZ     string
	cpu.InfoStat
}
type CPU struct {
	//CPU名
	Name string
	//类型
	Type string
	//核心数
	Cores int
	//芯片组型号
	CPUInterface string
	//线程数
	Threads int
	//大核频率
	MainGHZ string
	//小核频率
	AccGHZ string
	//频率
	Ghz string
	//功耗
	CPUPower string
	//上市时间
	CreateTime string
	//首发价格
	Price float64
	//图片地址
	ImgUrl string
}

// NewCPU
// @Description: 初始化CPU结构体用于向前端抛出CPU方法
// @return *CPU
func NewCPU() *CPU {
	return &CPU{}
}

// GetCpuInfo
// @Description: 获取CPU信息
// @return cpu.InfoStat
func GetCpuInfo() CPUInfo {
	var CpuInfo CPUInfo
	CpuInfo.Cores, _ = cpu.Counts(false)
	CpuInfo.Threads, _ = cpu.Counts(true)
	cpuInfo, _ := cpu.Info()
	fmt.Println(cpuInfo, "-------")
	CpuInfo.Name, CpuInfo.Mhz = cpuInfo[0].ModelName, cpuInfo[0].Mhz
	CpuInfo.GHZ = fmt.Sprintf("%.1f", CpuInfo.Mhz/1000)
	return CpuInfo
}

// GetUsingCpuInfo
// @Description: 获取CPU运行时的信息
// @return float64
func GetUsingCpuInfo() float64 {
	useInfo, _ := cpu.Percent(1*time.Second, false)
	fmt.Println(useInfo, "useInfo")
	return useInfo[0]
}

func GetCpuJSONData() interface{} {
	data, err := ioutil.ReadFile("./Json/CPU.json")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var jsonData map[string]interface{}
	e := json.Unmarshal(data, &jsonData)
	if e != nil {
		fmt.Println(e)
	}
	return jsonData["cpu"]
}

func (c *CPU) GetCPUList() {

}
