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

//sys_user_role - 用户-角色
type SysUserRole struct {
	ID        string `gorm:"column:id" json:"ID" form:"id" query:"id" valid:"Required;MaxSize(64)"`
	SysUserID string `gorm:"column:sys_user_id" json:"SysUserID" form:"sys_user_id" query:"sys_user_id" valid:"Required;MaxSize(64)" validAlias:"用户编号"` // 用户编号
	SysRoleID string `gorm:"column:sys_role_id" json:"SysRoleID" form:"sys_role_id" query:"sys_role_id" valid:"Required;MaxSize(64)" validAlias:"角色编号"` // 角色编号

}

// TableName sets the insert table name for this struct type
func (s *SysUserRole) TableName() string {
	return "sys_user_role"
}
