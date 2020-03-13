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

//sys_menu - 菜单表
type SysMenu struct {
	ID          string      `gorm:"column:id;primary_key" json:"id" form:"id" query:"id" valid:"Required;MaxSize(64)" validAlias:"编号"`                                   // 编号
	ParentID    string      `gorm:"column:parent_id" json:"parent_id" form:"parent_id" query:"parent_id" valid:"Required;MaxSize(64)" validAlias:"父级编号"`                 // 父级编号
	RelationIds string      `gorm:"column:relation_ids" json:"relation_ids" form:"relation_ids" query:"relation_ids" valid:"Required;MaxSize(2000)" validAlias:"所有父级编号"` // 所有父级编号
	Name        string      `gorm:"column:name" json:"name" form:"name" query:"name" valid:"Required;MaxSize(100)" validAlias:"名称"`                                      // 名称
	Sort        float64     `gorm:"column:sort" json:"sort" form:"sort" query:"sort" valid:"Required" validAlias:"排序"`                                                   // 排序
	Href        null.String `gorm:"column:href" json:"href" form:"href" query:"href" valid:"MaxSize(100)" validAlias:"链接"`                                               // 链接
	Target      null.String `gorm:"column:target" json:"target" form:"target" query:"target" valid:"MaxSize(100)" validAlias:"目标"`                                       // 目标
	Icon        null.String `gorm:"column:icon" json:"icon" form:"icon" query:"icon" valid:"MaxSize(100)" validAlias:"图标"`                                               // 图标
	IsShow      string      `gorm:"column:is_show" json:"is_show" form:"is_show" query:"is_show" valid:"MaxSize(1)" validAlias:"是否在菜单中显示"`                      // 是否在菜单中显示
	Permission  null.String `gorm:"column:permission" json:"permission" form:"permission" query:"permission" valid:"MaxSize(200)" validAlias:"权限标识"`                     // 权限标识
	CreatedBy   string      `gorm:"column:created_by" json:"created_by" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`              // 创建者
	CreatedAt   time.Time   `gorm:"column:created_at" json:"created_at" form:"created_at" query:"created_at" validAlias:"创建时间"`                                          // 创建时间
	UpdatedBy   string      `gorm:"column:updated_by" json:"updated_by" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`              // 更新者
	UpdatedAt   time.Time   `gorm:"column:updated_at" json:"updated_at" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                                          // 更新时间
	Remarks     null.String `gorm:"column:remarks" json:"remarks" form:"remarks" query:"remarks" valid:"MaxSize(255)" validAlias:"备注信息"`                                 // 备注信息
	DeletedAt   null.Time   `gorm:"column:deleted_at" json:"deleted_at" form:"deleted_at" query:"deleted_at" validAlias:"删除时间"`                                          // 删除时间

}

// TableName sets the insert table name for this struct type
func (s *SysMenu) TableName() string {
	return "sys_menu"
}
