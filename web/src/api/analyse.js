import request from '@/utils/request'

const api = {
    analyse: '/v1/w/a/analyse',
    charts: '/v1/w/a/charts'
}

export function getAnalyse (parameter) {
    return request({
        url: api.analyse,
        method: 'post',
        data: parameter
    })
}

export function getStatistics (parameter) {
    return request({
        url: api.charts,
        method: 'post',
        data: parameter
    })
}
