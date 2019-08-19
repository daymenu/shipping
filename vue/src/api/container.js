import request from '@/utils/request'

const baseUrl = '/api/containerService'

export function page(data) {
  return request({
    url: baseUrl + '/page',
    method: 'get',
    params: data
  })
}

export function get(data) {
  return request({
    url: baseUrl + '/get',
    method: 'get',
    params: data
  })
}

export function use(data) {
  return request({
    url: baseUrl + '/use',
    method: 'post',
    data
  })
}
export function giveback(data) {
  return request({
    url: baseUrl + '/giveback',
    method: 'post',
    data
  })
}
export function create(data) {
  return request({
    url: baseUrl + '/create',
    method: 'post',
    data
  })
}
export function update(data) {
  return request({
    url: baseUrl + '/update',
    method: 'post',
    data
  })
}
