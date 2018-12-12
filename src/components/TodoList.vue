<template>
  <div >
    <input type="text"
           class="todo-input"
           placeholder="what needs to be done"
           v-model="newTodo"
           @keyup.enter="addTodo">
    <div v-for="(todo, index) in todos"
         :key="todo.id"
         class="todo-item">
      <div class="todo-item-left">
        <input type="checkbox" v-model="todo.completed">
        <div class="todo-item-label" v-if="!todo.editing"
              @dblclick="editTodo(todo)"
              :class="{ completed : todo.completed}">
          {{ todo.title }}</div>
        <input v-if="todo.editing"
               @blur="doneEdit(todo)"
               @keyup.enter="doneEdit(todo)"
               @keyup.esc="cancelEdit(todo)"
               class="todo-item-edit"
               type="text"
               v-model="todo.title" v-focus>
      </div>
      <div class="remove-item"
           @click="removeTodo(index)">
        &times;
      </div>
    </div>
    <div class="extra-container">
      <div><label><input type="checkbox"
                         :checked="!anyRemaining"
                         @change="checkAllTodos">Check All</label></div>
      <div>{{ remaining }} items left</div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'todo-list',
    directives: {
      focus: {
        inserted: function(el) {
          el.focus();
        }
      }
    },
    data () {
          return {
            newTodo: '',
            beforeditCache:'',
            idForTodo:4,
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
          };
        },
    computed: {
      remaining() {
        return this.todos.filter(todo => !todo.completed).length
      },
      anyRemaining() {
        return this.remaining != 0
      }
    },
    methods: {
      addTodo() {
        if (this.newTodo.trim().length == 0) {
          return;
        };
        this.todos.push({
          id: this.idForTodo,
          title:this.newTodo,
          completed:false
        });
        this.newTodo = '';
        this.idForTodo++;
      },
      removeTodo(index) {
        this.todos.splice(index,1);
      },
      editTodo(todo) {
        this.beforeditCache = todo.title
        todo.editing = true
      },
      doneEdit(todo) {
        if (todo.title.trim() == '') {
          todo.title = this.beforeditCache
        };
        todo.editing = false
      },
      cancelEdit(todo) {
        todo.title =  this.beforeditCache
        todo.editing = false
      },
      checkAllTodos() {
        this.todos.forEach( (todo) => todo.completed = event.target.checked)
      }
    }
  };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss">

  .todo-input {
    width: 100%;
    padding: 10px 18px;
    font-size: 18px;
    margin-bottom: 16px;

    &:focus {
      outline: 0;
    }
  }

  .todo-item {
    margin-bottom: 12px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    animation-duration: 0.3s;
    font-size: 16px
  }

  .remove-item {
    cursor: pointer;
    margin-left: 14px;
    &:hover {
      color: black;
    }
  }

  .todo-item-left {
    display: flex;
    align-items: center;
  }

  .todo-item-label {
    padding: 10px;
    border: 1px solid white;
    margin-left: 12px;
  }

  .todo-item-edit {
    font-size: 24px;
    color: #2c3e50;
    margin-left: 12px;
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc; //override defaults
    font-family: "Avenir", Helvetica, Arial, sans-serif;

    &:focus {
      outline: none;
    }
  }

  .completed {
    text-decoration: line-through;
    color: grey;
  }

  .extra-container {
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 16px;
    border-top: 1px solid lightgrey;
    padding-top: 14px;
    margin-bottom: 14px;
  }

  .name-container {
    margin-bottom: 16px;
  }

  button {
    font-size: 14px;
    background-color: white;
    appearance: none;
    padding: 4px;

    &:hover {
      background: lightgreen;
    }

    &:focus {
      outline: none;
    }
  }

  .active {
    background: lightgreen;
  }

  // CSS Transitions
  .fade-enter-active,
  .fade-leave-active {
    transition: opacity 0.2s;
  }

  .fade-enter,
  .fade-leave-to {
    opacity: 0;
  }
</style>
