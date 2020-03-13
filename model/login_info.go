package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/zxbit2011/ant/utils/enum"
	"time"
)

//后台系统登陆
type SysUserLoginInfo struct {
	ID                string         //ID
	APIID             string         //第三方ID
	Name              string         //账号名称 - 通过第三方接口查询
	Sex               string         //性别
	SchoolId          string         //学校ID
	SchoolName        string         //学校
	ClassNo           string         //学生 - 编辑编号
	ClassName         string         //学生 - 班级名称
	CompanyId         string         //公司ID
	CompanyName       string         //公司名称
	DataSource        string         //数据来源，本地数据库：local，大渡口数据中心：ddk
	OfficeId          string         //组织ID
	OfficeParentId    string         //组织关系
	OfficeRelationIds string         //组织关系
	UserType          string         //账号类型
	Photo             string         //头像
	SysRoles          []VSysUserRole //角色，支持多角色
	//SysMenuThree      []SysMenuTree        //菜单树
	SysMenu      []VSysRoleMenu       //菜单
	SysMenuBtns  []VSysRoleMenuBtn    //菜单按钮权限
	IsSuperAdmin bool                 //是否为超级管理员
	IsAdmin      bool                 //是否管理员
	PathAuth     map[string]time.Time //url已通过的鉴权
}

// login token struct
type SysUserLoginToken struct {
	ID   string        //ID
	Unix int64         //时间戳
	jwt.StandardClaims //jwt
}

// token struct
type SysToken struct {
	Token       string //token
	TokenExpire int64  //token过期时间戳
}

//菜单树
type SysMenuTree struct {
	VSysRoleMenu
	Children []SysMenuTree
}

//登录token struct
func (info SysUserLoginInfo) GetLoginToken() *SysUserLoginToken {
	return &SysUserLoginToken{
		ID:   info.ID,
		Unix: time.Now().Unix(),
	}
}

//是否为超级管理员
func IsSuperAdmin(roles []VSysUserRole) bool {
	for i := 0; i < len(roles); i++ {
		if roles[i].RoleType == enum.UserPostSuperAdminCode {
			return true
		}
	}
	return false
}

//是否为超级管理员
func IsAdmin(roles []VSysUserRole) bool {
	for i := 0; i < len(roles); i++ {
		if roles[i].RoleType == enum.UserPostAdminCode {
			return true
		}
	}
	return false
}

//登录信息
func GetLoginInfo(vSysUser VSysUser, roles []VSysUserRole, menus []VSysRoleMenu, menuBtns []VSysRoleMenuBtn) SysUserLoginInfo {
	return SysUserLoginInfo{
		ID:                vSysUser.ID,
		APIID:             vSysUser.APIID,
		Name:              vSysUser.Name,
		UserType:          vSysUser.UserType,
		Photo:             vSysUser.Photo,
		CompanyId:         vSysUser.SysCompanyId,
		CompanyName:       vSysUser.SysCompanyName,
		DataSource:        vSysUser.DataSource,
		OfficeId:          vSysUser.OfficeID,
		OfficeParentId:    vSysUser.OfficeParentId,
		OfficeRelationIds: vSysUser.RelationIds,
		SchoolId:          vSysUser.SchoolID,
		SysRoles:          roles,
		//SysMenuThree:      GetSysMenuThree(menus, "0"),
		SysMenu:      menus,
		SysMenuBtns:  menuBtns,
		IsAdmin:      IsAdmin(roles),
		IsSuperAdmin: IsSuperAdmin(roles),
		PathAuth:     make(map[string]time.Time, 0),
	}
}

//获取树形菜单
func GetSysMenuThree(menus []VSysRoleMenu, parentId string) []SysMenuTree {
	var smuts = make([]SysMenuTree, 0)
	var smut SysMenuTree
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentID != parentId {
			continue
		}
		smut = SysMenuTree{}
		smut.ID = menus[i].ID
		smut.ParentID = menus[i].ParentID
		smut.Name = menus[i].Name
		smut.Href = menus[i].Href
		smut.Icon = menus[i].Icon
		smut.IsShow = menus[i].IsShow
		smut.Permission = menus[i].Permission
		smut.RelationIds = menus[i].RelationIds
		smut.Remarks = menus[i].Remarks
		smut.Target = menus[i].Target
		smut.Children = GetSysMenuThree(menus, menus[i].ID)
		smuts = append(smuts, smut)
	}
	return smuts
}

//获取树形菜单
func GetSysMenuTree(menus []SysMenu, parentId string) []SysMenuTree {
	var smuts = make([]SysMenuTree, 0)
	var smut SysMenuTree
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentID != parentId {
			continue
		}
		smut = SysMenuTree{}
		smut.ID = menus[i].ID
		smut.ParentID = menus[i].ParentID
		smut.Name = menus[i].Name
		smut.Href = menus[i].Href
		smut.Icon = menus[i].Icon
		smut.Sort = menus[i].Sort
		smut.RelationIds = menus[i].RelationIds
		smut.Remarks = menus[i].Remarks
		smut.Target = menus[i].Target
		smut.Children = GetSysMenuTree(menus, menus[i].ID)
		smuts = append(smuts, smut)
	}
	return smuts
}

//部门树
type SysOfficeTree struct {
	SysOffice
	Children []SysOfficeTree `json:"children,omitempty"`
}

type Tree struct {
	ParentID    string      `json:"parentId"`
	Title       string      `json:"title"`
	Key         string      `json:"key"`
	Icon        string      `json:"icon",omitempty`
	Children    []Tree      `json:"children,omitempty"`
	ScopedSlots ScopedSlots `json:"scopedSlots,omitempty"`
}

type ScopedSlots struct {
	Icon     string `json:"icon"`
	ShowIcon string `json:"showIcon"`
}

//获取树形菜单
func GetSysMenuTreeUI(menus []SysMenu, parentId string) []Tree {
	var tree = make([]Tree, 0)
	var t Tree
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentID != parentId {
			continue
		}
		t = Tree{}
		t.Key = menus[i].ID
		t.ParentID = menus[i].ParentID
		t.Title = menus[i].Name
		t.Icon = menus[i].Icon.String
		t.ScopedSlots = ScopedSlots{Icon: "custom", ShowIcon: t.Icon}
		t.Children = GetSysMenuTreeUI(menus, menus[i].ID)
		tree = append(tree, t)
	}
	return tree
}

//获取树形菜单
func GetSysOfficeTreeUI(menus []SysOffice, parentId string) []Tree {
	var tree = make([]Tree, 0)
	var t Tree
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentID != parentId {
			continue
		}
		t = Tree{}
		t.Key = menus[i].ID
		t.ParentID = menus[i].ParentID
		t.Title = menus[i].Name
		t.Children = GetSysOfficeTreeUI(menus, menus[i].ID)
		tree = append(tree, t)
	}
	return tree
}
