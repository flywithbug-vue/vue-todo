<template>
  <div class="todo-item">
    <div class="todo-item-left">
      <input type="checkbox"
             v-model="completed"
             @change="doneEdit">
      <div v-if="!editing"
           @dblclick="editTodo"
           class="todo-item-label"
           :class="{ completed : completed }">{{ title }}</div>
      <input v-else
             class="todo-item-edit"
             type="text"
             v-model="title"
             @blur="doneEdit"
             @keyup.enter="doneEdit"
             @keyup.esc="cancelEdit" v-focus>
    </div>
    <div>
      <button @click="pluralize">Plural</button>
      <span class="remove-item"
            @click="removeTodo(index)">
      ✘
    </span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'todo-item',
  directives: {
    focus: {
      inserted: function(el) {
        el.focus();
      }
    }
  },
  props: {
    todo: {
      type: Object,
      required: true
    },
    index: {
      type: Number,
      required: true
    },
    checkAll: {
      type: Boolean,
      required: true
    }
  },
  data() {
    return {
      'id': this.todo.id,
      'title': this.todo.title,
      'completed': this.todo.completed,
      'editing': this.todo.editing,
      'beforeEditCache': ''
    };
  },
  watch: {
    checkAll() {
      this.completed = this.checkAll ? true : this.todo.completed;
    }
  },
  created() {
    eventBus.$on('pluralize', this.handlePluralize);
  },
  beforeDestroy() {
    eventBus.$off('pluralize', this.handlePluralize)
  },
  methods: {
    removeTodo(index) {
      eventBus.$emit('removeTodo', index);
    },
    editTodo() {
      this.beforeEditCache = this.title;
      this.editing = true;
    },
    doneEdit() {
      if (this.title.trim() == '') {
        this.title = this.beforeditCache;
      }
      this.beforeEditCache = this.title;
      this.editing = false;
      // this.todo.title = this.title
      eventBus.$emit('finishedEdit', {
        'index': this.index,
        'todo': {
          'id': this.id,
          'title': this.title,
          'completed': this.completed,
          'editing': this.editing
        }
      });
    },
    cancelEdit() {
      this.title = this.beforeditCache;
      this.editing = false;
    },
    pluralize() {
      eventBus.$emit('pluralize');
    },
    handlePluralize() {
      this.title = this.title + 's';
      eventBus.$emit('finishedEdit', {
        'index': this.index,
        'todo': {
          'id': this.id,
          'title': this.title,
          'completed': this.completed,
          'editing': this.editing
        }
      });
    }
  }
};
</script>

<style scoped>

</style>
