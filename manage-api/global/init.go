package global

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils/config"
	"github.com/zxbit2011/gifCaptcha"
	"image/color"
	"time"
)

var (
	//DB 数据库操作
	DB *gorm.DB
	//Conf 配置相关
	Conf *config.Config
	//Log 日志
	Log *logs.BeeLogger
	//RD redis
	RD *config.RedisConnPool
	//Session
	Session = func(c echo.Context) *config.USession { return config.GetSession(c) }
	//Secret
	Secret           = []byte("5eF6Xj8z#pZxBOkavlcPq^MmC09*S*!8!8Jor8V7*m0F3*zWReV%o3taoH%DI@ni")
	VerificationCode = "VerificationCode"

	GifCaptcha = gifCaptcha.New()
)

const (
	//授权过期时间 24 小时
	AuthExpireTime     int64 = 2 * 60 * 60
	PathAuthExpireTime int64 = 10 * 60
	//登录地址
	RedirectUrl = "/login"
	//登录失效提示
	AuthInvalidMsg = "未登录或登陆已失效"
	//登录失败提示
	AuthErrorMsg          = "登陆失败，请重新登陆"
	AuthLoginInfoErrorMsg = "获取登录信息失败"
	//登录cookie名称
	TokenName     = "zm_token"
	LoginInfoName = "loginInfo"
)

//ConfPath 配置文件路径
type ConfPath struct {
	ConfigPath string
}

//InitGlobal 初始化各项配置
func InitGlobal(confPath *ConfPath) (err error) {
	// 初始化配置文件
	Conf, err = config.LoadGlobalConfig(confPath.ConfigPath)
	if err != nil {
		err = errors.New(fmt.Sprintf("初始化toml配置文件失败，error：%s", err.Error()))
		return
	}
	//初始化数据库连接
	DB, err = config.NewGorm(Conf)
	if err != nil {
		err = errors.New(fmt.Sprintf("初始化数据库失败，error：%s", err.Error()))
		return
	}
	//初始化redis
	RD, err = config.InitRedis(Conf.Redis, Conf.Project.Name)
	if err != nil {
		err = errors.New(fmt.Sprintf("初始化Redis数据库失败，error：%s", err.Error()))
		return
	}
	Log = config.InitLog(Conf.Log)
	GifCaptcha.SetFrontColor(color.Black, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	return
}

//登录信息
func GetLoginInfo(c echo.Context) model.SysUserLoginInfo {
	return c.Get(LoginInfoName).(model.SysUserLoginInfo)
}

/**
单点登录的唯一标记格式
*/
func GetSysUserLoginFlag(id string) string {
	return fmt.Sprintf("sys_user_login_token_%s", id)
}

var TokenCookieFlag = map[string]bool{
	"/files/*/auth/*": true,
}

/**
检查token有效性
cookieFlag cookie方式校验 一般用于Get请求的时候
*/
func GetToken(c echo.Context) (tokenStr string, err error) {
	//auth
	tokenStr = c.Request().Header.Get("Access-Token")
	path := c.Path()
	if TokenCookieFlag[path] {
		if tokenStr == "" {
			ck, err := c.Cookie(TokenName)
			if err == nil && ck.Value != "" {
				tokenStr = ck.Value
			}
		}
	}
	if tokenStr == "" {
		tokenStr = c.FormValue(TokenName)
	}
	if tokenStr == "" {
		err = errors.New(AuthInvalidMsg)
		return
	}
	return
}

/**
根据token获取Redis中的数据
tokenString：token参数
*/
func GetSysUserLoginInfo(tokenStr string) (loginInfo model.SysUserLoginInfo, err error) {
	tip := AuthInvalidMsg
	loginInfoStr, err := RD.GetString(tokenStr)
	if err != nil {
		Log.Error("get token redis error： %s", err.Error())
		err = errors.New(AuthInvalidMsg)
	}
	if loginInfoStr == "" {
		err = errors.New(tip)
		return
	}
	err = json.Unmarshal([]byte(loginInfoStr), &loginInfo)
	return
}

func QueryRows(db *gorm.DB, sqlStr string, val ...interface{}) (list []map[string]interface{}, err error) {
	var rows *sql.Rows
	rows, err = db.Raw(sqlStr, val...).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	var columns []string
	columns, err = rows.Columns()
	if err != nil {
		return
	}
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	// 这里需要初始化为空数组，否则在查询结果为空的时候，返回的会是一个未初始化的指针
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}

		ret := make(map[string]interface{})
		for i, col := range values {
			if col == nil {
				ret[columns[i]] = nil
			} else {
				switch val := (*scanArgs[i].(*interface{})).(type) {
				case []byte:
					ret[columns[i]] = string(val)
					break
				case time.Time:
					ret[columns[i]] = val.Format("2006-01-02 15:04:05")
					break
				default:
					ret[columns[i]] = val
				}
			}
		}
		list = append(list, ret)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
