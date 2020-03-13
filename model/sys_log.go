package model

import (
	"database/sql"
	"github.com/guregu/null"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

//sys_log - 日志表
type SysLog struct {
	ID         string    `gorm:"column:id;primary_key" json:"id" form:"id" query:"id" valid:"Required;MaxSize(64)" validAlias:"编号"`                              // 编号
	Type       string    `gorm:"column:type;default:'1'" json:"type" form:"type" query:"type" valid:"Required;MaxSize(1)" validAlias:"日志类型"`                     // 日志类型
	Title      string    `gorm:"column:title;default:''" json:"title" form:"title" query:"title" valid:"Required;MaxSize(255)" validAlias:"日志标题"`                // 日志标题
	CreatedBy  string    `gorm:"column:created_by" json:"created_by" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`         // 创建者
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at" form:"created_at" query:"created_at" validAlias:"创建时间"`                                     // 创建时间
	RemoteAddr string    `gorm:"column:remote_addr" json:"remote_addr" form:"remote_addr" query:"remote_addr" valid:"Required;MaxSize(255)" validAlias:"操作IP地址"` // 操作IP地址
	UserAgent  string    `gorm:"column:user_agent" json:"user_agent" form:"user_agent" query:"user_agent" valid:"Required;MaxSize(255)" validAlias:"用户代理"`       // 用户代理
	RequestURI string    `gorm:"column:request_url" json:"request_url" form:"request_url" query:"request_url" valid:"Required;MaxSize(255)" validAlias:"请求URI"`  // 请求URI
	Method     string    `gorm:"column:method" json:"method" form:"method" query:"method" valid:"Required;MaxSize(5)" validAlias:"操作方式"`                         // 操作方式
	Params     string    `gorm:"column:params" json:"params" form:"params" query:"params" valid:"MaxSize(65535)" validAlias:"操作提交的数据"`                           // 操作提交的数据
	Exception  string    `gorm:"column:exception" json:"exception" form:"exception" query:"exception" valid:"MaxSize(65535)" validAlias:"异常信息"`                  // 异常信息

}

// TableName sets the insert table name for this struct type
func (s *SysLog) TableName() string {
	return "sys_log"
}
