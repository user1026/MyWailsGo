package Computer

type GPU struct {
	//显卡名
	Name string
	//型号
	Model string
	//图片地址
	ImgUrl string
	//显卡功耗
	GPUPower string
	//上市时间
	CreateTime string
	//首发价格
	Price float64
	//bt
	Bit int
	//显存大小
	RAM int
	//显存频率
	RAMGhz string
	//制造商
	OEM string
	//显存类型
	RAMType string
	//流处理器数量
	LCPUNumber int
	//长度
	Length string
	//宽度
	Width string
	//高度
	Height string
}

func GetGPUList() {

}
