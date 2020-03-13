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

//file_log -
type FileLog struct {
	ID        string    `gorm:"column:id" json:"ID" form:"ID" query:"id" valid:"Required;MaxSize(64)"`
	Name      string    `gorm:"column:name" json:"Name" form:"Name" query:"name" valid:"Required;MaxSize(100)" validAlias:"文件名称"`                     // 文件名称
	Path      string    `gorm:"column:path" json:"Path" form:"Path" query:"path" valid:"Required;MaxSize(100)" validAlias:"路径"`                       // 路径
	Size      int64     `gorm:"column:size" json:"Size" form:"Size" query:"size" valid:"Required" validAlias:"大小"`                                    // 大小
	Ext       string    `gorm:"column:ext" json:"Ext" form:"Ext" query:"ext" valid:"Required;MaxSize(50)" validAlias:"扩展名"`                           // 扩展名
	CreatedBy string    `gorm:"column:created_by" json:"CreatedBy" form:"CreatedBy" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"` // 创建者
	CreatedAt time.Time `gorm:"column:created_at" json:"CreatedAt" form:"CreatedAt" query:"created_at" validAlias:"创建时间"`                             // 创建时间
	UpdatedBy string    `gorm:"column:updated_by" json:"UpdatedBy" form:"UpdatedBy" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"` // 更新者
	UpdatedAt time.Time `gorm:"column:updated_at" json:"UpdatedAt" form:"UpdatedAt" query:"updated_at" validAlias:"更新时间"`                             // 更新时间
	DeletedAt null.Time `gorm:"column:deleted_at" json:"DeletedAt" form:"DeletedAt" query:"deleted_at" validAlias:"删除时间"`                             // 删除时间
	IP        string    `gorm:"column:ip" json:"IP" form:"IP" query:"ip" valid:"Required;MaxSize(20)" validAlias:"ip"`                                // ip
}

// TableName sets the insert table name for this struct type
func (f *FileLog) TableName() string {
	return "file_log"
}
