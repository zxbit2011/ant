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

//sys_role_menu_btn - 角色-菜单
type SysRoleMenuBtn struct {
	ID           string `gorm:"column:id" json:"ID" form:"id" query:"id" valid:"Required;MaxSize(64)"`
	SysRoleID    string `gorm:"column:sys_role_id" json:"SysRoleID" form:"sys_role_id" query:"sys_role_id" valid:"Required;MaxSize(64)" validAlias:"角色编号"`                  // 角色编号
	SysMenuID    string `gorm:"column:sys_menu_id" json:"SysMenuID" form:"sys_menu_id" query:"sys_menu_id" valid:"Required;MaxSize(64)" validAlias:"按钮权限编号"`                // 按钮权限编号
	SysMenuBtnID string `gorm:"column:sys_menu_btn_id" json:"SysMenuBtnID" form:"sys_menu_btn_id" query:"sys_menu_btn_id" valid:"Required;MaxSize(64)" validAlias:"按钮权限编号"` // 按钮权限编号

}

// TableName sets the insert table name for this struct type
func (s *SysRoleMenuBtn) TableName() string {
	return "sys_role_menu_btn"
}
