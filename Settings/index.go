package Settings

import (
	"MyWailsGo/Global"
	"MyWailsGo/utils"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

type Settings struct {
	ExportFileUrl  string
	ExportFileType string
	UpdateUrl      string
}

// SetExportFileUrl
// @Description: 设置导出的文件存放地址
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
// @Description: 设置导出文件类型
// @return string
func SetExportFileType(value string) bool {
	if err := WriteSettings("ExportFileType", value); err != nil {
		return false
	}
	return true
}

// ReadSettings
// @Description: 读取配置文件
// @return SettingsJson
func ReadSettings() map[string]string {
	path := utils.GetPath()
	jsonFile, e := os.Open(path + "\\settings.json")
	if utils.ErrorDialog(e, "读取配置文件发生错误") {
		fmt.Println("e:", e)
		return map[string]string{}
	}
	fileInfo, _ := os.Lstat(path + "\\settings.json")
	data := make([]byte, fileInfo.Size())
	jsonFile.Read(data)
	defer jsonFile.Close()
	var settingsJson = map[string]string{}
	json.Unmarshal(data, &settingsJson)
	return settingsJson
}

// WriteSettings
// @Description: 写入配置文件
// @param key string
// @param value string
// @return error
func WriteSettings(key string, value string) error {
	path := utils.GetPath()
	SettingsMap := ReadSettings()
	jsonFile, e := os.Create(path + "\\settings.json")
	if utils.ErrorDialog(e, "读取配置文件发生错误") {
		fmt.Println("e:", e)
		return e
	}
	SettingsMap[key] = value
	jsonData, _ := json.Marshal(SettingsMap)
	_, writeErr := jsonFile.Write(jsonData)
	if utils.ErrorDialog(writeErr, "配置文件写入发生错误") {
		fmt.Println("writeErr:", writeErr)
		return writeErr
	}
	return nil
}

func ExportFile(FileType string, data string) bool {
	Settings := ReadSettings()
	if FileType == "txt" {
		file, err := os.Create(Settings["ExportFileUrl"] + "\\电脑配置.txt")
		defer file.Close()
		if utils.ErrorDialog(err, "创建导出文件失败") {
			return false
		}
		_, WriteErr := file.Write([]byte(data))
		if utils.ErrorDialog(WriteErr, "导出失败") {

		}

	} else {
		return false
	}
	return true
}
