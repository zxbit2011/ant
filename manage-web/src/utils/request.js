import Vue from 'vue'
import axios from 'axios'
import qs from 'qs'
import store from '@/store/'
import { VueAxios } from './axios'
import notification from 'ant-design-vue/es/notification'
import { ACCESS_TOKEN } from '@/store/mutation-types'

// 创建 axios 实例
const service = axios.create({
  baseURL: '/api', // api base_url
  timeout: 30000, // 请求超时时间
  responseType: 'json',
  transformRequest: [function (data) {
    return qs.stringify(data)
  }]
})

const err = (error) => {
  notification.error({
    duration: 2,
    message: 'HTTP ERROR',
    description: error.message
  })
  return { ret: 400, msg: error.message }
}

// request interceptor
service.interceptors.request.use(config => {
  const token = Vue.ls.get(ACCESS_TOKEN)
  if (token) {
    config.headers['Access-Token'] = token // 让每个请求携带自定义 token 请根据实际情况自行修改
  }
  return config
}, err)

// response interceptor
service.interceptors.response.use((response) => {
  // dataAxios 是 axios 返回数据中的 data
  const dataAxios = response.data
  // 这个状态码是和后端约定的
  const { ret, msg, data } = dataAxios
  if (ret === undefined) {
    return dataAxios
  } else {
    switch (ret) {
      case 200:
      case 700:
        return dataAxios
      case 401:
        notification.warning({
          duration: 2,
          message: '系统消息',
          description: msg,
          onClose: () => {
            store.commit('SET_TOKEN', '')
            store.commit('SET_ROLES', [])
            Vue.ls.remove(ACCESS_TOKEN)
            if (location.pathname !== '/user/login') {
              location.href = '/user/login'
            }
          }
        })
        break
      case 400: // 错误
      case 501: // 图片验证码错误
        if (msg === '参数验证失败') {
          const errs = new Array()
          for (let i = 0; i < data.length; i++) {
            errs.push(`${i + 1}、${data[i]}`)
          }
          notification.error({
            duration: 2,
            message: '系统消息',
            description: errs.join('')
          })
        } else {
          notification.error({
            duration: 2,
            message: '系统消息',
            description: msg
          })
        }
        break
      default:
        notification.warning({
          duration: 2,
          message: '系统消息',
          description: msg
        })
        break
    }
    return dataAxios
  }
}, err)

const installer = {
  vm: {},
  install (Vue) {
    Vue.use(VueAxios, service)
  }
}

export {
  installer as VueAxios,
  service as axios
}
