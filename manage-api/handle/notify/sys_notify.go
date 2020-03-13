package notify

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/manage-api/handle"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils"
	"github.com/zxbit2011/ant/utils/enum"
	"github.com/zxbit2011/ant/utils/validation"
	"strings"
)

// 获取通知消息详情
func GetSysNotify(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sn model.SysNotify
	if err := global.DB.Model(&model.SysNotify{}).First(&sn, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, "该通知消息不存在")
		}
		global.Log.Error("GetSysNotify error：", err)
		return utils.ErrorNull(c, "获取通知消息失败")
	}
	if err := handle.PowerCheck(loginInfo, sn.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}
	return utils.SuccessNullMsg(c, sn)
}

// 获取通知消息详情页
func GetSysNotifyInfo(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sn model.SysNotify
	if err := global.DB.First(&sn, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, "该通知消息不存在")
		}
		global.Log.Error("GetSysNotifyInfo error：", err)
		return utils.ErrorNull(c, "获取通知消息失败")
	}
	if err := handle.PowerCheck(loginInfo, sn.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}
	var snf []model.SysNotifyFile
	if err := global.DB.Find(&snf, "sys_notify_id=?", sn.ID).Error; err != nil {
		global.Log.Error("GetSysNotifyInfo error：", err)
		return utils.ErrorNull(c, "获取通知消息文件失败")
	}
	return utils.SuccessNullMsg(c, map[string]interface{}{
		"notify":     sn,
		"notifyFile": snf,
	})
}

// 获取消息通知
func GetSysNotifyList(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	limit := c.FormValue("limit")
	order := "created_at DESC"
	var list []model.VSysNotify
	db := global.DB.Order(order)
	//【数据权限控制】判断角色，如果是超级管理员则不进行验证
	if !loginInfo.IsSuperAdmin {
		db = db.Where("company_id=? AND relation_ids LIKE ?", loginInfo.CompanyId,
			fmt.Sprintf("%s%%", loginInfo.OfficeRelationIds))
	}
	if limit != "" && utils.IsValidNumber(limit) {
		db = db.Limit(10)
	}
	if err := db.Find(&list).Error; err != nil {
		global.Log.Error("GetProjectList error：", err)
		return utils.ErrorNull(c, "获取消息通知失败")
	}
	return utils.SuccessNullMsg(c, list)
}

/**
通知消息分页
*/
func GetSysNotifyPage(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	title := c.FormValue("title")
	typ := c.FormValue("type")
	status := c.FormValue("status")
	startTime := c.FormValue("startTime")
	endTime := c.FormValue("endTime")
	//排序
	order := "created_at DESC"
	sortField := c.FormValue("sortField")
	sortOrder := c.FormValue("sortOrder")
	if sortField != "" && sortOrder != "" {
		order = fmt.Sprintf("%s %s", sortField, sortOrder)
	}
	db := global.DB.Model(&model.VSysNotify{}).Order(order)

	//【数据权限控制】判断角色，如果是超级管理员则不进行验证
	if !loginInfo.IsSuperAdmin {
		db = db.Where("company_id=? AND relation_ids LIKE ?", loginInfo.CompanyId,
			fmt.Sprintf("%s%%", loginInfo.OfficeRelationIds))
	}
	//通知消息名称搜索
	if title != "" {
		where := "%" + title + "%"
		db = db.Where("title like ?", where)
	}
	//通知消息类型
	if typ != "" {
		db = db.Where("type = ?", typ)
	}
	//通知消息发布状态
	if status != "" && utils.IsValidNumber(status) {
		db = db.Where("status = ?", status)
	}
	//通知消息发表起止时间
	if startTime != "" && endTime != "" && utils.IsValidDateTime(startTime) && utils.IsValidDateTime(endTime) {
		db = db.Where("created_at > ? AND created_at<?", startTime, endTime)
	}
	var count int
	var list []model.VSysNotify
	if err := db.Count(&count).Error; err != nil {
		global.Log.Error("GetSysNotifyPage error：", err)
		return utils.ErrorNull(c, "获取消息失败")
	}
	pageIndex := utils.GetPageIndex(c.FormValue("pageNo"))
	pageSize := utils.GetPageSize(c.FormValue("pageSize"))
	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	return utils.SuccessNullMsg(c, &utils.PageData{
		PageIndex:  pageIndex,
		PageSize:   pageSize,
		Count:      count,
		PageNumber: utils.GetPageNumber(count, pageSize),
		Data:       list,
	})
}

