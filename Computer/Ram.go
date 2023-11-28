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
	//内存名
	Name string
	//总容量
	Total string
	//上市时间
	CreateTime string
	//时序
	CL string
	//频率
	MHZ string
	//价格
	Price float64
	//图片地址
	ImgUrl string
	//内存类型
	Type string
}

// GetRamInfo
// @Description: 获取本机内存信息
// @return RamInfo
func GetRamInfo() RamInfo {
	m := RamInfo{}
	ramInfo, _ := mem.VirtualMemory()
	m.Total = utils.BToGb(ramInfo.Total)
	m.Used = utils.BToGb(ramInfo.Used)
	m.UsePr = ramInfo.UsedPercent
	return m
}

func GetRamList() {

}
