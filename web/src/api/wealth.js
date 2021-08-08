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
    url: api.cost,
    method: 'put',
    data: parameter
  })
}
