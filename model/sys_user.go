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

//sys_user - 用户表
type SysUser struct {
	ID        string      `gorm:"column:id;primary_key" json:"id" form:"id" query:"id" valid:"Required;MaxSize(64)" validAlias:"编号"`                                  // 编号
	CompanyID string      `gorm:"column:company_id" json:"company_id" form:"company_id" query:"company_id" valid:"Required;MaxSize(64)" validAlias:"归属公司"`            // 归属公司
	OfficeID  string      `gorm:"column:office_id" json:"office_id" form:"office_id" query:"office_id" valid:"Required;MaxSize(64)" validAlias:"归属部门"`                // 归属部门
	Username  string      `gorm:"column:username" json:"username" form:"username" query:"username" valid:"Required;MaxSize(50)" validAlias:"登录名"`                     // 登录名
	Password  string      `gorm:"column:password" json:"password" form:"password" query:"password" valid:"Required;MaxSize(50)" validAlias:"密码"`                      // 密码
	Salt      string      `gorm:"column:salt" json:"salt" form:"salt" query:"salt" valid:"Required;MaxSize(10)" validAlias:"盐"`                                       // 盐
	Name      string      `gorm:"column:name" json:"name" form:"name" query:"name" valid:"Required;MaxSize(100)" validAlias:"姓名"`                                     // 姓名
	Email     null.String `gorm:"column:email" json:"email" form:"email" query:"email" valid:"MaxSize(200)" validAlias:"邮箱"`                                          // 邮箱
	Phone     null.String `gorm:"column:phone" json:"phone" form:"phone" query:"phone" valid:"MaxSize(200)" validAlias:"电话"`                                          // 电话
	Mobile    null.String `gorm:"column:mobile" json:"mobile" form:"mobile" query:"mobile" valid:"MaxSize(200)" validAlias:"手机"`                                      // 手机
	APIID     string      `gorm:"column:api_id;default:'0'" json:"api_id" form:"api_id" query:"api_id" valid:"Required;MaxSize(100)" validAlias:"第三方API的用户ID"`        // 第三方API的用户ID
	UserType  string      `gorm:"column:user_type" json:"user_type" form:"user_type" query:"user_type" valid:"Required;MaxSize(30)" validAlias:"用户类型"`                // 用户类型，sec：教务，teacher：教师，student：学生
	Photo     string      `gorm:"column:photo" json:"photo" form:"photo" query:"photo" valid:"MaxSize(500)" validAlias:"用户头像"`                                        // 用户头像
	LoginIP   null.String `gorm:"column:login_ip" json:"login_ip" form:"login_ip" query:"login_ip" valid:"MaxSize(100)" validAlias:"最后登陆IP"`                          // 最后登陆IP
	LoginDate null.Time   `gorm:"column:login_date" json:"login_date" form:"login_date" query:"login_date" validAlias:"最后登陆时间"`                                       // 最后登陆时间
	LoginFlag string      `gorm:"column:login_flag;default:'0'" json:"login_flag" form:"login_flag" query:"login_flag" valid:"Required;MaxSize(1)" validAlias:"登录标记"` // 登录标记，0、不可登录，1、可登录
	CreatedBy string      `gorm:"column:created_by" json:"created_by" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`             // 创建者
	CreatedAt time.Time   `gorm:"column:created_at" json:"created_at" form:"created_at" query:"created_at" validAlias:"创建时间"`                                         // 创建时间
	UpdatedBy string      `gorm:"column:updated_by" json:"updated_by" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`             // 更新者
	UpdatedAt time.Time   `gorm:"column:updated_at" json:"updated_at" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                                         // 更新时间
	Remarks   null.String `gorm:"column:remarks" json:"remarks" form:"remarks" query:"remarks" valid:"MaxSize(255)" validAlias:"备注信息"`                                // 备注信息
	DeletedAt null.Time   `gorm:"column:deleted_at" json:"deleted_at" form:"deleted_at" query:"deleted_at" validAlias:"删除时间"`                                         // 删除时间
	SchoolID  string      `gorm:"column:school_id" json:"school_id" form:"school_id" query:"school_id" validAlias:"学校ID"`                                             // 学校ID
}

// TableName sets the insert table name for this struct type
func (s *SysUser) TableName() string {
	return "sys_user"
}
