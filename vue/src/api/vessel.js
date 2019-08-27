import request from '@/utils/request'

const baseUrl = '/api/vessel'

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
