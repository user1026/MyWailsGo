package AppMenu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

func AddFileMnuList(menu *menu.Menu) {
	FileMenu := menu.AddSubmenu("文件")
	FileMenu.AddText("打开文件", keys.CmdOrCtrl("o"), openFile)
}

func openFile(_ *menu.CallbackData) {

}
