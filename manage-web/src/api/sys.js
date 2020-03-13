import { axios } from '@/utils/request'

// 管理平台 - 添加/修改菜单
export function SaveMenu (data) {
  return axios({
    url: '/auth/sys/menu/save',
    method: 'post',
    data
  })
}

// 管理平台 - 删除菜单
export function DelMenu (data) {
  return axios({
    url: '/auth/sys/menu/del',
    method: 'post',
    data
  })
}
// 管理平台 - 查询菜单
export function GetMenu (data) {
  return axios({
    url: '/auth/sys/menu/get',
    method: 'post',
    data
  })
}

// 管理平台 - 获取菜单按钮
export function GetBtnMenuList (data) {
  return axios({
    url: '/auth/sys/menu/btn/all',
    method: 'post',
    data
  })
}

// 管理平台 - 保存菜单按钮
export function SaveMenuBtn (data) {
  return axios({
    url: '/auth/sys/menu/btn/save',
    method: 'post',
    data
  })
}

// 管理平台 - 删除菜单按钮
export function DelMenuBtn (data) {
  return axios({
    url: '/auth/sys/menu/btn/del',
    method: 'post',
    data
  })
}

// 管理平台 - 获取菜单按钮
export function GetMenuBtn (data) {
  return axios({
    url: '/auth/sys/menu/btn/get',
    method: 'post',
    data
  })
}

// 获取所有的菜单列表
export function GetMenuList (data) {
  return axios({
    url: '/auth/sys/menu/all',
    method: 'post',
    data
  })
}

/* -----------角色管理------------- */

// 获取角色列表
export function GetRoleList (data) {
  return axios({
    url: '/auth/sys/role/all',
    method: 'post',
    data
  })
}

// 添加角色
export function AddRole (data) {
  return axios({
    url: '/auth/sys/role/save',
    method: 'post',
    data
  })
}

// 删除角色
export function DelRole (data) {
  return axios({
    url: '/auth/sys/role/del',
    method: 'post',
    data
  })
}
// 获取角色
export function GetRoleInfo (data) {
  return axios({
    url: '/auth/sys/role/get',
    method: 'post',
    data
  })
}

// 根据角色获取菜单资源
export function GetCheckedMenuLimit (data) {
  return axios({
    url: '/auth/sys/role/menu/all',
    method: 'post',
    data
  })
}
// 根据角色获取菜单资源按钮
export function GetCheckedMenuBtnLimit (data) {
  return axios({
    url: '/auth/sys/role/menu/btn/all',
    method: 'post',
    data
  })
}

// 操作角色资源关系表
export function AddRoleResource (data) {
  return axios({
    url: '/platformManage/role/resource/add',
    method: 'post',
    data
  })
}

/* -----------------用户管理--------------------- */

// 添加用户
export function AddSystemManage (data) {
  return axios({
    url: '/auth/sys/user/save',
    method: 'post',
    data
  })
}
// 禁用用户
export function DisableSysUser (data) {
  return axios({
    url: '/auth/sys/user/disable',
    method: 'post',
    data
  })
}

// 获取用户列表
export function GetSysUserList (data) {
  return axios({
    url: '/auth/sys/user/page',
    method: 'post',
    data
  })
}

// 获取用户列表
export function DelSystemManage (data) {
  return axios({
    url: '/auth/sys/user/del',
    method: 'post',
    data
  })
}

// 重置密码
export function ResetPassword (data) {
  return axios({
    url: '/auth/sys/user/pwd/reset',
    method: 'post',
    data
  })
}

// 添加菜单权限
export function UpdateRoleMenuByIds (data) {
  return axios({
    url: '/auth/sys/role/menu/save',
    method: 'post',
    data
  })
}
// 添加菜单按钮权限
export function UpdateRoleMenuBtnById (data) {
  return axios({
    url: '/auth/sys/role/menu/btn/save',
    method: 'post',
    data
  })
}

/* ------------------部门管理----------------- */

export function SaveSysOffice (data) {
  return axios({
    url: '/auth/sys/office/save',
    method: 'post',
    data
  })
}

export function GetSysOfficeAll (data) {
  return axios({
    url: '/auth/sys/office/all',
    method: 'post',
    data
  })
}

export function DelSysOffice (data) {
  return axios({
    url: '/auth/sys/office/del',
    method: 'post',
    data
  })
}

/* ------------操作日志---------------- */

// 获取操作列表
export function GetManageActionLog (data) {
  return axios({
    url: '/auth/sys/log/page',
    method: 'post',
    data
  })
}

/* ------------登陆者权限---------------- */

// 登录的菜单资源
export function GetLoginMenu (data) {
  return axios({
    url: '/auth/login/menu',
    method: 'post',
    data
  })
}
