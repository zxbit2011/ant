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

//sys_role - 角色表
type SysRole struct {
	ID           string      `gorm:"column:id;primary_key" json:"id" form:"id" query:"id" valid:"Required;MaxSize(64)" validAlias:"编号"`                                       // 编号
	SysCompanyID string      `gorm:"column:sys_company_id" json:"sys_company_id" form:"sys_company_id" query:"sys_company_id" valid:"Required;MaxSize(64)" validAlias:"归属机构"` // 归属机构
	Name         string      `gorm:"column:name" json:"name" form:"name" query:"name" valid:"Required;MaxSize(100)" validAlias:"角色名称"`                                        // 角色名称
	Enname       null.String `gorm:"column:enname" json:"enname" form:"enname" query:"enname" valid:"MaxSize(255)" validAlias:"英文名称"`                                         // 英文名称
	RoleType     string      `gorm:"column:role_type;default:'0'" json:"role_type" form:"role_type" query:"role_type" valid:"Required;MaxSize(1)" validAlias:"角色类型"`          // 角色类型，2、超级管理员，1、管理员，0、普通账户
	DataScope    null.String `gorm:"column:data_scope" json:"data_scope" form:"data_scope" query:"data_scope" valid:"MaxSize(1)" validAlias:"数据范围"`                           // 数据范围
	Useable      null.String `gorm:"column:useable" json:"useable" form:"useable" query:"useable" valid:"MaxSize(1)" validAlias:"是否可用"`                                       // 是否可用，1、可用，0、不可用
	CreatedBy    string      `gorm:"column:created_by" json:"created_by" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`                  // 创建者
	CreatedAt    time.Time   `gorm:"column:created_at" json:"created_at" form:"created_at" query:"created_at" validAlias:"创建时间"`                                              // 创建时间
	UpdatedBy    string      `gorm:"column:updated_by" json:"updated_by" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`                  // 更新者
	UpdatedAt    time.Time   `gorm:"column:updated_at" json:"updated_at" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                                              // 更新时间
	Remarks      null.String `gorm:"column:remarks" json:"remarks" form:"remarks" query:"remarks" valid:"MaxSize(255)" validAlias:"备注信息"`                                     // 备注信息
	DeletedAt    null.Time   `gorm:"column:deleted_at" json:"deleted_at" form:"deleted_at" query:"deleted_at" validAlias:"删除时间"`                                              // 删除时间
}

// TableName sets the insert table name for this struct type
func (s *SysRole) TableName() string {
	return "sys_role"
}
