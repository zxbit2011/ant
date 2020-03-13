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

//sys_notify - 通知通告
type SysNotify struct {
	ID        string      `gorm:"column:id;primary_key" json:"id" form:"id" query:"id" valid:"Required;MaxSize(64)" validAlias:"编号"`                       // 编号
	Type      string      `gorm:"column:type" json:"type" form:"type" query:"type" valid:"MaxSize(1)" validAlias:"类型"`                                     // 类型
	ObjId     null.String `gorm:"column:obj_id" json:"obj_id" form:"obj_id" query:"obj_id" valid:"MaxSize(64)" validAlias:"发送对象id"`                        // 类型
	Title     string      `gorm:"column:title" json:"title" form:"title" query:"title" valid:"MaxSize(200)" validAlias:"标题"`                               // 标题
	Content   string      `gorm:"column:content" json:"content" form:"content" query:"content" valid:"MaxSize(2000)" validAlias:"内容"`                      // 内容
	Status    string      `gorm:"column:status" json:"status" form:"status" query:"status" valid:"MaxSize(1)" validAlias:"状态"`                             // 状态
	CreatedBy string      `gorm:"column:created_by" json:"created_by" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`  // 创建者
	CreatedAt time.Time   `gorm:"column:created_at" json:"created_at" form:"created_at" query:"created_at" validAlias:"创建时间"`                              // 创建时间
	UpdatedBy string      `gorm:"column:updated_by" json:"updated_by" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`  // 更新者
	UpdatedAt time.Time   `gorm:"column:updated_at" json:"updated_at" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                              // 更新时间
	DeletedAt null.Time   `gorm:"column:deleted_at" json:"deleted_at" form:"deleted_at" query:"deleted_at" validAlias:"删除时间"`                              // 删除时间
	CompanyID string      `gorm:"column:company_id" json:"company_id" form:"company_id" query:"company_id" valid:"Required;MaxSize(64)" validAlias:"归属公司"` // 归属公司
	OfficeID  string      `gorm:"column:office_id" json:"office_id" form:"office_id" query:"office_id" valid:"Required;MaxSize(64)" validAlias:"归属部门"`     // 归属部门
}

// TableName sets the insert table name for this struct type
func (s *SysNotify) TableName() string {
	return "sys_notify"
}
