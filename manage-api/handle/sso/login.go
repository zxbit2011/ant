package sso

import (
	"errors"
	"fmt"
	"image/gif"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils"
	"github.com/zxbit2011/ant/utils/convert"
	"github.com/zxbit2011/ant/utils/encrypt"
	"github.com/zxbit2011/ant/utils/enum"
)

func ImageCode(c echo.Context) error {
	c.Set("Content-Type", "image/gif")
	gifData, code := global.GifCaptcha.RangCaptcha()
	global.Session(c).AddValue(global.VerificationCode, strings.ToLower(code)).Saves()
	err := gif.EncodeAll(c.Response().Writer, gifData)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	c.Response().Flush()
	return nil
}

//账号密码登录
func Login(c echo.Context) error {
	//账号验证
	userName := c.FormValue("username")
	password := c.FormValue("password")
	if userName == "" {
		return utils.ErrorNull(c, "账号不能为空")
	}
	if password == "" {
		return utils.ErrorNull(c, "密码不能为空")
	}
	if len(userName) > 50 {
		return utils.ErrorNull(c, "登录名格式错误")
	}
	if len(password) > 300 {
		return utils.ErrorNull(c, "密码格式错误")
	}
	imgCode := c.FormValue("imgCode")
	if imgCode == "" {
		return utils.ErrorNull(c, "请输入验证码")
	}
	//判断验证码是否正确
	if strings.ToLower(imgCode) != global.Session(c).GetValue(global.VerificationCode) {
		return utils.ErrorNull(c, "验证码错误")
	}
	var vSysUser model.VSysUser
	if err := global.DB.First(&vSysUser, "username=?", userName).Error; err != nil && err != gorm.ErrRecordNotFound {
		global.Log.Error("query sys_user error： %s", err.Error())
		return utils.ErrorNull(c, "登陆失败")
	}
	if encrypt.Sha1Encode(password+vSysUser.Salt) != vSysUser.Password {
		return utils.ErrorNull(c, "帐号或密码错误")
	}
	_, sysToken, err := LoginLib(vSysUser, c.RealIP())
	if err != nil {
		return utils.ErrorNull(c, err.Error())
	}
	//cookie设置
	cookie := new(http.Cookie)
	cookie.Name = global.TokenName
	cookie.Value = sysToken.Token
	cookie.Expires = time.Now().Add(time.Duration(global.AuthExpireTime) * time.Second)
	cookie.Path = "/"
	cookie.Domain = strings.Split(c.Request().Host, ":")[0]
	c.SetCookie(cookie)
	return utils.Success(c, "操作成功", sysToken)
}

//公共登录方法
func LoginLib(vSysUser model.VSysUser, ip string) (loginInfo model.SysUserLoginInfo, sysToken model.SysToken, err error) {
	if vSysUser.ID == "" {
		err = errors.New("账号不存在")
		return
	}
	if vSysUser.LoginFlag == enum.LoginFlagNo {
		err = errors.New("账号以被禁用")
		return
	}
	//获取用户角色
	var vSysUserRole []model.VSysUserRole
	if err = global.DB.Where("sys_user_id=?", vSysUser.ID).Find(&vSysUserRole).Error; err != nil {
		global.Log.Error("query v_sys_user_role error： %s", err.Error())
		err = errors.New("登录失败")
		return
	}
	var roleIds []string
	if len(vSysUserRole) > 0 {
		for i := 0; i < len(vSysUserRole); i++ {
			roleIds = append(roleIds, vSysUserRole[i].ID)
		}
	}
	//获取用户角色菜单
	var vSysRoleMenu []model.VSysRoleMenu
	if len(roleIds) > 0 {
		if err = global.DB.Where("sys_role_id in (?)", roleIds).Group("id").Find(&vSysRoleMenu).Error; err != nil {
			global.Log.Error("query v_sys_role_menu error： %s", err.Error())
			err = errors.New("登录失败")
			return
		}
	}
	//获取用户角色菜单按钮
	var vSysRoleMenuBtn []model.VSysRoleMenuBtn
	if len(roleIds) > 0 {
		if err = global.DB.Where("sys_role_id in (?)", roleIds).Group("id").Find(&vSysRoleMenuBtn).Error; err != nil {
			global.Log.Error("query v_sys_role_menu_btn error： %s", err.Error())
			err = errors.New("登录失败")
			return
		}
	}

	//登录信息
	loginInfo = model.GetLoginInfo(vSysUser, vSysUserRole, vSysRoleMenu, vSysRoleMenuBtn)

	// Generate encoded token and send it as response.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, loginInfo.GetLoginToken())
	tokenStr, err := token.SignedString(global.Secret)
	if err != nil {
		global.Log.Error("jwt token SignedString %s", err.Error())
		err = errors.New("登录失败")
		return
	}
	now := time.Now()
	// 更新登录时间登录IP
	if err = global.DB.Model(&model.SysUser{ID: loginInfo.ID}).Updates(map[string]interface{}{"login_ip": ip, "login_date": now}).Error; err != nil {
		err = errors.New("更新登录信息失败")
		return
	}

	sysUserFlag := global.GetSysUserLoginFlag(loginInfo.ID)
	//清除该帐号下的其他Token
	oldToken, err := global.RD.GetString(sysUserFlag)
	if err == nil {
		_, _ = global.RD.DelKey(oldToken)
	}

	//Redis 2小时
	//登录信息缓存
	_, err = global.RD.SetAndExpire(tokenStr, convert.MustJsonString(loginInfo), global.AuthExpireTime)
	if err != nil {
		global.Log.Error("redis SetAndExpire token error： %s", err.Error())
		err = errors.New("登录失败")
		return
	}
	//单点登录统一标记
	_, err = global.RD.SetAndExpire(sysUserFlag, token, global.AuthExpireTime)
	if err != nil {
		global.Log.Error("redis SetAndExpire token error： %s", err.Error())
		err = errors.New("登录失败")
		return
	}
	sysToken.TokenExpire = global.AuthExpireTime
	sysToken.Token = tokenStr
	return
}

/**
退出
*/
func LogOut(c echo.Context) error {
	token, err := global.GetToken(c)
	if err != nil {
		return utils.SuccessNull(c, "已安全退出")
	}
	loginInfo, err := global.GetSysUserLoginInfo(token)
	if err != nil {
		return utils.SuccessNull(c, "已安全退出")
	}
	sysUserFlag := global.GetSysUserLoginFlag(loginInfo.ID)
	_, err = global.RD.DelKey(sysUserFlag)
	if err != nil {
		global.Log.Error("redis DelKey sysUserFlag error： %s", err.Error())
	}
	_, err = global.RD.DelKey(token)
	if err != nil {
		global.Log.Error("redis DelKey token error： %s", err.Error())
	}
	cookie, _ := c.Cookie(global.TokenName)
	if cookie != nil {
		cookie.Expires = time.Now().Add(-10 * time.Minute)
	}
	return utils.SuccessNull(c, "已安全退出")
}

/**
获取登录的菜单资源权限
*/
func GetLoginMenu(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	return utils.SuccessNullMsg(c, map[string]interface{}{
		"loginInfo": loginInfo,
	})
}
