package Computer

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type Computer struct {
}

func GetCpuInfo() []cpu.InfoStat {
	cpuInfo, _ := cpu.Info()
	fmt.Println(cpuInfo)
	return cpuInfo
}

func GetUsingCpuInfo() []float64 {
	useInfo, _ := cpu.Percent(2*time.Second, false)
	fmt.Println(useInfo, "useInfo")
	return useInfo
}
