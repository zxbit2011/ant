package sys

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/manage-api/handle"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils"
	"github.com/zxbit2011/ant/utils/validation"
)

// 菜单表获取单条数据
func GetSysMenu(c echo.Context) error {
	id := c.FormValue("id")
	var sysMenu model.SysMenu
	if err := global.DB.First(&sysMenu, "id=?", id).Error; err != nil {
		global.Log.Error("GetSysMenu error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, sysMenu)
}

// 菜单表获取所有数据
func GetSysMenuAll(c echo.Context) error {
	parentId := c.FormValue("parent_id")
	t := c.FormValue("type")
	db := global.DB.Model(&model.SysMenu{}).Order("sort ASC,created_at ASC")
	if parentId != "" {
		if parentId == "0" {
			db = db.Where("parent_id = 0")
		} else {
			var menu model.SysMenu
			if err := global.DB.Find(&menu, "id=?", parentId).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return utils.ErrorNull(c, "菜单不存在")
				}
				global.Log.Error("GetSysMenuAll error：%v", err)
				return utils.ErrorNull(c, utils.GetFailResult)
			}
			db = db.Where("relation_ids LIKE ?", menu.RelationIds+"%")
		}
	}
	var sysMenu []model.SysMenu
	if err := db.Find(&sysMenu).Error; err != nil {
		global.Log.Error("GetSysMenuAll error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	if t == "antTree" {
		return utils.SuccessNullMsg(c, model.GetSysMenuTreeUI(sysMenu, "0"))
	} else {
		return utils.SuccessNullMsg(c, model.GetSysMenuTree(sysMenu, "0"))
	}
}

// 菜单表获取分页数据
func GetSysMenuPage(c echo.Context) error {
	db := global.DB.Model(&model.SysMenu{}).Order("id DESC")
	/*	条件搜索范例
		name := c.FormValue("name")
		if name != "" {
			db = db.Where("name like ?", fmt.Sprintf("%%s%%",name))
		}
	*/
	var count int
	var list []model.SysMenu
	if err := db.Count(&count).Error; err != nil {
		global.Log.Error("GetSysMenuPage error：", err)
		return utils.ErrorNull(c, "获取菜单失败")
	}
	pageIndex := utils.GetPageIndex(c.FormValue("page_index"))
	pageSize := utils.GetPageSize(c.FormValue("page_size"))
	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("GetSysMenuPage error：%v", err)
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

// 菜单表保存数据
func SaveSysMenu(c echo.Context) error {

	loginInfo := global.GetLoginInfo(c)
	var sysMenu model.SysMenu
	idStr := c.FormValue("id")
	isEditFlag := false
	if idStr != "" {
		if err := global.DB.First(&sysMenu, "id=?", idStr).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}

		if err := handle.PowerCheck(loginInfo, sysMenu.CreatedBy); err != nil {
			return utils.ErrorNull(c, err.Error())
		}
		isEditFlag = true
	} else {
		//新增 false
		sysMenu.ID = utils.ID()
	}

	//需限制入参参数
	if err := new(utils.CustomBinder).Bind(&sysMenu, c); err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}

	sysMenu.UpdatedBy = loginInfo.ID
	if !isEditFlag {
		sysMenu.ID = utils.ID()
		sysMenu.CreatedBy = loginInfo.ID
	}

	if sysMenu.ParentID == "" {
		sysMenu.ParentID = "0"
	}
	if sysMenu.ParentID == "0" {
		sysMenu.RelationIds = fmt.Sprintf("-0-%s-", sysMenu.ID)
	} else {
		var parentSysMenu model.SysMenu
		if err := global.DB.Model(&model.SysMenu{}).First(&parentSysMenu, "id=?", sysMenu.ParentID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return utils.ErrorNull(c, utils.GetDataNullResult)
			}
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataFailResult)
		}
		sysMenu.RelationIds = fmt.Sprintf("%s%s-", parentSysMenu.RelationIds, sysMenu.ID)
	}

	errs := validation.Valid(&sysMenu)
	if len(errs) > 0 {
		return utils.Error(c, "参数验证失败", errs)
	}
	if isEditFlag {
		//修改
		if err := global.DB.Save(&sysMenu).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.UpdateFailResult)
		}
		return utils.SuccessNull(c, utils.UpdateSuccessResult)
	} else {
		if err := global.DB.Create(&sysMenu).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.AddFailResult)
		}
		return utils.SuccessNull(c, utils.AddSuccessResult)
	}
}

// 菜单表删除数据
func DelSysMenu(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sysMenu model.SysMenu
	if err := global.DB.Model(&model.SysMenu{}).First(&sysMenu, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("GetProjectType error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}

	if err := handle.PowerCheck(loginInfo, sysMenu.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}

	//判断是否存在子菜单
	var childrenCount int
	if err := global.DB.Model(&model.SysMenu{}).Where("parent_id=?", sysMenu.ID).Count(&childrenCount).Error; err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}
	if childrenCount > 0 {
		return utils.ErrorNull(c, "旗下存在子菜单无法进行删除！")
	}

	// 开始事务
	tx := global.DB.Begin()
	if err := tx.Delete(&sysMenu).Error; err != nil {
		tx.Rollback()
		global.Log.Error("DelSysMenu error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	if err := tx.Delete(&model.SysMenuBtn{}, "sys_menu_id=?", sysMenu.ID).Error; err != nil {
		tx.Rollback()
		global.Log.Error("DelSysMenu SysMenuBtn error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	tx.Commit()

	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
