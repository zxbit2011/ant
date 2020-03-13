package sys

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils"
)

// 日志表获取分页数据
func GetSysLogPage(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	//排序
	order := "created_at DESC"
	sortField := c.FormValue("sortField")
	sortOrder := c.FormValue("sortOrder")
	if sortField != "" && sortOrder != "" {
		order = fmt.Sprintf("%s %s", sortField, sortOrder)
	}
	db := global.DB.Model(&model.VSysLog{}).Where("company_id=?", loginInfo.CompanyId).Order(order)
	//条件搜索范例
	name := c.FormValue("name")
	createdBy := c.FormValue("created_by")
	t := c.FormValue("type")
	url := c.FormValue("url")
	ip := c.FormValue("ip")
	title := c.FormValue("title")
	startTime := c.FormValue("startTime")
	endTime := c.FormValue("endTime")
	if name != "" {
		db = db.Where("name like ?", fmt.Sprintf("%%%s%%", name))
	}
	if title != "" {
		db = db.Where("title like ?", fmt.Sprintf("%%%s%%", title))
	}
	if url != "" {
		db = db.Where("request_url like ?", fmt.Sprintf("%%%s%%", url))
	}
	if ip != "" {
		db = db.Where("remote_addr like ?", fmt.Sprintf("%%%s%%", ip))
	}
	if t != "" {
		db = db.Where("type = ?", t)
	}
	if createdBy != "" {
		db = db.Where("created_by=?", createdBy)
	}
	if startTime != "" && endTime != "" {
		db = db.Where("created_at>=? AND created_at<=?", startTime, endTime)
	}
	var count int
	var list []model.VSysLog
	var err error
	if err := db.Count(&count).Error; err != nil {
		global.Log.Error("GetSysLogPage error：", err)
		return utils.ErrorNull(c, "获取日志失败")
	}
	pageIndex := utils.GetPageIndex(c.FormValue("pageNo"))
	pageSize := utils.GetPageSize(c.FormValue("pageSize"))
	if err = db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("GetSysLogPage error：%v", err)
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
