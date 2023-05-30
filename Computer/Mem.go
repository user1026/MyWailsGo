package Computer

import (
	"changeme/utils"
	"github.com/shirou/gopsutil/mem"
)

type MemInfo struct {
	Total float64 `json:"total"`
	Used  float64 `json:"used"`
	UsePr float64 `json:"usePr"`
}

func GetMemInfo() MemInfo {
	m := MemInfo{}
	memInfo, _ := mem.VirtualMemory()
	m.Total = utils.BToGb(memInfo.Total)
	m.Used = utils.BToGb(memInfo.Used)
	m.UsePr = memInfo.UsedPercent
	return m
}
