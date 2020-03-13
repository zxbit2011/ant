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

//v_sys_user - VIEW
type VSysLog struct {
	SysLog
	Name     string `gorm:"column:name" json:"name" form:"name" query:"name" valid:"Required;MaxSize(50)" validAlias:"姓名"`                               // 姓名，Table[search]
	APIID    string `gorm:"column:api_id;default:'0'" json:"api_id" form:"api_id" query:"api_id" valid:"Required;MaxSize(100)" validAlias:"第三方API的用户ID"` // 第三方API的用户IDUserType  string      `gorm:"column:user_type" json:"user_type" form:"user_type" query:"user_type" valid:"Required;MaxSize(30)" validAlias:"用户类型"`                // 用户类型，sec：教务，teacher：教师，student：学生
	UserType string `gorm:"column:user_type" json:"user_type" form:"user_type" query:"user_type" valid:"Required;MaxSize(30)" validAlias:"用户类型"`         // 用户类型，sec：教务，teacher：教师，student：学生
}

// TableName sets the insert table name for this struct type
func (v *VSysLog) TableName() string {
	return "v_sys_log"
}
