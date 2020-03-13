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
type VSysNotify struct {
	SysNotify
	RelationIds string `gorm:"column:relation_ids" json:"relation_ids"`
	OfficeName  string `gorm:"column:office_name" json:"office_name"`
}

// TableName sets the insert table name for this struct type
func (s *VSysNotify) TableName() string {
	return "v_sys_notify"
}
