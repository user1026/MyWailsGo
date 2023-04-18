package AppMenu

import (
	"changeme/Event"
	"changeme/Global"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var event = Event.EventList

func AddFileMnuList(menu *menu.Menu) {
	FileMenu := menu.AddSubmenu("文件")
	FileMenu.AddText("打开文件", keys.CmdOrCtrl("o"), openFile)
}

func openFile(_ *menu.CallbackData) {
	fmt.Println(Global.Global_ConText)
	file, err := runtime.OpenFileDialog(Global.Global_ConText, runtime.OpenDialogOptions{
		DefaultDirectory: "C:",
		Title:            "请选择文件",
	})
	if err != nil {
		runtime.MessageDialog(Global.Global_ConText, runtime.MessageDialogOptions{
			Type:          "error",
			Title:         "发生错误",
			Message:       "选择文件时发生错误",
			DefaultButton: "YES",
		})
	}
	fmt.Println(file)
	if file != "" {
		event.GotoJs("openFile", file)
	}

	//runtime.EventsEmit(Global.Global_ConText, "openFile", "123")
}
