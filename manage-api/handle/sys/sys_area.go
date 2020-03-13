package sys

import (
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils"
)

// 根据项目查询指标列表
func GetSysAreaList(c echo.Context) error {
	parentId := c.FormValue("parentId")
	if parentId == "" {
		return utils.ErrorNull(c, "请先选择父级地区")
	}
	var sysAreas []model.SysArea
	if err := global.DB.Model(&model.SysArea{}).Find(&sysAreas, "parent_id=?", parentId).Error; err != nil {
		global.Log.Error("GetSysArea error：", err)
		return utils.ErrorNull(c, "获取地区失败")
	}
	return utils.SuccessNullMsg(c, sysAreas)
}