// 公共保存通知消息类型
func SaveSysNotify(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	idStr := c.FormValue("id")
	isEditFlag := false
	var sn model.SysNotify
	if idStr != "" {
		if err := global.DB.First(&sn, "id=?", idStr).Error; err != nil {
			return utils.ErrorNull(c, "操作无效")
		}
		if err := handle.PowerCheck(loginInfo, sn.CreatedBy); err != nil {
			return utils.ErrorNull(c, err.Error())
		}
		isEditFlag = true
	}
	if err := new(utils.CustomBinder).Bind(&sn, c); err != nil {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}

	//业务逻辑
	sn.UpdatedBy = loginInfo.ID
	if !isEditFlag {
		sn.ID = utils.ID()
		sn.Status = enum.NotRelease
		sn.CreatedBy = loginInfo.ID
		sn.OfficeID = loginInfo.OfficeId
		sn.CompanyID = loginInfo.CompanyId

	}

	//参数验证
	errs := validation.Valid(&sn)
	if len(errs) > 0 {
		return utils.Error(c, "参数验证失败", errs)
	}

	// 开始事务
	tx := global.DB.Begin()
	//保存文件
	fileList := c.FormValue("fileList")
	if fileList != "" {
		var formFile []model.FormFile
		if err := json.Unmarshal([]byte(fileList), &formFile); err != nil {
			global.Log.Error("SaveSysNotify Unmarshal fileList error", err)
			return utils.ErrorNull(c, "文件上传失败")
		}
		for _, v := range formFile {
			if err := tx.Create(&model.SysNotifyFile{
				ID:          utils.ID(),
				SysNotifyID: sn.ID,
				Name:        v.Name,
				URL:         v.Url,
				CreatedBy:   loginInfo.ID,
				UpdatedBy:   loginInfo.ID,
			}).Error; err != nil {
				tx.Rollback()
				global.Log.Error("SaveSysNotify Create error", err)
				return utils.ErrorNull(c, "添加失败")
			}
		}
	}
	//数据保存
	if !isEditFlag {
		if err := tx.Create(&sn).Error; err != nil {
			tx.Rollback()
			global.Log.Error("SaveSysNotify Create error", err)
			return utils.ErrorNull(c, "添加失败")
		}
	} else {
		if err := tx.Save(&sn).Error; err != nil {
			tx.Rollback()
			global.Log.Error("SaveSysNotify Save error", err)
			return utils.ErrorNull(c, "修改失败")
		}
	}
	tx.Commit()
	return utils.Success(c, "保存成功", sn)
}

/**
通知消息审核
*/
func ReleaseSysNotify(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var pt model.VSysNotify
	if err := global.DB.Model(&model.VSysNotify{}).First(&pt, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, "该通知消息不存在")
		}
		global.Log.Error("get VSysNotify error：", err)
		return utils.ErrorNull(c, "获取通知消息失败")
	}
	if !loginInfo.IsSuperAdmin && !strings.Contains(pt.RelationIds, loginInfo.OfficeRelationIds) {
		return utils.ErrorNull(c, "无权限操作此通知消息")
	}
	var reviewStatus, reviewMsg string
	if pt.Status == enum.NotRelease {
		reviewStatus = enum.Release
		reviewMsg = "通知消息已发布！"
	} else {
		reviewStatus = enum.NotRelease
		reviewMsg = "通知消息已取消发布！"
	}
	if err := global.DB.Model(&pt).Updates(map[string]interface{}{"status": reviewStatus, "updated_by": loginInfo.ID}).Error; err != nil {
		return utils.ErrorNull(c, "操作失败")
	}
	return utils.SuccessNull(c, reviewMsg)
}

// 通知消息表批量删除数据
func DelSysNotify(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	ids := c.FormValue("ids")
	if ids == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	idsArr := strings.Split(ids, ",")
	//查询所有的通知消息进行权限校验
	db := global.DB.Model(&model.VSysNotify{}).Where("id IN (?)", idsArr)
	//【数据权限控制】判断角色，如果是超级管理员则不进行验证
	if !loginInfo.IsSuperAdmin {
		db = db.Where("company_id=? AND relation_ids LIKE ?", loginInfo.CompanyId,
			fmt.Sprintf("%s%%", loginInfo.OfficeRelationIds))
	}
	var count int
	if err := db.Count(&count).Error; err != nil {
		global.Log.Error("DelSysNotify error：", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	if count != len(idsArr) {
		return utils.ErrorNull(c, "删除失败，无权限或一些通知消息不存在")
	}
	if err := global.DB.Where("id IN (?)", idsArr).Delete(&model.SysNotify{}).Error; err != nil {
		global.Log.Error("DelSysNotify error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
