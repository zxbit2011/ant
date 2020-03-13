package sys

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils"
)

// 角色-菜单获取单条数据
func GetSysRoleMenu(c echo.Context) error {
	id := c.FormValue("id")
	var sysRoleMenu model.SysRoleMenu
	if err := global.DB.First(&sysRoleMenu, "id=?", id).Error; err != nil {
		global.Log.Error("GetSysRoleMenu error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}

	return utils.SuccessNullMsg(c, sysRoleMenu)
}

// 角色-菜单获取所有数据
func GetSysRoleMenuAll(c echo.Context) error {
	id := c.FormValue("id")
	var vSysRoleMenu []model.VSysRoleMenu
	if err := global.DB.Find(&vSysRoleMenu, "sys_role_id=?", id).Error; err != nil {
		global.Log.Error("GetVSysRoleMenuAll error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, vSysRoleMenu)

}

// 角色-菜单获取分页数据
func GetSysRoleMenuPage(c echo.Context) error { //视图查询
	db := global.DB.Model(&model.VSysRoleMenu{}).Order("id DESC")

	/*	条件搜索范例
		name := c.FormValue("name")
		if name != "" {
			db = db.Where("name like ?", fmt.Sprintf("%%s%%",name))
		}
	*/
	var count int
	var list []model.VSysRoleMenu
	if err := db.Count(&count).Error; err != nil {
		global.Log.Error("GetSysRoleMenuPage error：", err)
		return utils.ErrorNull(c, "获取角色菜单失败")
	}
	pageIndex := utils.GetPageIndex(c.FormValue("page_index"))
	pageSize := utils.GetPageSize(c.FormValue("page_size"))
	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("GetVSysRoleMenuPage error：%v", err)
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

// 角色 - 菜单保存数据
func SaveSysRoleMenu(c echo.Context) error {
	ids := c.FormValue("ids")
	roleID := c.FormValue("roleId")
	add := c.FormValue("add")
	if ids == "" || roleID == "" || add == "" {
		return utils.ErrorNull(c, "无效操作")
	}
	tip := "保存菜单资源失败"
	var menuIds []string
	if err := json.Unmarshal([]byte(ids), &menuIds); err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, tip)
	}

	var sysRole model.SysRole
	if err := global.DB.Find(&sysRole, "id=?", roleID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, "无角色数据")
		}
		global.Log.Error("SaveSysRoleMenuBtn error:%s", err.Error())
		return utils.ErrorNull(c, "保存失败")
	}

	var menuCount int
	if err := global.DB.Model(&model.SysMenu{}).Where("id in(?)", menuIds).Count(&menuCount).Error; err != nil {
		global.Log.Error("SaveSysRoleMenu error:%s", err.Error())
		return utils.ErrorNull(c, "保存失败")
	}
	if menuCount != len(menuIds) {
		return utils.ErrorNull(c, "保存失败")
	}

	// 开始事务
	tx := global.DB.Begin()
	if add == "true" {
		//新增菜单
		if len(menuIds) > 0 {
			//添加资源权限
			for _, value := range menuIds {
				if err := tx.Create(&model.SysRoleMenu{ID: utils.ID(), SysRoleID: roleID, SysMenuID: value}).Error; err != nil {
					tx.Rollback()
					global.Log.Error(err.Error())
					return utils.ErrorNull(c, tip)
				}
			}
		}
	} else {
		//删除现有资源权限以及按钮权限
		if err := tx.Delete(&model.SysRoleMenu{}, "sys_menu_id in(?)", menuIds).Error; err != nil {
			tx.Rollback()
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, tip)
		}
		if err := tx.Delete(&model.SysRoleMenuBtn{}, "sys_menu_id in(?)", menuIds).Error; err != nil {
			tx.Rollback()
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, tip)
		}
	}
	tx.Commit()
	return utils.SuccessNull(c, "保存成功")
}

// 角色-菜单删除数据
func DelSysRoleMenu(c echo.Context) error {

	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sysRoleMenu model.SysRoleMenu
	if err := global.DB.Model(&model.SysRoleMenu{}).First(&sysRoleMenu, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("GetProjectType error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}

	if err := global.DB.Delete(&sysRoleMenu).Error; err != nil {
		global.Log.Error("DelSysRoleMenu error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
