package handle

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/model"
	"strings"
)

// 根据创建人进行权限判断 - 数据操作权限判断【可以加缓存】
// userInfo 登录人
// createBy 数据的创建人
func PowerCheck(userInfo model.SysUserLoginInfo, createBy string) (err error) {
	if !userInfo.IsSuperAdmin {
		//非超级管理员权限
		if userInfo.IsAdmin {
			//管理员
			var vSysUser model.VSysUser
			if err = global.DB.First(&vSysUser, "id=?", createBy).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					err = errors.New("创建者不存在")
					return
				}
				global.Log.Error("PowerCheck error：", err)
				err = errors.New("获取创建者失败")
				return
			}
			if !strings.Contains(vSysUser.RelationIds, userInfo.OfficeRelationIds) {
				err = errors.New("无权限操作")
				return
			}
		} else {
			//普通账号
			if userInfo.ID != createBy {
				err = errors.New("无权限操作")
				return
			}
		}
	}
	return
}
