import { axios } from '@/utils/request'

const api = {
  GetGroup: '/auth/index/group'
}

export function GetGroup (parameter) {
  return axios({
    url: api.GetGroup,
    method: 'post',
    params: parameter
  })
}
