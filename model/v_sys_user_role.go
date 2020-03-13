package model

type VSysUserRole struct {
	SysRole
	SysUserId string `gorm:"column:sys_user_id" json:"sys_user_id"`
}

func (s *VSysUserRole) TableName() string {
	return "v_sys_user_role"
}
