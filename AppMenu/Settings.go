package AppMenu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

func AddSettingsMenuList(menu *menu.Menu) {
	SettingsMenu := menu.AddSubmenu("设置")
	SettingsMenu.AddText("路径", keys.CmdOrCtrl("p"), SetPath)
}

func SetPath(_ *menu.CallbackData) {

}
