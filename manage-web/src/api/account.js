import { axios } from '@/utils/request'

const api = {
  // project
  AccountSecPage: '/auth/account/sec/page',
  AccountTeacherPage: '/auth/account/teacher/page',
  AccountStudentPage: '/auth/account/student/page',
  AccountSchoolPage: '/auth/account/school/page'
}

export default api

export function getAccountSecPage (parameter) {
  return axios({
    url: api.AccountSecPage,
    method: 'POST',
    data: parameter
  })
}

export function getAccountTeacherPage (parameter) {
  return axios({
    url: api.AccountTeacherPage,
    method: 'POST',
    data: parameter
  })
}

export function getAccountStudentPage (parameter) {
  return axios({
    url: api.AccountStudentPage,
    method: 'POST',
    data: parameter
  })
}

export function getAccountSchoolPage (parameter) {
  return axios({
    url: api.AccountSchoolPage,
    method: 'POST',
    data: parameter
  })
}
