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
	RadiatorHeight string
	//接口
	Interface map[string]string
	//详细信息
	Info string
	//支持的主板类型
	SupportMainBoardTypes string
	//支持的电源类型
	SupportPowerTypes string
	//支持的显卡长度
	SupportGPULength string
}

// NewChassis
// @Description:
// @return *Chassis
func NewChassis() *Chassis {
	return &Chassis{}
}

// GetChassisList
// @Description:
// @receiver c *Chassis
func (c *Chassis) GetChassisList() {

}

// GetChassisListByHttp
// @Description:
// @receiver c *Chassis
func (c *Chassis) GetChassisListByHttp() {

}
