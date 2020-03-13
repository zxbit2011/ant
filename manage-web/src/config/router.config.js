// eslint-disable-next-line
import {BasicLayout, PageView, RouteView, UserLayout} from '@/layouts'

export const asyncRouterMap = [
  {
    path: '/',
    name: 'home',
    component: BasicLayout,
    meta: { title: '首页' },
    redirect: '/index',
    children: [
      {
        path: '/index',
        name: 'index',
        component: () => import('@/views/Index'),
        meta: { title: '工作台', keepAlive: true, icon: 'home', permission: ['index'] }
      },
      // message 信息管理
      {
        path: '/message',
        name: 'message',
        component: () => import('@/views/message/List'),
        meta: { title: '信息管理', icon: 'message', permission: ['message'] }
      },
      // sys user 系统设置 - 个人中心
      {
        path: '/personal',
        name: 'personal',
        meta: { title: '个人资料', icon: 'user', hideHeader: true },
        component: () => import('@/views/personal/Index'),
        hideChildrenInMenu: true,
        redirect: '/personal/base',
        children: [
          {
            path: '/personal/base',
            name: 'PersonalBase',
            component: () => import('@/views/personal/BaseSetting'),
            meta: { title: '个人资料', hidden: true, keepAlive: true  }
          },
          {
            path: '/personal/pwd',
            name: 'PersonalPwd',
            component: () => import('@/views/personal/Pwd'),
            meta: { title: '修改密码', hidden: true, keepAlive: true }
          },
          {
            path: '/personal/custom',
            name: 'PersonalCustom',
            component: () => import('@/views/personal/Custom'),
            meta: { title: '个性化设置', hidden: true, keepAlive: true }
          },
          {
            path: '/personal/notification',
            name: 'PersonalNotification',
            component: () => import('@/views/personal/Notification'),
            meta: { title: '新消息通知', hidden: true, keepAlive: true }
          }
        ]
      },
      {
        path: '/sys',
        name: 'sys',
        component: PageView,
        meta: { title: '系统设置', icon: 'setting', hideHeader: true, permission: ['setting'] },
        redirect: '/sys/menu/list',
        children: [
          {
            path: '/sys/menu/list',
            name: 'SysMenuList',
            component: () => import('@/views/sys/menu/List'),
            meta: { title: '菜单管理', hidden: true, permission: ['sysMenu'] }
          },
          {
            path: '/sys/role/list',
            name: 'SysRoleList',
            component: () => import('@/views/sys/role/List'),
            meta: { title: '角色列表', hidden: true, permission: ['sysRole'] }
          },
          {
            path: '/sys/user/list',
            name: 'SysUserList',
            component: () => import('@/views/sys/user/List'),
            meta: { title: '系统用户', hidden: true, permission: ['sysUser'] }
          },
          /* {
            path: '/sys/permission/List',
            name: 'SysPermissionList',
            component: () => import('@/views/sys/permission/List'),
            meta: { title: '权限列表', hidden: true, keepAlive: true, permission: ['sysPermission'] }
          },
          {
            path: '/sys/dic/List',
            name: 'SysDicList',
            component: () => import('@/views/sys/dic/List'),
            meta: { title: '数据字典', hidden: true, keepAlive: true, permission: ['sysDic'] }
          }, */
          {
            path: '/sys/log/list',
            name: 'SysLogList',
            component: () => import('@/views/sys/log/List'),
            meta: { title: '日志信息', hidden: true, permission: ['sysLog'] }
          }
        ]
      }
    ]
  },
  {
    path: '*', redirect: '/404', hidden: true
  }
]

/**
 * 基础路由
 * @type { *[] }
 */
export const constantRouterMap = [
  {
    path: '/user',
    component: UserLayout,
    redirect: '/user/login',
    hidden: true,
    children: [
      {
        path: 'login',
        name: 'login',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Login')
      },
      {
        path: 'register',
        name: 'register',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Register')
      },
      {
        path: 'result',
        name: 'result',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Result')
      },
      {
        path: 'register-result',
        name: 'registerResult',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/RegisterResult')
      }
    ]
  },
  {
    path: '/401',
    component: () => import('@/views/exception/401'),
    meta: { title: '401' }
  },
  {
    path: '/403',
    component: () => import('@/views/exception/403'),
    meta: { title: '403' }
  },
  {
    path: '/404',
    component: () => import('@/views/exception/404'),
    meta: { title: '404' }
  }
]
