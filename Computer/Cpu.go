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
	Interface string
	//线程数
	Threads int
	//大核频率
	MainGHZ string
	//小核频率
	AccGHZ string
	//频率
	Ghz string
	//功耗
	Power string
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
func (c *CPU) GetCpuInfo() CPUInfo {
	var CpuInfo CPUInfo
	cpuInfo, _ := cpu.Info()
	Cores, _ := cpu.Counts(false)
	Threads, _ := cpu.Counts(true)
	CpuInfo.Name, CpuInfo.Mhz = cpuInfo[0].ModelName, cpuInfo[0].Mhz
	CpuInfo.GHZ = fmt.Sprintf("%.1f", CpuInfo.Mhz/1000)
	return CPUInfo{
		Cores:   Cores,
		Threads: Threads,
		Name:    cpuInfo[0].ModelName,
		GHZ:     fmt.Sprintf("%.1f", cpuInfo[0].Mhz/1000),
	}
}

// GetUsingCpuInfo
// @Description: 获取CPU运行时的信息
// @return float64
func (c *CPU) GetUsingCpuInfo() float64 {
	useInfo, _ := cpu.Percent(5*time.Second, false)
	fmt.Println(useInfo, "useInfo")
	return useInfo[0]
}

func (c *CPU) GetCpuJSONData() interface{} {
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
