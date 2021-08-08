import { getEarningList, createEarning, updateEarning } from '@/api/earning'

const earning = {
    actions: {
      CostList ({ commit }, params) {
        return new Promise((resolve, reject) => {
            getEarningList(params).then(response => {
                resolve(response)
            }).catch(error => {
                reject(error)
            })
        })
      },
      CreateEarning ({ commit }, params) {
          console.log('params:', params)
        return new Promise((resolve, reject) => {
            createEarning(params).then(response => {
                resolve(response)
            }).catch(error => {
                reject(error)
            })
        })
      },
      UpdateEarning ({ commit }, params) {
        return new Promise((resolve, reject) => {
          updateEarning(params).then(response => {
            resolve(response)
          }).catch(error => {
            reject(error)
          })
        })
      }
    }
}

export default earning
