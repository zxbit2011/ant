package sys

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils"
	"github.com/zxbit2011/ant/utils/validation"
)

// 用户-角色获取单条数据
func GetSysUserRole(c echo.Context) error {
	id := c.FormValue("id")
	var sysUserRole model.SysUserRole
	if err := global.DB.First(&sysUserRole, "id=?", id).Error; err != nil {
		global.Log.Error("GetSysUserRole error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}

	return utils.SuccessNullMsg(c, sysUserRole)
}

// 用户-角色获取所有数据
func GetSysUserRoleAll(c echo.Context) error {
	id := c.FormValue("id")
	var vSysUserRole []model.VSysUserRole
	if err := global.DB.Find(&vSysUserRole, id).Error; err != nil {
		global.Log.Error("GetVSysUserRoleAll error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, vSysUserRole)
}

// 用户-角色获取分页数据
func GetSysUserRolePage(c echo.Context) error {
	db := global.DB.Model(&model.SysUserRole{}).Order("id DESC")
	/*	条件搜索范例
		name := c.FormValue("name")
		if name != "" {
			db = db.Where("name like ?", fmt.Sprintf("%%s%%",name))
		}
	*/
	var count int
	var list []model.SysUserRole
	if err := db.Count(&count).Error; err != nil {
		global.Log.Error("GetSysUserRolePage error：", err)
		return utils.ErrorNull(c, "获取系统角色失败")
	}
	pageIndex := utils.GetPageIndex(c.FormValue("page_index"))
	pageSize := utils.GetPageSize(c.FormValue("page_size"))
	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("GetSysUserRolePage error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, &utils.PageData{
		PageIndex:  pageIndex,
		PageSize:   pageSize,
		Count:      count,
		PageNumber: utils.GetPageNumber(count, pageSize),
		Data:       list,
	})

}

// 用户-角色保存数据
func SaveSysUserRole(c echo.Context) error {
	var sysUserRole model.SysUserRole
	idStr := c.FormValue("id")
	isEditFlag := false
	if idStr != "" {
		if err := global.DB.First(&sysUserRole, "id=?", idStr).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}

		isEditFlag = true
	} else {
		//新增 false
		sysUserRole.ID = utils.ID()
	}

	//需限制入参参数
	if err := new(utils.CustomBinder).Bind(&sysUserRole, c); err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	errs := validation.Valid(&sysUserRole)
	if len(errs) > 0 {
		return utils.Error(c, "参数验证失败", errs)
	}
	if isEditFlag {
		//修改
		if err := global.DB.Save(&sysUserRole).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.UpdateFailResult)
		}
		return utils.SuccessNull(c, utils.UpdateSuccessResult)
	} else {
		if err := global.DB.Create(&sysUserRole).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.AddFailResult)
		}
		return utils.SuccessNull(c, utils.AddSuccessResult)
	}
}

// 用户-角色删除数据
func DelSysUserRole(c echo.Context) error {

	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sysUserRole model.SysUserRole
	if err := global.DB.Model(&model.SysUserRole{}).First(&sysUserRole, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("GetProjectType error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}

	if err := global.DB.Delete(&sysUserRole).Error; err != nil {
		global.Log.Error("DelSysUserRole error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
