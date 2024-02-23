package Computer

type SSD struct {
	//固态名
	Name string
	//容量
	CAP string
	//颗粒类型
	Type string
	//硬盘长度
	Length string
	//接口类型
	InterfaceType string
	//图片地址
	ImgUrl string
	//接口
	Interface string
}

// NewSSD
// @Description:
// @return *SSD
func NewSSD() *SSD {
	return &SSD{}
}

// GetSSDList
// @Description:
// @receiver s *SSD
func (s *SSD) GetSSDList() {

}

// GetSSDListByHttp
// @Description:
// @receiver s *SSD
func (s *SSD) GetSSDListByHttp() {

}
