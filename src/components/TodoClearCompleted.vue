<template>
  <button v-if="showClearCompletedButton"
          @click="clearCompleted">Clear Completed</button>
</template>

<script>
export default {
  name: 'todo-clear-completed',
  computed: {
    showClearCompletedButton() {
      return this.$store.getters.showClearCompletedButton
    }
  },
  methods: {
    clearCompleted() {
      this.$store.commit('clearCompleted')
    }
  }
};
</script>

<style scoped>

</style>
