package config

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

func InitLog(logConf Log) *logs.BeeLogger {
	var Logger = logs.GetBeeLogger()
	c := fmt.Sprintf(
		`{
			"filename": "%s",
			"maxdays": %d,
			"daily": %s,
			"rotate": %s,
			"level": %d,
			"separate": "[%s]"
		}`,
		logConf.Path,
		logConf.MaxDays,
		"true",
		"true",
		logConf.Level,
		logConf.Separate,
	)

	logs.SetLogger(logs.AdapterMultiFile, c)
	logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)
	return Logger
}
