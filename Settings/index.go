package Settings

import (
	"MyWailsGo/Global"
	"MyWailsGo/utils"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
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
	m := map[string]interface{}{}
	json.Unmarshal([]byte(data), &m)
	Settings := ReadSettings()
	if FileType == "txt" {
		file, err := os.Create(Settings["ExportFileUrl"] + "\\电脑配置.txt")
		defer file.Close()
		if utils.ErrorDialog(err, "创建导出文件失败") {
			return false
		}

		txt := fmt.Sprintf("CPU %v\n"+
			"主板  %v\n"+
			"显卡  %v\n"+
			"内存  %v\n"+
			"散热器 %v\n"+
			"机箱  %v\n"+
			"固态硬盘 %v\n"+
			"机械硬盘 %v\n", m["CPU"], m["MainBorad"], m["GPU"], m["RAM"], m["Radiator"], m["Chassis"], m["SSD"], m["HDD"])
		_, WriteErr := file.Write([]byte(txt))
		if utils.ErrorDialog(WriteErr, "导出失败") {
			return false
		}

	} else {
		excel := excelize.NewFile()
		i := 1
		for k, v := range m {
			excel.SetCellValue("Sheet1", "A"+strconv.Itoa(i), Global.Computer[k])
			excel.SetCellValue("Sheet1", "D"+strconv.Itoa(i), v)
			fmt.Println(Global.Computer[k], v)
			i++
		}
		excel.SetCellValue("Sheet1", "A1", Global.Computer["CPU"])
		excel.SetCellValue("Sheet1", "D1", m["CPU"])
		excel.SetCellValue("Sheet1", "A2", Global.Computer["MainBoard"])
		excel.SetCellValue("Sheet1", "D2", m["MainBoard"])
		excel.SetCellValue("Sheet1", "A3", Global.Computer["GPU"])
		excel.SetCellValue("Sheet1", "D3", m["GPU"])
		excel.SetCellValue("Sheet1", "A4", Global.Computer["RAM"])
		excel.SetCellValue("Sheet1", "D4", m["RAM"])
		excel.SetCellValue("Sheet1", "A5", Global.Computer["Radiator"])
		excel.SetCellValue("Sheet1", "D5", m["Radiator"])
		excel.SetCellValue("Sheet1", "A6", Global.Computer["Power"])
		excel.SetCellValue("Sheet1", "D6", m["Power"])
		excel.SetCellValue("Sheet1", "A7", Global.Computer["Chassis"])
		excel.SetCellValue("Sheet1", "D7", m["Chassis"])
		excel.SetCellValue("Sheet1", "A8", Global.Computer["SSD"])
		excel.SetCellValue("Sheet1", "D8", m[""])
		excel.SetCellValue("Sheet1", "A9", Global.Computer["HDD"])
		excel.SetCellValue("Sheet1", "D9", m[""])
		err := excel.SaveAs(Settings["ExportFileUrl"] + "\\电脑配置.xlsx")
		if utils.ErrorDialog(err, "导出失败") {
			return false
		}
		defer excel.Close()
	}
	return true
}
func ExportFile2(FileType string, data Global.PCHardwork) bool {
	Settings := ReadSettings()
	if FileType == "txt" {
		file, err := os.Create(Settings["ExportFileUrl"] + "\\电脑配置.txt")
		defer file.Close()
		if utils.ErrorDialog(err, "创建导出文件失败") {
			return false
		}
		var SSD, HDD string
		for k, v := range data.SSD {
			if k != 0 {
				SSD += "        	" + v + "\n"
			} else {
				SSD += v + "\n"
			}
		}
		for k, v := range data.HDD {
			if k != 0 {
				HDD += "        	" + v + "\n"
			} else {
				HDD += v + "\n"
			}
		}
		txt := fmt.Sprintf("CPU %v\n"+
			"主板  %v\n"+
			"显卡  %v\n"+
			"内存  %v\n"+
			"散热器 %v\n"+
			"机箱  %v\n"+
			"固态硬盘 %v\n"+
			"机械硬盘 %v\n", data.CPU, data.MainBoard, data.GPU, data.RAM, data.Radiator, data.Chassis, SSD, HDD)
		_, WriteErr := file.Write([]byte(txt))
		if utils.ErrorDialog(WriteErr, "导出失败") {
			return false
		}

	} else {
		excel := excelize.NewFile()
		//i := 1
		//for k, v := range data {
		//	excel.SetCellValue("Sheet1", "A"+strconv.Itoa(i), Global.Computer[k])
		//	excel.SetCellValue("Sheet1", "D"+strconv.Itoa(i), v)
		//	fmt.Println(Global.Computer[k], v)
		//	i++
		//}
		excel.SetCellValue("Sheet1", "A1", Global.Computer["CPU"])
		excel.SetCellValue("Sheet1", "D1", data.CPU)
		excel.SetCellValue("Sheet1", "A2", Global.Computer["MainBoard"])
		excel.SetCellValue("Sheet1", "D2", data.MainBoard)
		excel.SetCellValue("Sheet1", "A3", Global.Computer["GPU"])
		excel.SetCellValue("Sheet1", "D3", data.GPU)
		excel.SetCellValue("Sheet1", "A4", Global.Computer["RAM"])
		excel.SetCellValue("Sheet1", "D4", data.RAM)
		excel.SetCellValue("Sheet1", "A5", Global.Computer["Radiator"])
		excel.SetCellValue("Sheet1", "D5", data.Radiator)
		excel.SetCellValue("Sheet1", "A6", Global.Computer["Chassis"])
		excel.SetCellValue("Sheet1", "D6", data.Chassis)
		excel.SetCellValue("Sheet1", "A7", Global.Computer["SSD"])
		for i := 0; i < len(data.SSD); i++ {
			excel.SetCellValue("Sheet1", "D"+strconv.Itoa(7+i), data.SSD[i])
		}
		if len(data.HDD) != 0 {
			var n = 7 + len(data.SSD)
			excel.SetCellValue("Sheet1", "A"+strconv.Itoa(n), Global.Computer["HDD"])
			for i := 0; i < len(data.SSD); i++ {
				excel.SetCellValue("Sheet1", "D"+strconv.Itoa(n+i), data.HDD[i])
			}
		}
		err := excel.SaveAs(Settings["ExportFileUrl"] + "\\电脑配置.xlsx")
		if utils.ErrorDialog(err, "导出失败") {
			return false
		}
		defer excel.Close()
	}
	return true
}
