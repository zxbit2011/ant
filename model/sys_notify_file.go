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

//sys_notify_file - 通知记录文件
type SysNotifyFile struct {
	ID          string    `gorm:"column:id;primary_key" json:"id" form:"id" query:"id" valid:"Required;MaxSize(64)"`
	SysNotifyID string    `gorm:"column:sys_notify_id" json:"sys_notify_id" form:"sys_notify_id" query:"sys_notify_id" valid:"Required;MaxSize(64)" validAlias:"通知消息id"` // 通知消息id
	Name        string    `gorm:"column:name" json:"name" form:"name" query:"name" valid:"Required;MaxSize(100)" validAlias:"文件名称"`                                      // 文件名称
	URL         string    `gorm:"column:url" json:"url" form:"url" query:"url" valid:"Required;MaxSize(255)" validAlias:"文件地址"`                                          // 文件地址
	CreatedBy   string    `gorm:"column:created_by" json:"created_by" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`                // 创建者
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at" form:"created_at" query:"created_at" validAlias:"创建时间"`                                            // 创建时间
	UpdatedBy   string    `gorm:"column:updated_by" json:"updated_by" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`                // 更新者
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                                            // 更新时间
	DeletedAt   null.Time `gorm:"column:deleted_at" json:"deleted_at" form:"deleted_at" query:"deleted_at" validAlias:"删除时间"`                                            // 删除时间
}

// TableName sets the insert table name for this struct type
func (p *SysNotifyFile) TableName() string {
	return "sys_notify_file"
}
