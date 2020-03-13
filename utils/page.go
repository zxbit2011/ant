package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zxbit2011/ant/utils/convert"
	"math"
)

type PageTable struct {
	Fields    string
	Table     string
	Where     string
	PageIndex int
	PageSize  int
	Order     string
	GroupBy   string
}

type PageData struct {
	PageIndex  int         `json:"pageNo"`
	PageSize   int         `json:"pageSize"`
	PageNumber int         `json:"pageNum"`
	Count      int         `json:"totalCount"`
	Data       interface{} `json:"data"`
	ExtendData interface{} `json:"extendData"`
}

const PageSize = 10

func GetPageIndex(pi string) int {
	pageIndex := 1
	if pi != "" && IsValidNumber(pi) {
		pageIndex = convert.MustInt(pi)
	}
	return pageIndex
}

func GetPageSize(ps string) int {
	pageSize := PageSize
	if ps != "" && IsValidNumber(ps) {
		pageSize = convert.MustInt(ps)
	}
	if pageSize > 100 {
		//防止大量数据查询
		pageSize = 100
	}
	return pageSize
}

type Count struct {
	count int
}

func GetPageNumber(count, pageSize int) int {
	return int(math.Ceil(float64(count) / float64(pageSize)))
}

func QueryPageMaps(db *gorm.DB, page *PageTable, args ...interface{}) (*PageData, error) {
	return QueryPage(db, page, []map[string]interface{}{}, args...)
}

func QueryPage(db *gorm.DB, page *PageTable, data interface{}, args ...interface{}) (*PageData, error) {
	if page.Where != "" {
		page.Where = " WHERE " + page.Where
	}
	if page.Order != "" {
		page.Order = " ORDER BY " + page.Order
	}
	var sql string
	if page.GroupBy != "" {
		page.GroupBy = " GROUP BY " + page.GroupBy
		sql = fmt.Sprintf("SELECT COUNT(1) count FROM (SELECT %s FROM %s %s %s) c", page.Fields, page.Table, page.Where, page.GroupBy)
	} else {
		sql = fmt.Sprintf("SELECT COUNT(1) count FROM  %s %s", page.Table, page.Where)
	}
	println(fmt.Sprintf("QueryPage COUNT sql：%s", sql))
	var c Count
	err := db.Raw(sql, args...).Scan(&c).Error
	if err != nil {
		return nil, err
	}
	sql = fmt.Sprintf("SELECT %s FROM %s %s %s %s limit %d,%d", page.Fields, page.Table, page.Where, page.GroupBy, page.Order, (page.PageIndex-1)*page.PageSize, page.PageSize)
	println(fmt.Sprintf("QueryPage sql：%s", sql))
	err = db.Raw(sql, args...).Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return &PageData{
		PageIndex:  page.PageIndex,
		PageSize:   page.PageSize,
		Count:      c.count,
		Data:       data,
		PageNumber: int(math.Ceil(float64(c.count) / float64(page.PageSize))),
	}, nil
}
