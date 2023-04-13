package AppMenu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func AllMenuList() *menu.Menu {
	AppMenu := menu.NewMenu()
	AddFileMnuList(AppMenu)
	AddSettingsMenuList(AppMenu)
	return AppMenu
}
