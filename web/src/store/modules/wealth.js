import { getCostList, getCostCategories } from '@/api/wealth'

const cost = {
  actions: {
    CostList ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getCostList(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    CostCategories ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getCostCategories(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default cost
