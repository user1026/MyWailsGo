package Global

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	Global_ConText context.Context
)

type Global struct {
	Context context.Context
	Error   error
}

// ErrorDialog
// @Description: 错误弹框
// @param err error
// @param str string
// @return bool
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

// InfoDialog
// @Description: 信息弹框
// @param str string
func InfoDialog(str string) {
	_, err := runtime.MessageDialog(Global_ConText, runtime.MessageDialogOptions{
		Type:          "Info",
		Title:         "通知",
		Message:       str,
		DefaultButton: "是",
	})
	if err != nil {
		return
	}
}
