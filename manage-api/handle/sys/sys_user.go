package sys

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/manage-api/handle"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils"
	"github.com/zxbit2011/ant/utils/encrypt"
	"github.com/zxbit2011/ant/utils/enum"
	"github.com/zxbit2011/ant/utils/validation"
)

//判定第三方数据中心的用户是否存在
func CheckSysUser(apiId, userType string) (sysUser model.VSysUser, err error) {
	err = global.DB.First(&sysUser, "api_id=? AND user_type=?", apiId, userType).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

// 用户表获取单条数据
func GetSysUser(c echo.Context) error {
	id := c.FormValue("id")
	var sysUser model.SysUser
	if err := global.DB.First(&sysUser, "id=?", id).Error; err != nil {
		global.Log.Error("GetSysUser error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}

	loginInfo := global.GetLoginInfo(c)
	if err := handle.PowerCheck(loginInfo, sysUser.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}
	return utils.SuccessNullMsg(c, sysUser)
}

// 用户表获取所有数据
func GetSysUserAll(c echo.Context) error {
	id := c.FormValue("id")
	var sysUser []model.SysUser
	if err := global.DB.Find(&sysUser, id).Error; err != nil {
		global.Log.Error("GetSysUserAll error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, sysUser)
}

// 用户表获取分页数据
func GetSysUserPage(c echo.Context) error { //视图查询
	//排序
	order := "created_at DESC"
	sortField := c.FormValue("sortField")
	sortOrder := c.FormValue("sortOrder")
	if sortField != "" && sortOrder != "" {
		order = fmt.Sprintf("%s %s", sortField, sortOrder)
	}
	db := global.DB.Model(&model.VSysUser{}).Order(order)
	loginInfo := global.GetLoginInfo(c)
	// 非超级管理员 && 数据权限验证
	if !loginInfo.IsSuperAdmin {
		if loginInfo.IsAdmin { //管理员 - 可查看旗下组织机构的数据
			db = db.Where("relation_ids LIKE ?", loginInfo.OfficeRelationIds+"%")
		} else {
			//普通账号 - 仅能看自己的数据
			db = db.Where("created_by=?", loginInfo.ID)
		}
	}

	//条件搜索范例
	name := c.FormValue("name")
	usertype := c.FormValue("usertype")
	roleId := c.FormValue("role_id")
	officeId := c.FormValue("office_id")
	startTime := c.FormValue("time1")
	endTime := c.FormValue("time2")

	valid := validation.Validation{}
	valid.MaxSize(name, 50, "name", "姓名字数超长")
	valid.MaxSize(roleId, 64, "role_id", "角色id字数超长")
	valid.MaxSize(officeId, 64, "office_id", "部门id字数超长")
	if errs := validation.ValidSingle(valid); errs != nil {
		return utils.ParamsError(c, errs)
	}

	if name != "" {
		db = db.Where("name like ?", fmt.Sprintf("%%%s%%", name))
	}
	if roleId != "" {
		db = db.Where("FIND_IN_SET(?,role_ids)", roleId)
	}
	if officeId != "" {
		db = db.Where("office_id=?", officeId)
	}
	if usertype != "" {
		db = db.Where("user_type=?", usertype)
	}
	if startTime != "" && endTime != "" {
		db = db.Where("created_at>=? AND created_at<=?", startTime, endTime)
	}
	var err error
	var count int
	var list []model.VSysUser
	if err := db.Count(&count).Error; err != nil {
		global.Log.Error("GetSysUserPage error：", err)
		return utils.ErrorNull(c, "获取用户失败")
	}
	pageIndex := utils.GetPageIndex(c.FormValue("page_index"))
	pageSize := utils.GetPageSize(c.FormValue("page_size"))
	if err = db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("GetVSysUserPage error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	var roleList = make(map[string][]string)
	//获取角色
	if len(list) > 0 {
		var ids []string
		for _, v := range list {
			ids = append(ids, v.ID)
		}
		//获取用户角色
		var vSysUserRole []model.VSysUserRole
		if err = global.DB.Where("sys_user_id IN (?)", ids).Find(&vSysUserRole).Error; err != nil {
			global.Log.Error("query v_sys_user_role error： %s", err.Error())
			return utils.ErrorNull(c, "获取用户所属角色失败")
		}
		if len(vSysUserRole) > 0 {
			for _, v := range vSysUserRole {
				roleList[v.SysUserId] = append(roleList[v.SysUserId], v.Name)
			}
		}
	}
	return utils.SuccessNullMsg(c, &utils.PageData{
		PageIndex:  pageIndex,
		PageSize:   pageSize,
		Count:      count,
		PageNumber: utils.GetPageNumber(count, pageSize),
		Data:       list,
		ExtendData: map[string]interface{}{
			"roleList":        roleList,
		},
	})

}

// 用户表保存数据
func SaveSysUser(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	var sysUser model.SysUser
	idStr := c.FormValue("id")
	isEditFlag := false
	if idStr != "" {
		if err := global.DB.First(&sysUser, "id=?", idStr).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		if err := handle.PowerCheck(loginInfo, sysUser.CreatedBy); err != nil {
			return utils.ErrorNull(c, err.Error())
		}
		isEditFlag = true
	} else {
		//新增 false
		sysUser.ID = utils.ID()
	}

	//需限制入参参数
	if err := new(utils.CustomBinder).Bind(&sysUser, c); err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}

	sysUser.UpdatedBy = loginInfo.ID
	if !isEditFlag {
		sysUser.ID = utils.ID()
		sysUser.CreatedBy = loginInfo.ID
		sysUser.CompanyID = loginInfo.CompanyId
		sysUser.OfficeID = loginInfo.OfficeId
		sysUser.LoginFlag = enum.LoginFlagOK
	}
	pwd := c.FormValue("password")
	if pwd != "" {
		//密码加密
		salt := utils.RangeString(5)
		sysUser.Salt = salt
		sysUser.Password = encrypt.Sha1Encode(sysUser.Password + salt)
	}
	errs := validation.Valid(&sysUser)
	if len(errs) > 0 {
		return utils.ParamsError(c, errs)
	}

	//账号是否占用
	var isUsername int
	db := global.DB.Model(&model.SysUser{}).Where("username = ?", sysUser.Username)
	if isEditFlag {
		db = db.Where("id<>?", sysUser.ID)
	}
	if err := db.Count(&isUsername).Error; err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, "检测用户账号失败")
	}
	if isUsername > 0 {
		return utils.ErrorNull(c, "用户账号已被占用")
	}
	//判断手机号码是否占用
	var isMobile int
	db = global.DB.Model(&model.SysUser{}).Where("mobile = ?", sysUser.Mobile)
	if isEditFlag {
		db = db.Where("id<>?", sysUser.ID)
	}
	if err := db.Count(&isMobile).Error; err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, "检测手机号码失败")
	}
	if isMobile > 0 {
		return utils.ErrorNull(c, "手机号码已被占用")
	}
	//添加用户角色
	roleId := c.FormValue("role_id")
	if roleId == "" {
		return utils.ErrorNull(c, "请选择用户角色")
	}
	var roleIds []string
	err := json.Unmarshal([]byte(roleId), &roleIds)
	if err != nil {
		return utils.ErrorNull(c, "请选择用户角色")
	}
	if len(roleIds) == 0 {
		return utils.ErrorNull(c, "请选择用户角色")
	}
	var isRoleId int
	if err := global.DB.Model(&model.SysRole{}).Where("id in (?)", roleIds).Count(&isRoleId).Error; err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, "检测用户角色失败")
	}
	if isRoleId != len(roleIds) {
		return utils.ErrorNull(c, "选择的用户角色不存在")
	}

	// 开始事务
	tx := global.DB.Begin()
	if isEditFlag {

		//删除角色
		if err := tx.Delete(&model.SysUserRole{}, "sys_user_id=?", sysUser.ID).Error; err != nil {
			tx.Rollback()
			return utils.ErrorNull(c, utils.AddFailResult)
		}
		//创建角色
		for _, roleId := range roleIds {
			userRole := model.SysUserRole{}
			userRole.ID = utils.ID()
			userRole.SysRoleID = roleId
			userRole.SysUserID = sysUser.ID
			if err := tx.Create(&userRole).Error; err != nil {
				tx.Rollback()
				global.Log.Error(err.Error())
				return utils.ErrorNull(c, utils.AddFailResult)
			}
		}
		//修改
		if err := tx.Save(&sysUser).Error; err != nil {
			tx.Rollback()
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.UpdateFailResult)
		}
	} else {
		//创建用户
		if err := tx.Create(&sysUser).Error; err != nil {
			tx.Rollback()
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.AddFailResult)
		}
		//创建角色
		for _, roleId := range roleIds {
			userRole := model.SysUserRole{}
			userRole.ID = utils.ID()
			userRole.SysRoleID = roleId
			userRole.SysUserID = sysUser.ID
			if err := tx.Create(&userRole).Error; err != nil {
				tx.Rollback()
				global.Log.Error(err.Error())
				return utils.ErrorNull(c, utils.AddFailResult)
			}
		}
	}
	tx.Commit()
	return utils.SuccessNull(c, "保存成功")
}

