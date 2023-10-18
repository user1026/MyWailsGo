package Computer

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
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
	Name       string
	Cores      int
	Threads    int
	GHZ        int
	CPUPower   string
	CreateTime string
	Price      float64
	ImgUrl     string
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
