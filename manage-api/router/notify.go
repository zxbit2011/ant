package router

import (
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/handle/notify"
)

//账号
func SysNotifyRouter(e *echo.Group) {
	api := e.Group("/sys/notify")
	{
		api.POST("/get", notify.GetSysNotify)
		api.POST("/info", notify.GetSysNotifyInfo)
		api.POST("/save", notify.SaveSysNotify)
		api.POST("/list", notify.GetSysNotifyList)
		api.POST("/page", notify.GetSysNotifyPage)
		api.POST("/release", notify.ReleaseSysNotify)
		api.POST("/remove", notify.DelSysNotify)

		// 通知信息文件
		api.POST("/file/remove", notify.DelSysNotifyFile)
	}
}
