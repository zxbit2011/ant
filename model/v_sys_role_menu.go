package model

type VSysRoleMenu struct {
	SysMenu
	SysRoleId string `gorm:"column:sys_role_id" json:"sys_role_id"`
}

func (s *VSysRoleMenu) TableName() string {
	return "v_sys_role_menu"
}
