<template>
    <div>
        Logout
    </div>
</template>

<script>
export default {
  name: "Logout",
  created() {
    this.$store.dispatch('clearTodos')
    this.$store.dispatch('destroyToken').then(() => {
      this.$router.push({name: 'home'})
    }).catch(error => {
      this.$router.push({name: 'home'})
    })
  }

}
</script>

<style scoped>

</style>
