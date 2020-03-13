package router

import (
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/handle/sys"
)

func SysRouter(e *echo.Group) {
	//地区
	area := e.Group("/sys/area")
	{
		area.POST("/sub/list", sys.GetSysAreaList)
	}
	//日志
	log := e.Group("/sys/log")
	{
		log.POST("/page", sys.GetSysLogPage)
	}
	//系统用户
	user := e.Group("/sys/user")
	{
		user.POST("/get", sys.GetSysUser)
		user.POST("/all", sys.GetSysUserAll)
		user.POST("/page", sys.GetSysUserPage)
		user.POST("/save", sys.SaveSysUser)
		user.POST("/del", sys.DelSysUser)
		user.POST("/pwd/reset", sys.ResetPassword)
		user.POST("/disable", sys.DisableSysUser)
	}
	//系统用户
	userRole := e.Group("/sys/user/role")
	{
		userRole.POST("/get", sys.GetSysUserRole)
		userRole.POST("/all", sys.GetSysUserRoleAll)
		userRole.POST("/page", sys.GetSysUserRolePage)
		userRole.POST("/save", sys.SaveSysUserRole)
		userRole.POST("/del", sys.DelSysUserRole)
	}
	//菜单
	menu := e.Group("/sys/menu")
	{
		menu.POST("/get", sys.GetSysMenu)
		menu.POST("/all", sys.GetSysMenuAll)
		menu.POST("/page", sys.GetSysMenuPage)
		menu.POST("/save", sys.SaveSysMenu)
		menu.POST("/del", sys.DelSysMenu)
	}
	//菜单按钮
	menuBtn := e.Group("/sys/menu/btn")
	{
		menuBtn.POST("/get", sys.GetSysMenuBtn)
		menuBtn.POST("/all", sys.GetSysMenuBtnAll)
		menuBtn.POST("/page", sys.GetSysMenuBtnPage)
		menuBtn.POST("/save", sys.SaveSysMenuBtn)
		menuBtn.POST("/del", sys.DelSysMenuBtn)
	}
	//部门
	office := e.Group("/sys/office")
	{
		office.POST("/get", sys.GetSysOffice)
		office.POST("/all", sys.GetSysOfficeAll)
		office.POST("/page", sys.GetSysOfficePage)
		office.POST("/save", sys.SaveSysOffice)
		office.POST("/del", sys.DelSysOffice)
	}
	//角色
	role := e.Group("/sys/role")
	{
		role.POST("/get", sys.GetSysRole)
		role.POST("/all", sys.GetSysRoleAll)
		role.POST("/page", sys.GetSysRolePage)
		role.POST("/save", sys.SaveSysRole)
		role.POST("/del", sys.DelSysRole)
	}
	//角色菜单
	roleMenu := e.Group("/sys/role/menu")
	{
		roleMenu.POST("/get", sys.GetSysRoleMenu)
		roleMenu.POST("/all", sys.GetSysRoleMenuAll)
		roleMenu.POST("/page", sys.GetSysRoleMenuPage)
		roleMenu.POST("/save", sys.SaveSysRoleMenu)
		roleMenu.POST("/del", sys.DelSysRoleMenu)
	}
	//角色菜单按钮
	roleMenuBtn := e.Group("/sys/role/menu/btn")
	{
		roleMenuBtn.POST("/get", sys.GetSysRoleMenuBtn)
		roleMenuBtn.POST("/all", sys.GetSysRoleMenuBtnAll)
		roleMenuBtn.POST("/page", sys.GetSysRoleMenuBtnPage)
		roleMenuBtn.POST("/save", sys.SaveSysRoleMenuBtn)
		roleMenuBtn.POST("/del", sys.DelSysRoleMenuBtn)
	}
}
