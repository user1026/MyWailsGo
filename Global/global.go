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
