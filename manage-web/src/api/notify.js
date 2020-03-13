import { axios } from '@/utils/request'

const api = {
  GetSysNotify: '/auth/sys/notify/get',
  GetSysNotifyInfo: '/auth/sys/notify/info',
  SaveSysNotify: '/auth/sys/notify/save',
  GetSysNotifyPage: '/auth/sys/notify/page',
  GetSysNotifyList: '/auth/sys/notify/list',
  ReleaseSysNotify: '/auth/sys/notify/release',
  DelSysNotify: '/auth/sys/notify/remove',
  DelSysNotifyFile: '/auth/sys/notify/file/remove'
}

export default api

export function GetSysNotify (parameter) {
  return axios({
    url: api.GetSysNotify,
    method: 'POST',
    data: parameter
  })
}

export function GetSysNotifyInfo (parameter) {
  return axios({
    url: api.GetSysNotifyInfo,
    method: 'POST',
    data: parameter
  })
}

export function SaveSysNotify (parameter) {
  return axios({
    url: api.SaveSysNotify,
    method: 'POST',
    data: parameter
  })
}

export function GetSysNotifyList (parameter) {
  return axios({
    url: api.GetSysNotifyList,
    method: 'POST',
    data: parameter
  })
}
export function GetSysNotifyPage (parameter) {
  return axios({
    url: api.GetSysNotifyPage,
    method: 'POST',
    data: parameter
  })
}

export function DelSysNotify (parameter) {
  return axios({
    url: api.DelSysNotify,
    method: 'POST',
    data: parameter
  })
}

export function DelSysNotifyFile (parameter) {
  return axios({
    url: api.DelSysNotifyFile,
    method: 'POST',
    data: parameter
  })
}

export function ReleaseSysNotify (parameter) {
  return axios({
    url: api.ReleaseSysNotify,
    method: 'POST',
    data: parameter
  })
}
