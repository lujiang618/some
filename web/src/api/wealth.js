import request from '@/utils/request'

const api = {
  user: '/v1/user',
  role: '/v1/role',
  cost: '/v1/w/c/cost',
  costCategories: '/v1/w/c/cost/categories',
  permission: '/v1/permission',
  permissionNoPager: '/v1/permission/no-pager',
  orgTree: '/v1/org/tree'
}

export function getUserList (parameter) {
  return request({
    url: api.user,
    method: 'get',
    params: parameter
  })
}

export function getRoleList (parameter) {
  return request({
    url: api.role,
    method: 'get',
    params: parameter
  })
}

export function getCostList (parameter) {
  return request({
    url: api.cost,
    method: 'get',
    params: parameter
  })
}

export function getCostCategories (parameter) {
  return request({
    url: api.costCategories,
    method: 'get',
    params: parameter
  })
}

export function createCost (parameter) {
  return request({
    url: api.cost,
    method: 'post',
    data: parameter
  })
}

export function updateCost (parameter) {
  return request({
    url: api.orgTree,
    method: 'put',
    params: parameter
  })
}

export function saveService (parameter) {
  return request({
    url: api.service,
    method: parameter.id === 0 ? 'post' : 'put',
    data: parameter
  })
}

export function saveSub (sub) {
  return request({
    url: '/sub',
    method: sub.id === 0 ? 'post' : 'put',
    data: sub
  })
}
