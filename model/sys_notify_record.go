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

//sys_notify_record - 通知通告发送记录
type SysNotifyRecord struct {
	ID          string    `gorm:"column:id;primary_key" json:"id" form:"id" query:"id" valid:"Required;MaxSize(64)" validAlias:"编号"`                                     // 编号
	SysNotifyID string    `gorm:"column:sys_notify_id" json:"sys_notify_id" form:"sys_notify_id" query:"sys_notify_id" valid:"Required;MaxSize(64)" validAlias:"通知通告ID"` // 通知通告ID
	AccountID   string    `gorm:"column:account_id" json:"account_id" form:"account_id" query:"account_id" valid:"Required;MaxSize(64)" validAlias:"接受人"`                // 接受人
	ReadFlag    string    `gorm:"column:read_flag;default:'0'" json:"read_flag" form:"read_flag" query:"read_flag" valid:"Required;MaxSize(1)" validAlias:"阅读标记"`        // 阅读标记
	ReadAt      null.Time `gorm:"column:read_at" json:"read_at" form:"read_at" query:"read_at" validAlias:"阅读时间"`                                                        // 阅读时间
	CreatedBy   string    `gorm:"column:created_by" json:"created_by" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`                // 创建者
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at" form:"created_at" query:"created_at" validAlias:"创建时间"`                                            // 创建时间
	UpdatedBy   string    `gorm:"column:updated_by" json:"updated_by" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`                // 更新者
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                                            // 更新时间
	DeletedAt   null.Time `gorm:"column:deleted_at" json:"deleted_at" form:"deleted_at" query:"deleted_at" validAlias:"删除时间"`                                            // 删除时间

}

// TableName sets the insert table name for this struct type
func (s *SysNotifyRecord) TableName() string {
	return "sys_notify_record"
}
