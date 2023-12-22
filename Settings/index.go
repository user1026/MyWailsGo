package Settings

import (
	"MyWailsGo/Global"
	"MyWailsGo/utils"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

type SettingsJson struct {
	ExportFileUrl  string
	ExportFileType string
	UpdateUrl      string
}

// SetExportFileUrl
// @Description:
// @return string
func SetExportFileUrl() string {
	fileName, err := runtime.OpenDirectoryDialog(Global.Global_ConText, runtime.OpenDialogOptions{
		DefaultDirectory: "C:",
		Title:            "请选择文件夹",
	})
	if utils.ErrorDialog(err, "选择文件夹发生错误") {
		fmt.Println("err:", err)
		fileName = ""
		return ""
	}
	if writeErr := WriteSettings("ExportFileUrl", fileName); writeErr != nil {
		return ""
	}
	return fileName
}

// SetExportFileType
// @Description:
// @return string
func SetExportFileType(value string) bool {
	if err := WriteSettings("ExportFileType", value); err != nil {
		return false
	}
	return true
}

// ReadSettings
// @Description:
// @return SettingsJson
func ReadSettings() SettingsJson {
	path := utils.GetPath()
	jsonFile, e := os.Open(path + "\\settings.json")
	if utils.ErrorDialog(e, "读取配置文件发生错误") {
		fmt.Println("e:", e)
		return SettingsJson{}
	}
	fileInfo, _ := os.Lstat(path + "\\settings.json")
	data := make([]byte, fileInfo.Size())
	jsonFile.Read(data)
	defer jsonFile.Close()
	var settingsJson SettingsJson
	json.Unmarshal(data, &settingsJson)
	return settingsJson
}

func WriteSettings(key string, value string) error {
	path := utils.GetPath()
	jsonFile, e := os.Create(path + "\\settings.json")
	if utils.ErrorDialog(e, "读取配置文件发生错误") {
		fmt.Println("e:", e)
		return e
	}
	fileInfo, _ := os.Lstat(path + "\\settings.json")
	data := make([]byte, fileInfo.Size())
	jsonFile.Read(data)
	defer jsonFile.Close()
	var SettingsMap = map[string]string{}
	json.Unmarshal(data, &SettingsMap)
	SettingsMap[key] = value
	jsonData, _ := json.Marshal(SettingsMap)
	_, writeErr := jsonFile.Write(jsonData)
	if utils.ErrorDialog(writeErr, "配置文件写入发生错误") {
		fmt.Println("writeErr:", writeErr)
		return writeErr
	}
	return nil
}
