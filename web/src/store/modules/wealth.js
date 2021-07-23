import { getCostList, getCostCategories, createCost, updateCost } from '@/api/wealth'

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
    },
    CreateCost ({ commit }, params) {
      return new Promise((resolve, reject) => {
        createCost(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    UpdateCost ({ commit }, params) {
      return new Promise((resolve, reject) => {
        updateCost(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default cost