// 用户表删除数据
func DelSysUser(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sysUser model.SysUser
	if err := global.DB.Model(&model.SysUser{}).First(&sysUser, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("DisableSysUser error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}

	if err := handle.PowerCheck(loginInfo, sysUser.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}

	if err := global.DB.Delete(&sysUser).Error; err != nil {
		global.Log.Error("DelSysUser error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}

// 禁用系统用户
func DisableSysUser(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sysUser model.SysUser
	var err error
	if err := global.DB.Model(&model.SysUser{}).First(&sysUser, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("DisableSysUser error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}

	if err := handle.PowerCheck(loginInfo, sysUser.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}

	//获取用户角色
	var vSysUserRole []model.VSysUserRole
	if err = global.DB.Where("sys_user_id=?", sysUser.ID).Find(&vSysUserRole).Error; err != nil {
		global.Log.Error("query v_sys_user_role error： %s", err.Error())
		return utils.ErrorNull(c, "获取用户所属角色失败")
	}

	// 判断是否存在超级管理员角色，存在则不与删除
	for _, v := range vSysUserRole {
		if v.RoleType == enum.UserPostSuperAdminCode {
			return utils.ErrorNull(c, "无法禁用超级管理员角色")
		}
	}

	loginFlag := enum.LoginFlagNo
	if sysUser.LoginFlag == enum.LoginFlagNo {
		loginFlag = enum.LoginFlagOK
	}

	if err := global.DB.Model(&sysUser).Update("login_flag", loginFlag).Error; err != nil {
		global.Log.Error("DelSysUser error：%v", err)
		return utils.ErrorNull(c, utils.UpdateFailResult)
	}
	return utils.SuccessNull(c, utils.UpdateSuccessResult)
}

/*
密码重置
参数：
	管理员id：id
*/
func ResetPassword(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sysUser model.SysUser
	if err := global.DB.First(&sysUser, "id=?", id).Error; err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, "数据不存在")
	}
	if err := handle.PowerCheck(loginInfo, sysUser.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}
	password := "666666"
	sysUser.Password = encrypt.Sha1Encode(password + sysUser.Salt)
	if err := global.DB.Save(&sysUser).Error; err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, "重置密码失败")
	}
	return utils.SuccessNullMsg(c, fmt.Sprintf("重置密码成功，默认密码：%s", password))
}
