package sys

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/manage-api/handle"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils"
	"github.com/zxbit2011/ant/utils/validation"
)

// 菜单表获取单条数据
func GetSysMenuBtn(c echo.Context) error {
	id := c.FormValue("id")
	var sysMenuBtn model.SysMenuBtn
	if err := global.DB.First(&sysMenuBtn, "id=?", id).Error; err != nil {
		global.Log.Error("GetSysMenuBtn error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, sysMenuBtn)
}

// 菜单表获取所有数据
func GetSysMenuBtnAll(c echo.Context) error {
	menuId := c.FormValue("menuId")
	var sysMenuBtn []model.SysMenuBtn
	if err := global.DB.Order("created_at DESC").Find(&sysMenuBtn, "sys_menu_id=?", menuId).Error; err != nil {
		global.Log.Error("GetSysMenuBtnAll error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, sysMenuBtn)
}

// 菜单表获取分页数据
func GetSysMenuBtnPage(c echo.Context) error {
	db := global.DB.Model(&model.SysMenuBtn{}).Order("id DESC")
	/*	条件搜索范例
		name := c.FormValue("name")
		if name != "" {
			db = db.Where("name like ?", fmt.Sprintf("%%s%%",name))
		}
	*/
	var count int
	var list []model.SysMenuBtn
	if err := db.Count(&count).Error; err != nil {
		global.Log.Error("GetSysMenuBtnPage error：", err)
		return utils.ErrorNull(c, "获取菜单按钮权限失败")
	}
	pageIndex := utils.GetPageIndex(c.FormValue("page_index"))
	pageSize := utils.GetPageSize(c.FormValue("page_size"))
	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("GetSysMenuBtnPage error：%v", err)
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
func SaveSysMenuBtn(c echo.Context) error {

	loginInfo := global.GetLoginInfo(c)
	var sysMenuBtn model.SysMenuBtn
	idStr := c.FormValue("id")
	isEditFlag := false
	if idStr != "" {
		if err := global.DB.First(&sysMenuBtn, "id=?", idStr).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}

		if err := handle.PowerCheck(loginInfo, sysMenuBtn.CreatedBy); err != nil {
			return utils.ErrorNull(c, err.Error())
		}
		isEditFlag = true
	} else {
		//新增 false
		sysMenuBtn.ID = utils.ID()
	}

	//需限制入参参数
	if err := new(utils.CustomBinder).Bind(&sysMenuBtn, c); err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	sysMenuBtn.UpdatedBy = loginInfo.ID
	if !isEditFlag {
		sysMenuBtn.ID = utils.ID()
		sysMenuBtn.CreatedBy = loginInfo.ID
	}
	errs := validation.Valid(&sysMenuBtn)
	if len(errs) > 0 {
		return utils.Error(c, "参数验证失败", errs)
	}
	if isEditFlag {
		//修改
		if err := global.DB.Save(&sysMenuBtn).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.UpdateFailResult)
		}
		return utils.SuccessNull(c, utils.UpdateSuccessResult)
	} else {
		if err := global.DB.Create(&sysMenuBtn).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.AddFailResult)
		}
		return utils.SuccessNull(c, utils.AddSuccessResult)
	}
}

// 菜单表删除数据
func DelSysMenuBtn(c echo.Context) error {

	loginInfo := global.GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sysMenuBtn model.SysMenuBtn
	if err := global.DB.Model(&model.SysMenuBtn{}).First(&sysMenuBtn, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("GetProjectType error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}

	if err := handle.PowerCheck(loginInfo, sysMenuBtn.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}

	if err := global.DB.Delete(&sysMenuBtn).Error; err != nil {
		global.Log.Error("DelSysMenuBtn error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
