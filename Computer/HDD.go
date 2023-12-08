package Computer

type HDD struct {
	//机械硬盘名
	Name string
	//型号
	Model string
	//容量大小
	CAP string
	//硬盘类型
	Type string
	//转速
	Speed string
	//上市日期
	CreatedTime string
	//首发价格
	Price string
}

// NewHDD
// @Description:
// @return *HDD
func NewHDD() *HDD {
	return &HDD{}
}

// GetHDDList
// @Description: 获取本地机械硬盘产品列表
// @receiver h *HDD
func (h *HDD) GetHDDList() {

}

// GetHDDListByHttp
// @Description: 发起请求获取机械硬盘产品列表
// @receiver h *HDD
func (h *HDD) GetHDDListByHttp() {

}
