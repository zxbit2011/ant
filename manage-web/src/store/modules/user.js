import Vue from 'vue'
import { getInfo, login, logout, loginResult } from '@/api/login'
import { ACCESS_TOKEN } from '@/store/mutation-types'
import { welcome } from '@/utils/util'
import router from '../../router'
import store from '../../store'
import notification from 'ant-design-vue/es/notification'
import cookies from '@/utils/cookies'

const user = {
  state: {
    token: '',
    name: '',
    welcome: '',
    avatar: '',
    roles: [],
    info: {}
  },

  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_NAME: (state, { name, welcome }) => {
      state.name = name
      state.welcome = welcome
    },
    SET_AVATAR: (state, avatar) => {
      state.avatar = avatar
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    SET_INFO: (state, info) => {
      state.info = info
    }
  },

  actions: {
    // 登录
    Login ({ commit }, userInfo) {
      return new Promise((resolve, reject) => {
        login(userInfo).then(res => {
          if (res.ret === 200) {
            const result = res.data

            commit('SET_TOKEN', '')
            commit('SET_ROLES', [])
            Vue.ls.remove(ACCESS_TOKEN)

            Vue.ls.set(ACCESS_TOKEN, result.Token, result.TokenExpire * 1000)
            commit('SET_TOKEN', result.Token)

            cookies.set('token', result.Token, { expires: 0.083 })
          }
          resolve(res)
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 获取用户信息
    GetInfo ({ commit }) {
      return new Promise((resolve, reject) => {
        getInfo().then(response => {
          if (response.ret === 200) {
            const result = response.data
            const loginInfo = result.loginInfo

            // 权限
            if (loginInfo.SysMenu && loginInfo.SysMenu.length > 0) {
              loginInfo.permissions = loginInfo.SysMenuBtns.map(permission => {
                return permission.permission
              })
              loginInfo.permissionList = loginInfo.SysMenu.map(permission => {
                return permission.permission
              })
              commit('SET_ROLES', loginInfo)
              commit('SET_INFO', loginInfo)
            } else {
              // 未分配任何权限
              // reject(new Error('getInfo: roles must be a non-null array !'))
              if (!loginInfo.IsSuperAdmin) {
                notification.warning({
                  duration: 2,
                  message: '系统消息',
                  description: '无任何权限'
                })
                commit('SET_TOKEN', '')
                commit('SET_ROLES', [])
                Vue.ls.remove(ACCESS_TOKEN)
              } else {
                commit('SET_ROLES', loginInfo)
                commit('SET_INFO', loginInfo)
              }
            }
            // 资源
            commit('SET_NAME', { name: loginInfo.Name, welcome: welcome() })
            commit('SET_AVATAR', loginInfo.Photo)
          }
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    LoginCheck ({ commit }, query) {
      return new Promise((resolve, reject) => {
        loginResult(query).then(res => {
          if (res.ret === 200) {
            /* const result = res.data
            const loginInfo = result.loginInfo

            // 权限
            if (loginInfo.SysMenuBtns && loginInfo.SysMenuBtns.length > 0) {
              loginInfo.permission = loginInfo.SysMenuBtns.map(permission => {
                return permission.permission
              })
              loginInfo.permissionList = loginInfo.SysMenu.map(permission => {
                return permission.permission
              })

              commit('SET_ROLES', loginInfo)
              commit('SET_INFO', loginInfo)
            }
            // 资源
            commit('SET_NAME', { name: loginInfo.Name, welcome: welcome() })
            commit('SET_AVATAR', loginInfo.Photo)
            Vue.ls.set(ACCESS_TOKEN, result.token.Token, result.token.TokenExpire * 1000)
            commit('SET_TOKEN', result.token.Token) */

            const result = res.data
            Vue.ls.set(ACCESS_TOKEN, result.token.Token, result.token.TokenExpire * 1000)
            commit('SET_TOKEN', result.token.Token)
            resolve(res)
          }
          resolve(res)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 登出
    Logout ({ commit, state }) {
      return new Promise((resolve) => {
        logout(state.token).then(() => {
          commit('SET_TOKEN', '')
          commit('SET_ROLES', [])
          Vue.ls.remove(ACCESS_TOKEN)
          resolve()
        }).catch(() => {
          resolve()
        })
      })
    }

  }
}

export default user
