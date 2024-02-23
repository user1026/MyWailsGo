package Global

import (
	"context"
)

var (
	Global_ConText context.Context
	JsonUrl        = "./Json/"
	Computer       = map[string]string{
		"CPU":       "CPU",
		"GPU":       "显卡",
		"MainBoard": "主板",
		"RAM":       "内存条",
		"Power":     "电源",
		"SSD":       "固态硬盘",
		"HDD":       "机械硬盘",
		"Chassis":   "机箱",
		"Radiator":  "散热器",
	}
)

type PCHardwork struct {
	CPU       string   `json:"CPU"`
	GPU       string   `json:"GPU"`
	MainBoard string   `json:"MainBoard"`
	RAM       string   `json:"RAM"`
	Power     string   `json:"Power"`
	SSD       []string `json:"SSD"`
	HDD       []string `json:"HDD"`
	Chassis   string   `json:"Chassis"`
	Radiator  string   `json:"Radiator"`
}
