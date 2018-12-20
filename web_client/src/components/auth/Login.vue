<template>
  <div class="login-form">
    <h2 class="login-heading">Login</h2>
    <form action="#">
      <div class="form-control">
        <label for="account">Account</label>
        <input type="text" name="account" id="account" class="login-input" v-model="account">
      </div>
      <div class="form-control mb-more">
        <label for="password">Password</label>
        <input type="password" name="password" id="password" class="login-input" v-model="password">
      </div>
    </form>
    <div class="form-control">
      <button type="submit" class="btn-submit" @click="login">Login</button>
    </div>
  </div>
</template>

<script>
    export default {
        name: "Login",
      data() {
          return {
            account: '',
            password: '',
          }
      },
      methods: {
          login() {
            this.$store.dispatch('retrieveToken',{
              account:this.account,
              password:this.password,
            }).then(response => {
              this.$router.push({name: 'todo'})
            })
          }
      }
    }
</script>

<style >


</style>
