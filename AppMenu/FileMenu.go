package AppMenu

import (
	"changeme/Event"
	"changeme/Global"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

var event = Event.EventList

func AddFileMnuList(menu *menu.Menu) {
	FileMenu := menu.AddSubmenu("文件")
	FileMenu.AddText("打开文件", keys.CmdOrCtrl("o"), openFile)
}

func openFile(_ *menu.CallbackData) {
	fmt.Println(Global.Global_ConText)
	fileName, err := runtime.OpenFileDialog(Global.Global_ConText, runtime.OpenDialogOptions{
		DefaultDirectory: "C:",
		Title:            "请选择文件",
	})
	if Global.ErrorDialog(err, "选择文件时发生错误") {
		return
	}
	//file, openErr := excelize.OpenFile()
	file, openErr := os.Open(fileName)
	if Global.ErrorDialog(openErr, "打开文件时发生错误") {
		return
	}
	fmt.Println(file)
	fileInfo, er := file.Stat()
	if Global.ErrorDialog(er, "获取文件信息时发生错误") {
		return
	}
	var fileData = make([]byte, fileInfo.Size())
	_, readerr := file.Read(fileData)
	if Global.ErrorDialog(readerr, "读取文件信息时发生错误") {
		return
	}
	fmt.Println(string(fileData))
	if file != nil {
		event.GotoJs("openFile", file)
	}
	defer func() {
		file.Close()
	}()
	//runtime.EventsEmit(Global.Global_ConText, "openFile", "123")
}
