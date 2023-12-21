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
