import { getAnalyse } from '@/api/analyse'

const analyse = {
    actions: {
        GetAnalyse ({ commit }, params) {
            return new Promise((resolve, reject) => {
                getAnalyse(params).then(response => {
                    resolve(response)
                }).catch(error => {
                    reject(error)
                })
            })
        }
    }
}

export default analyse
