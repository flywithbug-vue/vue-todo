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
             v-model.trim="title"
             @blur="doneEdit"
             @keyup.enter="doneEnterEdit"
             @keyup.esc="cancelEdit" v-focus>
    </div>
    <div>
      <button @click="pluralize">Plural</button>
      <span class="remove-item"
            @click="removeTodo()">
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
    removeTodo() {
      this.$store.dispatch('deleteTodo', this.id)
    },
    editTodo() {
      this.beforeEditCache = this.title;
      this.editing = true;
    },
    doneEnterEdit() {
      this.editing = false
      if (this.title.trim() == '') {
        this.title = this.beforeEditCache;
      }
    },
    doneEdit() {
      this.editing = false
      if (this.title.trim() == '') {
        this.title = this.beforeEditCache;
      }
      if (this.title == this.beforeEditCache){
        return
      }

        this.$store.dispatch('updateTodo',{
        'id': this.id,
        'title': this.title,
        'completed': this.completed,
        'editing': this.editing
      })
    },
    cancelEdit() {
      this.title = this.beforeEditCache;
      this.editing = false;
    },
    pluralize() {
      eventBus.$emit('pluralize');
    },
    handlePluralize() {
      this.title = this.title + 's';
      this.$store.dispatch('updateTodo',{
        'id': this.id,
        'title': this.title,
        'completed': this.completed,
        'editing': this.editing
      })
    },
  }
};
</script>

<style scoped>

</style>
