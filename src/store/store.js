import Vue from 'vue'
import Vuex from 'vuex'


Vue.use(Vuex)


export const store = new Vuex.Store({
  state: {
    filter: 'all',
    todos:[
      {
        'id': 1,
        'title': 'Finish Vue Screencast',
        'completed': false,
        'editing': false
      },
      {
        'id': 2,
        'title': 'Take over world',
        'completed': false,
        'editing': false
      },
      {
        'id': 3,
        'title': 'hello world',
        'completed': false,
        'editing': false
      }
    ]
  }
})
