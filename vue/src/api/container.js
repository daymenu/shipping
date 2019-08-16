import request from '@/utils/request'

export function page(data) {
  return request({
    url: '/container/containerService/page',
    method: 'get',
    params: data
  })
}

export function get(data) {
  return request({
    url: '/container/containerService/get',
    method: 'get',
    params: data
  })
}

export function use(data) {
  return request({
    url: '/container/containerService/use',
    method: 'post',
    data
  })
}
export function giveback(data) {
  return request({
    url: '/container/containerService/giveback',
    method: 'post',
    data
  })
}
export function create(data) {
  return request({
    url: '/container/containerService/create',
    method: 'post',
    data
  })
}
export function update(data) {
  return request({
    url: '/container/containerService/update',
    method: 'post',
    data
  })
}
