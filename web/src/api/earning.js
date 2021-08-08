import request from '@/utils/request'

const api = {
    earning: '/v1/w/e/earning'
}

export function getEarningList (parameter) {
    return request({
        url: api.earning,
        method: 'get',
        params: parameter
    })
}

export function createEarning (parameter) {
    console.log('parameter:', parameter)
    return request({
      url: api.earning,
      method: 'post',
      data: parameter
    })
 }

export function updateEarning (parameter) {
    return request({
        url: api.earning,
        method: 'put',
        data: parameter
    })
}
