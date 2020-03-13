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

//sys_area -
type SysArea struct {
	ID        int         `gorm:"column:id;primary_key" json:"id" form:"id" query:"id" valid:"Required" validAlias:"ID"`                                   // ID
	Name      string      `gorm:"column:name;default:''" json:"name" form:"name" query:"name" valid:"Required;MaxSize(50)" validAlias:"地区名"`               // 地区名
	ParentID  int         `gorm:"column:parent_id" json:"parent_id" form:"parent_id" query:"parent_id" valid:"Required" validAlias:"父ID"`                  // 父ID
	Zipcode   null.String `gorm:"column:zipcode" json:"zipcode" form:"zipcode" query:"zipcode" valid:"MaxSize(10)" validAlias:"邮编"`                        // 邮编
	Sort      null.Int    `gorm:"column:sort;default:'50'" json:"sort" form:"sort" query:"sort" validAlias:"排序"`                                           // 排序
	CreatedBy string      `gorm:"column:created_by" json:"created_by" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`  // 创建者
	CreatedAt time.Time   `gorm:"column:created_at" json:"created_at" form:"created_at" query:"created_at" validAlias:"创建时间"`                              // 创建时间
	UpdatedBy string      `gorm:"column:updated_by" json:"updated_by" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`  // 更新者
	UpdatedAt time.Time   `gorm:"column:updated_at" json:"updated_at" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                              // 更新时间
	Remarks   string      `gorm:"column:remarks;default:''" json:"remarks" form:"remarks" query:"remarks" valid:"Required;MaxSize(255)" validAlias:"备注信息"` // 备注信息
	DeletedAt null.Time   `gorm:"column:deleted_at" json:"deleted_at" form:"deleted_at" query:"deleted_at" valid:"Required" validAlias:"删除时间"`             // 删除时间

}

// TableName sets the insert table name for this struct type
func (s *SysArea) TableName() string {
	return "sys_area"
}
