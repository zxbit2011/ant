package utils

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type ResultParam struct {
	Ret  int64       `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	//成功
	SuccessCode = 200
	//错误
	ErrorCode = 400
	//授权失败
	AuthFailCode = 401
	//参数校验失败
	ParamsErrorCode = 402
	//权限不足
	PermissionDeniedCode = 403
	//数据为空
	NullDataCode = 404
	//图片验证码为空
	ImgCodeNull = 500
	//图片验证码错误
	ImgCodeFailCode = 501
	//询问
	ConfirmCode = 700
)

const (
	GetParsFailResult   = "获取参数失败"
	ParsFailResult      = "参数校验失败"
	GetDataNullResult   = "数据不存在"
	GetDataFailResult   = "获取数据失败"
	AddSuccessResult    = "添加成功"
	UpdateSuccessResult = "修改成功"
	DeleteSuccessResult = "删除成功"
	AddFailResult       = "添加失败"
	UpdateFailResult    = "修改失败"
	DeleteFailResult    = "删除失败"
	GetFailResult       = "获取失败"
)

func ToResultParam(b []byte) ResultParam {
	var rp ResultParam
	err := json.Unmarshal(b[:], &rp)
	if err != nil {
		println(err.Error())
		return ResultParam{}
	}
	return rp
}

func SuccessRespone(c echo.Context, data string) error {
	return c.HTML(http.StatusOK, data)
}

func Success(c echo.Context, msg string, data interface{}) error {
	return Result(c, SuccessCode, msg, data)
}

func SuccessNullMsg(c echo.Context, data interface{}) error {
	return Result(c, SuccessCode, "", data)
}
func SuccessNull(c echo.Context, msg string) error {
	return Result(c, SuccessCode, msg, nil)
}

func ImgCodeNUll(c echo.Context) error {
	return Result(c, ImgCodeNull, "请输入验证码", nil)
}

func ImgCodeFail(c echo.Context) error {
	return Result(c, ImgCodeFailCode, "验证码错误", nil)
}
func ConfirmNUll(c echo.Context, msg string) error {
	return Result(c, ConfirmCode, msg, nil)
}

func Confirm(c echo.Context, msg string, data interface{}) error {
	return Result(c, ConfirmCode, msg, data)
}

func Error(c echo.Context, msg string, data interface{}) error {
	return Result(c, ErrorCode, msg, data)
}

func ErrorNull(c echo.Context, msg string) error {
	return Result(c, ErrorCode, msg, nil)
}

func ParamsError(c echo.Context, data interface{}) error {
	return Result(c, ParamsErrorCode, ParsFailResult, data)
}

func NullData(c echo.Context) error {
	return Result(c, NullDataCode, "暂无数据", nil)
}

func AuthFail(c echo.Context, msg string) error {
	return Result(c, AuthFailCode, msg, nil)
}

func AuthFailNull(c echo.Context) error {
	return Result(c, AuthFailCode, "未登录或登陆已失效", nil)
}

func PermissionDenied(c echo.Context) error {
	return Result(c, PermissionDeniedCode, "权限不足，限制访问", nil)
}

func AuthFailOrgNull(c echo.Context) error {
	return Result(c, AuthFailCode, "无管理权限", nil)
}

func ResultApi(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func Result(c echo.Context, ret int64, msg string, data interface{}) error {
	resultMap := map[string]interface{}{
		"ret":  ret,
		"msg":  msg,
		"data": data,
	}
	return c.JSON(http.StatusOK, resultMap)
}

func Alert(c echo.Context, tip string) error {
	return RedirectAndAlert(c, tip, "")
}

func RedirectAndAlert(c echo.Context, tip, url string) error {
	var js string
	if tip != "" {
		js += fmt.Sprintf("alert('%v');", tip)
	}
	js += fmt.Sprintf("parent.location.href = '%v';", url)
	return ResultHtml(c, fmt.Sprintf("<script>%v</script>", js))
}

func Redirect(c echo.Context, url string) error {
	return c.Redirect(http.StatusMovedPermanently, url)
}

func ResultHtml(c echo.Context, html string) error {
	return c.HTML(http.StatusOK, html)
}

func ResultString(c echo.Context, str string) error {
	return c.String(http.StatusOK, str)
}
