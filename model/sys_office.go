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

//sys_office - 机构表
type SysOffice struct {
	ID           string      `gorm:"column:id;primary_key" json:"id" form:"id" query:"id" valid:"Required;MaxSize(64)" validAlias:"编号"`                                       // 编号
	SysCompanyID string      `gorm:"column:sys_company_id" json:"sys_company_id" form:"sys_company_id" query:"sys_company_id" valid:"Required;MaxSize(64)" validAlias:"公司id"` // 公司id
	APIID        string      `gorm:"column:api_id" json:"api_id" form:"api_id" query:"api_id" valid:"Required;MaxSize(64)" validAlias:"第三方id"`                                // 第三方id
	ParentID     string      `gorm:"column:parent_id" json:"parent_id" form:"parent_id" query:"parent_id" valid:"Required;MaxSize(64)" validAlias:"父级编号"`                     // 父级编号
	RelationIds  string      `gorm:"column:relation_ids" json:"relation_ids" form:"relation_ids" query:"relation_ids" valid:"Required;MaxSize(2000)" validAlias:"所有父级编号"`     // 所有父级编号
	Name         string      `gorm:"column:name" json:"name" form:"name" query:"name" valid:"Required;MaxSize(100)" validAlias:"名称"`                                          // 名称
	Sort         int         `gorm:"column:sort;default:'0'" json:"sort" form:"sort" query:"sort" valid:"Required" validAlias:"排序"`                                           // 排序
	AreaID       string      `gorm:"column:area_id" json:"area_id" form:"area_id" query:"area_id" valid:"Required;MaxSize(64)" validAlias:"归属区域"`                             // 归属区域
	Code         null.String `gorm:"column:code" json:"code" form:"code" query:"code" valid:"MaxSize(100)" validAlias:"区域编码"`                                                 // 区域编码
	Type         string      `gorm:"column:type" json:"type" form:"type" query:"type" valid:"Required;MaxSize(1)" validAlias:"机构类型"`                                          // 机构类型，0：教务，1：学校
	Grade        string      `gorm:"column:grade" json:"grade" form:"grade" query:"grade" valid:"Required;MaxSize(1)" validAlias:"机构等级"`                                      // 机构等级
	Address      null.String `gorm:"column:address" json:"address" form:"address" query:"address" valid:"MaxSize(255)" validAlias:"联系地址"`                                     // 联系地址
	ZipCode      null.String `gorm:"column:zip_code" json:"zip_code" form:"zip_code" query:"zip_code" valid:"MaxSize(100)" validAlias:"邮政编码"`                                 // 邮政编码
	Master       null.String `gorm:"column:master" json:"master" form:"master" query:"master" valid:"MaxSize(100)" validAlias:"负责人"`                                          // 负责人
	Phone        null.String `gorm:"column:phone" json:"phone" form:"phone" query:"phone" valid:"MaxSize(50)" validAlias:"电话"`                                                // 电话
	Fax          null.String `gorm:"column:fax" json:"fax" form:"fax" query:"fax" valid:"MaxSize(50)" validAlias:"传真"`                                                        // 传真
	Email        null.String `gorm:"column:email" json:"email" form:"email" query:"email" valid:"MaxSize(100)" validAlias:"邮箱"`                                               // 邮箱
	Useable      null.String `gorm:"column:useable;default:'1'" json:"useable" form:"useable" query:"useable" valid:"MaxSize(1)" validAlias:"是否启用"`                           // 是否启用
	CreatedBy    string      `gorm:"column:created_by" json:"CreatedBy" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`                 // 创建者
	CreatedAt    time.Time   `gorm:"column:created_at" json:"CreatedAt" form:"created_at" query:"created_at" validAlias:"创建时间"`                                             // 创建时间
	UpdatedBy    string      `gorm:"column:updated_by" json:"UpdatedBy" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`                 // 更新者
	UpdatedAt    time.Time   `gorm:"column:updated_at" json:"UpdatedAt" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                                             // 更新时间
	Remarks      null.String `gorm:"column:remarks" json:"Remarks" form:"remarks" query:"remarks" valid:"MaxSize(100)" validAlias:"备注信息"`                                    // 备注信息
	DeletedAt    null.Time   `gorm:"column:deleted_at" json:"DeletedAt" form:"deleted_at" query:"deleted_at" validAlias:"删除时间"`                                             // 删除时间
}

// TableName sets the insert table name for this struct type
func (s *SysOffice) TableName() string {
	return "sys_office"
}
