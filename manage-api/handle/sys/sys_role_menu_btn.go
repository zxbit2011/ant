package sys

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils"
)

// 角色-菜单获取单条数据
func GetSysRoleMenuBtn(c echo.Context) error {
	id := c.FormValue("id")
	var sysRoleMenuBtn model.SysRoleMenuBtn
	if err := global.DB.First(&sysRoleMenuBtn, "id=?", id).Error; err != nil {
		global.Log.Error("GetSysRoleMenuBtn error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}

	return utils.SuccessNullMsg(c, sysRoleMenuBtn)
}

// 角色-菜单获取所有数据
func GetSysRoleMenuBtnAll(c echo.Context) error {
	roleId := c.FormValue("roleId")
	menuId := c.FormValue("menuId")
	var vSysRoleMenuBtn []model.VSysRoleMenuBtn
	if err := global.DB.Find(&vSysRoleMenuBtn, "sys_role_id=? AND sys_menu_id=?", roleId, menuId).Error; err != nil {
		global.Log.Error("GetVSysRoleMenuBtnAll error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, vSysRoleMenuBtn)

}

// 角色-菜单获取分页数据
func GetSysRoleMenuBtnPage(c echo.Context) error { //视图查询
	db := global.DB.Model(&model.VSysRoleMenuBtn{}).Order("id DESC")

	/*	条件搜索范例
		name := c.FormValue("name")
		if name != "" {
			db = db.Where("name like ?", fmt.Sprintf("%%s%%",name))
		}
	*/
	var count int
	var list []model.VSysRoleMenuBtn
	if err := db.Count(&count).Error; err != nil {
		global.Log.Error("GetSysRoleMenuBtnPage error：", err)
		return utils.ErrorNull(c, "获取角色菜单按钮权限失败")
	}
	pageIndex := utils.GetPageIndex(c.FormValue("pageIndex"))
	pageSize := utils.GetPageSize(c.FormValue("pageSize"))
	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("GetVSysRoleMenuBtnPage error：%v", err)
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

// 角色-菜单保存数据
func SaveSysRoleMenuBtn(c echo.Context) error {
	id := c.FormValue("id")
	roleID := c.FormValue("roleId")
	add := c.FormValue("add")
	if id == "" || roleID == "" || add == "" {
		return utils.ErrorNull(c, "无效操作")
	}
	var sysRole model.SysRole
	if err := global.DB.Find(&sysRole, "id=?", roleID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, "无角色数据")
		}
		global.Log.Error("SaveSysRoleMenuBtn error:%s", err.Error())
		return utils.ErrorNull(c, "保存失败")
	}

	var sysMenuBtn model.SysMenuBtn
	if err := global.DB.Find(&sysMenuBtn, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, "无按钮权限数据")
		}
		global.Log.Error("SaveSysRoleMenuBtn error:%s", err.Error())
		return utils.ErrorNull(c, "保存失败")
	}

	if add == "true" {
		if err := global.DB.Create(&model.SysRoleMenuBtn{ID: utils.ID(), SysRoleID: roleID, SysMenuID: sysMenuBtn.SysMenuID, SysMenuBtnID: id}).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, "保存失败")
		}
	} else {
		if err := global.DB.Delete(&model.SysRoleMenuBtn{}, "sys_role_id=? AND sys_menu_btn_id=? ", roleID, id).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, "保存失败")
		}
	}
	return utils.SuccessNull(c, "保存成功")
}

// 角色-菜单删除数据
func DelSysRoleMenuBtn(c echo.Context) error {

	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sysRoleMenuBtn model.SysRoleMenuBtn
	if err := global.DB.Model(&model.SysRoleMenuBtn{}).First(&sysRoleMenuBtn, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("GetProjectType error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}

	if err := global.DB.Delete(&sysRoleMenuBtn).Error; err != nil {
		global.Log.Error("DelSysRoleMenuBtn error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
