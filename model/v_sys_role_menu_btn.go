package model

type VSysRoleMenuBtn struct {
	SysMenuBtn
	SysRoleId string `gorm:"column:sys_role_id" json:"sys_role_id"`
}

func (s *VSysRoleMenuBtn) TableName() string {
	return "v_sys_role_menu_btn"
}
