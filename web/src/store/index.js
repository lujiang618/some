import Vue from 'vue'
import Vuex from 'vuex'

import app from './modules/app'
import user from './modules/user'
import wealth from './modules/wealth'
import apps from './modules/apps'
import version from './modules/version'
import analyse from './modules/analyse'

// default router permission control
import permission from './modules/permission'

// dynamic router permission control (Experimental)
// import permission from './modules/async-router'
import getters from './getters'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    app,
    user,
    wealth,
    analyse,
    apps,
    version,
    permission
  },
  state: {

  },
  mutations: {

  },
  actions: {

  },
  getters
})
