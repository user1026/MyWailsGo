package Settings

import (
	"changeme/Global"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func SetExportFileUrl() string {
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
		fileName = ""
	}
	return fileName
}
