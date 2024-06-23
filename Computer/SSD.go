package Computer

type SSD struct {
	//固态名
	Name string
	//容量
	CAP string
	//闪存颗粒
	Flash string
	//芯片颗粒
	Chip string
	//硬盘规格
	Model string
	//寿命
	Life string
	//图片地址
	ImgUrl string
	//接口
	Interface string
	//读写速度
	RWSpeed string
	//缓存大小
	Cache string
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
