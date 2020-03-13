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

//sys_role_menu - 角色-菜单
type SysRoleMenu struct {
	ID        string `gorm:"column:id" json:"ID" form:"id" query:"id" valid:"Required;MaxSize(64)"`
	SysRoleID string `gorm:"column:sys_role_id" json:"SysRoleID" form:"sys_role_id" query:"sys_role_id" valid:"Required;MaxSize(64)" validAlias:"角色编号"` // 角色编号
	SysMenuID string `gorm:"column:sys_menu_id" json:"SysMenuID" form:"sys_menu_id" query:"sys_menu_id" valid:"Required;MaxSize(64)" validAlias:"菜单编号"` // 菜单编号

}

// TableName sets the insert table name for this struct type
func (s *SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
