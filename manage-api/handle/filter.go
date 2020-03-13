package handle

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils"
	"github.com/zxbit2011/ant/utils/convert"
	"github.com/zxbit2011/ant/utils/enum"
	"github.com/zxbit2011/ant/utils/validation"
	"strings"
	"time"
)

func Filter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := global.GetToken(c)
		if err != nil {
			return utils.AuthFail(c, err.Error())
		}
		loginInfo, err := global.GetSysUserLoginInfo(token)
		if err != nil {
			return utils.AuthFail(c, err.Error())
		}
		var sysUser model.SysUser
		query := "id,company_id,office_id,username,name,user_type,photo,login_ip,login_date,login_flag,created_at,updated_at"
		if err := global.DB.Select(query).First(&sysUser, "id=?", loginInfo.ID).Error; err != nil {
			global.Log.Error("query sys_user error： %s", err.Error())
			return utils.AuthFail(c, global.AuthLoginInfoErrorMsg)
		}
		if sysUser.LoginFlag == enum.LoginFlagNo {
			return utils.AuthFail(c, "账号以被禁用")
		}
		//更新登录缓存时间
		sysUserFlag := global.GetSysUserLoginFlag(loginInfo.ID)
		_, err = global.RD.SetAndExpire(token, convert.MustJson(loginInfo), global.AuthExpireTime)
		if err != nil {
			global.Log.Error("redis SetAndExpire token error： %s", err.Error())
		}
		//单点登录统一标记
		_, err = global.RD.SetAndExpire(sysUserFlag, token, global.AuthExpireTime)
		if err != nil {
			global.Log.Error("redis SetAndExpire token error： %s", err.Error())
		}

		method := c.Request().Method
		path := c.Request().URL.Path
		//鉴定权限
		if !loginInfo.IsSuperAdmin {
			if v, ok := loginInfo.PathAuth[path+method]; ok && v.After(time.Now()) {
				//未过期
				println("鉴权未过期")
			} else {
				//获取所有鉴权列表，判断当前请求是否需要权限判定
				resourceAll, err := GetResourceAll()
				if err != nil {
					return utils.ErrorNull(c, "获取资源权限失败")
				}
				//非超级管理员，根据资源进行鉴权
				if checkResourceAllPath(resourceAll, method, path) && !checkResourcePath(loginInfo.SysMenuBtns, method, path) {
					return utils.PermissionDenied(c)
				}
				//权限缓存，10分钟
				loginInfo.PathAuth[path+method] = time.Now().Add(10 * time.Minute)
			}
		}
		c.Set(global.LoginInfoName, loginInfo)
		SaveSysLog(c)
		return next(c)
	}
}

func GetResourceAll() (btns []model.SysMenuBtn, err error) {
	cacheResourceAllFlag := "cacheResourceAllFlag"
	cache, _ := global.RD.GetString(cacheResourceAllFlag)
	if cache == "" {
		if err = global.DB.Find(&btns).Error; err != nil {
			return
		}
		//redis缓存10分钟
		_, err = global.RD.SetAndExpire(cacheResourceAllFlag, convert.MustJsonString(btns), global.PathAuthExpireTime)
	} else {
		err = json.Unmarshal([]byte(cache), &btns)
	}
	return
}

//判断请求的Url是否在鉴权范围内
func checkResourceAllPath(btns []model.SysMenuBtn, method string, path string) (flag bool) {
	for _, v := range btns {
		if strings.ToLower(v.Path) == strings.ToLower(path) && strings.ToLower(v.Method) == strings.ToLower(method) {
			//在鉴权范围内
			flag = true
		}
	}
	return
}

//判断用户的所有权限是否包含当前的请求Url
func checkResourcePath(btns []model.VSysRoleMenuBtn, method string, path string) (flag bool) {
	for _, v := range btns {
		if strings.ToLower(v.Path) == strings.ToLower(path) && strings.ToLower(v.Method) == strings.ToLower(method) {
			//在鉴权范围内
			flag = true
		}
	}
	return
}

