import http from '../utils/request'

// login post request
export function LoginReq(data: any) {
  return http.request({
    url: '/v1/login',
    method: 'post',
    data
  })
}

// list request
export function ListReq() {
  return http.request({
    url: '/v1/list',
    method: 'post'
  })
}

// register request
export function RegisterReq(data: any) {
  return http.request({
    url: '/v1/register',
    method: 'post',
    data: data
  })
}