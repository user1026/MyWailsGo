package Computer

type Chassis struct {
	//机箱名
	Name string
	//型号
	Model string
	//长度
	Length string
	//宽度
	Width string
	//高度
	Height string
	//支持的散热器高度
	Radiator string
	//接口
	Interface map[string]string
	//详细信息
	Info string
}

func GetChassisList() {

}
