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

//sys_menu_btn - 菜单表
type SysMenuBtn struct {
	ID         string      `gorm:"column:id" json:"ID" form:"id" query:"id" valid:"Required;MaxSize(64)" validAlias:"编号"`                                     // 编号
	SysMenuID  string      `gorm:"column:sys_menu_id" json:"SysMenuID" form:"sys_menu_id" query:"sys_menu_id" valid:"Required;MaxSize(64)" validAlias:"菜单id"` // 菜单id
	Name       string      `gorm:"column:name" json:"Name" form:"name" query:"name" valid:"Required;MaxSize(100)" validAlias:"名称"`                            // 名称
	Permission null.String `gorm:"column:permission" json:"Permission" form:"permission" query:"permission" valid:"MaxSize(100)" validAlias:"权限标识"`           // 权限标识
	Remarks    null.String `gorm:"column:remarks" json:"Remarks" form:"remarks" query:"remarks" valid:"MaxSize(100)" validAlias:"备注信息"`                       // 备注信息
	CreatedBy  string      `gorm:"column:created_by" json:"CreatedBy" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`     // 创建者
	CreatedAt  time.Time   `gorm:"column:created_at" json:"CreatedAt" form:"created_at" query:"created_at" validAlias:"创建时间"`                                 // 创建时间
	UpdatedBy  string      `gorm:"column:updated_by" json:"UpdatedBy" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`     // 更新者
	UpdatedAt  time.Time   `gorm:"column:updated_at" json:"UpdatedAt" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                                 // 更新时间
	DeletedAt  null.Time   `gorm:"column:deleted_at" json:"DeletedAt" form:"deleted_at" query:"deleted_at" validAlias:"删除时间"`                                 // 删除时间
	Method     string      `gorm:"column:method" json:"Method" form:"method" query:"method" valid:"Required;MaxSize(20)" validAlias:"请求方式"`                   // 请求方式
	Path       string      `gorm:"column:path" json:"Path" form:"path" query:"path" valid:"Required;MaxSize(100)" validAlias:"请求路径"`                          // 请求路径
}

// TableName sets the insert table name for this struct type
func (s *SysMenuBtn) TableName() string {
	return "sys_menu_btn"
}
