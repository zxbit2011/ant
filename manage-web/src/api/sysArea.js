import { axios } from '@/utils/request'

const api = {
  GetSysAreaSubList: '/auth/sys/area/sub/list'
}

export default api

export function GetSysAreaSubList (parameter) {
  return axios({
    url: api.GetSysAreaSubList,
    method: 'POST',
    data: parameter
  })
}
