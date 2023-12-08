package Computer

type Power struct {
	//电源名
	Name string
	//型号
	Model string
	//功率
	Power string
	//价格
	Price float64
	//图片地址
	ImgUrl string
	//电源类型
	PowerType string
	//模组类型
	Type string
	//详细信息
	Info string
}

// NewPower
// @Description:
// @return *Power
func NewPower() *Power {
	return &Power{}
}

// GetPowerList
// @Description:
// @receiver p *Power
func (p *Power) GetPowerList() {

}

// GetPowerListByHttp
// @Description:
// @receiver p *Power
func (p *Power) GetPowerListByHttp() {

}
