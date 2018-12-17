import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'


Vue.use(Vuex)
axios.defaults.baseURL="http://localhost:6201/api"

export const store = new Vuex.Store({
  state: {
    filter: 'all',
    todos:[]
  },
  getters: {
    remaining(state) {
      return state.todos.filter(todo => !todo.completed).length;
    },
    anyRemaining(state, getters) {
      return getters.remaining != 0;
    },
    todosFiltered(state) {
      if (state.filter == 'all'){
        return state.todos;
      } else if (state.filter == 'active'){
        return state.todos.filter(todo => !todo.completed)
      } else if (state.filter == 'completed'){
        return state.todos.filter(todo => todo.completed)
      }
      return state.todos;
    },
    showClearCompletedButton(state) {
      return state.todos.filter(todo => todo.completed).length > 0
    }
  },
  mutations: {
    retrieveTodos(state,todos){
      state.todos = todos
    },
    addTodo(state, todo){
      state.todos.push({
        id: todo.id,
        title: todo.title,
        completed: false,
        editing: false,
      })
    },
    updateTodo(state, todo){
      const index = state.todos.findIndex(item => item.id == todo.id)
      state.todos.splice(index, 1,{
        'id': todo.id,
        'title': todo.title,
        'completed': todo.completed,
        'editing': todo.editing
      })
    },
    updateFilter(state,filter) {
      state.filter = filter
    },
    checkAll(state,checked) {
      state.todos.forEach((todo) => todo.completed = checked);
    },
    deleteTodo(state, id) {
      const index  = state.todos.findIndex(item => item.id == id)
      state.todos.splice(index,1);
    },
    clearCompleted(state) {
      state.todos = state.todos.filter(todo => !todo.completed)
    },
  },
  actions: {
    retrieveTodos(context) {
      axios.get('/todo/list').then(response => {
        context.commit('retrieveTodos',response.data.data.list)
      }).catch(error => {
        console.log(error)
      })
    },
    addTodo(context, todo){
      setTimeout(() => {
        context.commit('addTodo', todo)
      },1000)
    },
    updateTodo(context, todo){
      context.commit('updateTodo', todo)
    },
    updateFilter(context,filter) {
      context.commit('updateFilter',filter)

    },
    checkAll(context,checked) {
      context.commit('checkAll',checked)
    },
    deleteTodo(context, id) {
      context.commit('deleteTodo',id)
    },
    clearCompleted(context) {
      context.commit('clearCompleted')
    },
  }
})
