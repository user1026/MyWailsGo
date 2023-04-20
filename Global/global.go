package Global

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	Global_ConText context.Context
	CPU            string = "CPU"
	FSH            string = "散热器/水冷"
	RAM            string = "内存条"
	Motherboard    string = "主板"
	PSU            string = "电源"
	GPU            string = "显卡"
	SSD            string = "固态硬盘"
	HDD            string = "机械硬盘"
	CPU_P          string = "CPU功率"
	FSH_P          string = "散热器/水冷功率"
	Motherboard_P  string = "主板功率"
	GPU_P          string = "显卡功率"
	Other_P        string = "其他功率"
)

type Global struct {
	Context context.Context
	Error   error
}
type ExcelData struct {
	CPU []string
	FS  []string
	ZB  []string
}

func ErrorDialog(err error, str string) bool {
	if err != nil {
		runtime.MessageDialog(Global_ConText, runtime.MessageDialogOptions{
			Type:          "error",
			Title:         "发生错误",
			Message:       str,
			DefaultButton: "是",
		})
		return true
	}
	return false
}
func InfoDialog(str string) {
	runtime.MessageDialog(Global_ConText, runtime.MessageDialogOptions{
		Type:          "Info",
		Title:         "通知",
		Message:       str,
		DefaultButton: "是",
	})
}
