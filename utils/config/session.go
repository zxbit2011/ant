package config

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func GetSession(c echo.Context) *USession {
	ses, _ := session.Get("Go_Session_Id", c)
	ses.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 20,
		HttpOnly: true,
	}
	return &USession{
		ses,
		c,
	}
}

type USession struct {
	*sessions.Session
	echo.Context
}

func (us *USession) Saves() {
	us.Save(us.Context.Request(), us.Context.Response())
}

func (us *USession) AddValue(key, value string) *USession {
	us.Values[key] = value
	us.Saves()
	return us
}

func (us *USession) DelKey(key string) *USession {
	delete(us.Values, key)
	us.Saves()
	return us
}

func (us *USession) GetValue(key string) string {
	return fmt.Sprintf("%v", us.Values[key])
}
