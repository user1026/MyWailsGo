package Computer

type MainBoard struct {
	//主板
	Name string
	//型号
	Model string
	//平台
	Type string
	//芯片组
	CPUInterface string
	//上市时间
	CreateTime string
	//价格
	Price float64
	//接口
	Interface map[string]string
	//详细信息
	Info string
	//WIFI
	WIFI string
}
