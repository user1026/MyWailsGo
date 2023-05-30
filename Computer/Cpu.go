package Computer

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type Computer struct {
}

func GetCpuInfo() cpu.InfoStat {
	cpuInfo, _ := cpu.Info()
	return cpuInfo[0]
}

func GetUsingCpuInfo() float64 {
	useInfo, _ := cpu.Percent(1*time.Second, false)
	fmt.Println(useInfo, "useInfo")
	return useInfo[0]
}