var LogMap = map[string]string{
	"/api/auth/project/save":                   "项目保存",
	"/api/auth/sys/role/menu/btn/del":          "删除操作权限",
	"/api/auth/sys/role/del":                   "删除系统角色",
	"/api/auth/project/task/quota/save":        "项目任务指标保存",
	"/api/webapi/task/AddTaskVersion":          "添加任务版本",
	"/api/auth/logout":                         "系统退出",
	"/api/auth/project/task/quota/set/del":     "删除项目任务指标评价方式",
	"/api/auth/sys/office/save":                "保存组织架构",
	"/api/auth/project/task/quota/set/save":    "保存项目指标评价方式",
	"/api/auth/sys/menu/save":                  "保存系统菜单",
	"/api/auth/project/type/remove":            "删除项目类型",
	"/api/auth/sys/notify/file/remove":         "删除系统通知信息文件",
	"/api/auth/sys/office/del":                 "删除组织架构",
	"/api/webapi/task/DelTaskDataRecord":       "删除任务数据记录",
	"/api/auth/sys/user/role/del":              "删除系统用户角色",
	"/api/auth/project/task/file/remove":       "删除项目任务文件",
	"/api/auth/sys/notify/remove":              "删除系统通知信息",
	"/api/auth/project/quota/remove":           "删除项目指标",
	"/api/auth/sys/role/save":                  "保存系统角色",
	"/api/auth/project/quota/save":             "保存项目指标",
	"/api/auth/sys/user/pwd/reset":             "重置系统用户密码",
	"/api/webapi/task/over":                    "完成任务",
	"/api/auth/sys/notify/save":                "保存系统通知信息",
	"/api/auth/project/remove":                 "删除项目",
	"/api/auth/project/task/quota/del":         "删除项目任务指标",
	"/api/auth/sys/menu/btn/del":               "删除菜单操作权限",
	"/api/auth/project/type/save":              "保存项目类型",
	"/api/auth/sys/user/del":                   "删除系统用户",
	"/api/auth/project/task/save":              "保存项目任务",
	"/api/auth/project/review":                 "审核项目",
	"/api/webapi/task/RemoveTaskQuotaDataById": "删除任务指标数据",
	"/api/auth/sys/role/menu/del":              "删除角色菜单",
	"/api/auth/project/task/account/del":       "删除项目任务参与人员",
	"/api/auth/sys/user/save":                  "保存系统用户",
	"/api/auth/sys/notify/release":             "发布系统通知信息",
	"/api/auth/sys/role/menu/btn/save":         "保存菜单操作权限",
	"/api/auth/project/quota/type/remove":      "删除项目指标类型",
	"/api/auth/sys/role/menu/save":             "保存角色菜单",
	"/api/auth/project/task/account/save":      "保存项目任务参与人员",
	"/api/auth/project/task/remove":            "删除项目任务",
	"/api/auth/project/file/remove":            "删除项目文件",
	"/api/auth/project/quota/type/save":        "保存项目指标类型",
	"/api/auth/sys/user/role/save":             "保存系统用户角色",
	"/api/auth/sys/menu/del":                   "删除系统菜单",
	"/api/auth/sys/menu/btn/save":              "保存系统菜单操作权限",
	"/api/webapi/quota/Save":                   "保存指标",
}

// 日志表保存数据
func SaveSysLog(c echo.Context) {
	path := c.Request().URL.Path
	title, ok := LogMap[path]
	if ok {
		var t = "1" //后端
		if strings.Contains(path, "/api/webapi") {
			t = "2" //前端
		}
		loginInfo := global.GetLoginInfo(c)
		var sysLog model.SysLog
		//新增 false
		sysLog.ID = utils.ID()
		sysLog.Title = title
		sysLog.ID = utils.ID()
		sysLog.Type = t
		sysLog.CreatedBy = loginInfo.ID
		sysLog.RemoteAddr = c.RealIP()
		sysLog.UserAgent = c.Request().UserAgent()
		sysLog.RequestURI = c.Path()
		sysLog.Method = c.Request().Method
		uv, _ := c.FormParams()
		var arr []string
		for k, v := range uv {
			arr = append(arr, fmt.Sprintf("%s=%v", k, strings.Join(v, "")))
		}
		sysLog.Params = strings.Join(arr, "&")
		errs := validation.Valid(&sysLog)
		if len(errs) > 0 {
			global.Log.Error("参数验证失败，%v", errs)
			return
		}
		if err := global.DB.Create(&sysLog).Error; err != nil {
			global.Log.Error(err.Error())
		}

	}
}
