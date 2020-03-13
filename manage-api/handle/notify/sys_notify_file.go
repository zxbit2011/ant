package notify

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/manage-api/handle"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils"
)

// 通知消息文件数据删除
func DelSysNotifyFile(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var pf model.SysNotifyFile
	if err := global.DB.Model(&model.SysNotifyFile{}).First(&pf, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, "该通知消息文件不存在")
		}
		global.Log.Error("get SysNotifyFile error：", err)
		return utils.ErrorNull(c, "获取通知消息文件失败")
	}
	if err := handle.PowerCheck(loginInfo, pf.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}
	if err := global.DB.Where("id = ?", id).Delete(&model.SysNotifyFile{}).Error; err != nil {
		global.Log.Error("DelSysNotifyFile error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
