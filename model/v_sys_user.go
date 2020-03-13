package model

type VSysUser struct {
	SysUser
	SysCompanyId   string `gorm:"column:sys_company_id" json:"sys_company_id"`
	SysCompanyName string `gorm:"column:sys_company_name" json:"sys_company_name"`
	DataSource     string `gorm:"column:data_source" json:"data_source"`
	RelationIds    string `gorm:"column:relation_ids" json:"relation_ids"`
	OfficeParentId string `gorm:"column:office_parent_id" json:"office_parent_id"`
}

func (s *VSysUser) TableName() string {
	return "v_sys_user"
}
