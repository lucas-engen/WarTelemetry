package hudmsg

import (
	"strings"

	"github.com/lucas-engen/WarTelemetry/utils"
)

type Hudmsg struct {
	Events []string     `json:"events"`
	Damage []DamageData `json:"damage"`
}

type DamageData struct {
	Id      int    `json:"id"`
	Message string `json:"msg"`
	Sender  string `json:"sender"`
	Enemy   bool   `json:"enemy"`
	Mode    string `json:"mode"`
}

var path string = "hudmsg"

// GetURL function retrieves the URL to get data about gamechat
func GetURL() string {
	url := utils.GetBaseURL()
	url = strings.ReplaceAll(url, "$hostname$", utils.GetHostname())
	url = strings.ReplaceAll(url, "$path$", path)
	return url
}