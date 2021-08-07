import request from '@/utils/request'

const api = {
    analyse: '/v1/w/a/analyse'
}

export function getAnalyse (parameter) {
    return request({
        url: api.analyse,
        method: 'post',
        data: parameter
    })
}
