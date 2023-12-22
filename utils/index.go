package utils

import (
	"MyWailsGo/Global"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

func BToGb(num uint64) float64 {
	var Gb uint64 = 1024 * 1024 * 1024
	var Mb uint64 = 1024 * 1024
	if (num / Gb) > 1 {
		return float64(num) / float64(Gb)
	} else if (num / Mb) > 1 {
		return float64(num) / float64(Mb)
	}
	return 0.00
}

func OpenDir() string {
	fileName, err := runtime.OpenDirectoryDialog(Global.Global_ConText, runtime.OpenDialogOptions{
		DefaultDirectory: "C:",
		Title:            "请选择文件夹",
	})
	if err != nil {
		runtime.MessageDialog(Global.Global_ConText, runtime.MessageDialogOptions{
			Type:          runtime.QuestionDialog,
			Title:         "发生错误",
			Message:       "选择文件夹发生错误",
			DefaultButton: "确认",
		})
	}
	return fileName
}

func GetPath() string {
	str, _ := os.Getwd()
	return str
}

// ErrorDialog
// @Description: 错误弹框
// @param err error
// @param str string
// @return bool
func ErrorDialog(err error, str string) bool {
	if err != nil {
		runtime.MessageDialog(Global.Global_ConText, runtime.MessageDialogOptions{
			Type:          "error",
			Title:         "发生错误",
			Message:       str,
			DefaultButton: "是",
		})
		return true
	}
	return false
}

// InfoDialog
// @Description: 信息弹框
// @param str string
func InfoDialog(str string) {
	_, err := runtime.MessageDialog(Global.Global_ConText, runtime.MessageDialogOptions{
		Type:          "Info",
		Title:         "通知",
		Message:       str,
		DefaultButton: "是",
	})
	if err != nil {
		return
	}
}
