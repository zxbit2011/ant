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

//sys_dict - 字典表
type SysDict struct {
	ID          string    `gorm:"column:id" json:"ID" form:"id" query:"id" valid:"Required;MaxSize(64)" validAlias:"编号"`                                      // 编号
	Value       string    `gorm:"column:value" json:"Value" form:"value" query:"value" valid:"Required;MaxSize(100)" validAlias:"数据值"`                        // 数据值
	Label       string    `gorm:"column:label" json:"Label" form:"label" query:"label" valid:"Required;MaxSize(100)" validAlias:"标签名"`                        // 标签名
	Type        string    `gorm:"column:type" json:"Type" form:"type" query:"type" valid:"Required;MaxSize(100)" validAlias:"类型"`                             // 类型
	Description string    `gorm:"column:description" json:"Description" form:"description" query:"description" valid:"Required;MaxSize(100)" validAlias:"描述"` // 描述
	Sort        float64   `gorm:"column:sort" json:"Sort" form:"sort" query:"sort" valid:"Required" validAlias:"排序（升序）"`                                      // 排序（升序）
	ParentID    string    `gorm:"column:parent_id" json:"ParentID" form:"parent_id" query:"parent_id" valid:"Required;MaxSize(64)" validAlias:"父级编号"`         // 父级编号
	CreatedBy   string    `gorm:"column:created_by" json:"CreatedBy" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`      // 创建者
	CreatedAt   time.Time `gorm:"column:created_at" json:"CreatedAt" form:"created_at" query:"created_at" validAlias:"创建时间"`                                  // 创建时间
	UpdatedBy   string    `gorm:"column:updated_by" json:"UpdatedBy" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`      // 更新者
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"UpdatedAt" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                                  // 更新时间
	DeletedAt   null.Time `gorm:"column:deleted_at" json:"DeletedAt" form:"deleted_at" query:"deleted_at" validAlias:"删除时间"`                                  // 删除时间

}

// TableName sets the insert table name for this struct type
func (s *SysDict) TableName() string {
	return "sys_dict"
}
