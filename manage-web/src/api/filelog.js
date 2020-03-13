import { axios } from '@/utils/request'

const api = {
  RemoveFilesLog: '/auth/files/remove'
}

export default api

export function removeFilesLog (parameter) {
  return axios({
    url: api.RemoveFilesLog,
    method: 'POST',
    data: parameter
  })
}
