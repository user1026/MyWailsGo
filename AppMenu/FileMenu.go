package AppMenu

import (
	"MyWailsGo/Event"
	"MyWailsGo/Global"
	"MyWailsGo/utils"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/xuri/excelize/v2"
)

var event = Event.EventList

func AddFileMnuList(menu *menu.Menu) {
	FileMenu := menu.AddSubmenu("文件")
	FileMenu.AddText("打开文件", keys.CmdOrCtrl("o"), openFile)
}

func openFile(_ *menu.CallbackData) {
	fileName, err := runtime.OpenFileDialog(Global.Global_ConText, runtime.OpenDialogOptions{
		DefaultDirectory: "C:",
		Title:            "请选择文件",
	})
	if utils.ErrorDialog(err, "选择文件时发生错误") {
		return
	}
	file, openErr := excelize.OpenFile(fileName)
	if utils.ErrorDialog(openErr, "读取文件时发生错误") {
		return
	}
	fmt.Println(file)
	rows, _ := file.GetRows("Sheet1")
	fmt.Println(rows, "- - - - - -")
	if rows != nil {
		event.GotoJs("openFile", rows)
	}
	defer func() {
		file.Close()
	}()
}
