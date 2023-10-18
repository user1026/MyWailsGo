package Computer

import (
	"changeme/utils"
	"github.com/shirou/gopsutil/mem"
)

type RamInfo struct {
	Total float64 `json:"total"`
	Used  float64 `json:"used"`
	UsePr float64 `json:"usePr"`
}

type RAM struct {
	Name       string
	Total      int
	CreateTime string
	CL         string
	MHZ        int
	Price      float64
	ImgUrl     string
}

func GetRamInfo() RamInfo {
	m := RamInfo{}
	ramInfo, _ := mem.VirtualMemory()
	m.Total = utils.BToGb(ramInfo.Total)
	m.Used = utils.BToGb(ramInfo.Used)
	m.UsePr = ramInfo.UsedPercent
	return m
}
