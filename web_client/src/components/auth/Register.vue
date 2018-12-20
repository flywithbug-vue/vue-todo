<template>
<div class="login-form">
  <h2 class="login-heading">Register</h2>
  <form action="#" @submit.prevent="register">

    <div class="form-control">
      <label for="account">Account</label>
      <input type="text" name="account" id="account" class="login-input" v-model="account">
    </div>

    <div class="form-control">
      <label for="email">Email</label>
      <input type="email" name="email" id="email" class="login-input" v-model="email">
    </div>
    <div class="form-control mb-more">
      <label for="password">Password</label>
      <input type="password" name="password" id="password" class="login-input" v-model="password">
    </div>

    <div class="form-control">
      <button type="submit" class="btn-submit">Create Account</button>
    </div>

  </form>
</div>
</template>

<script>
  export default {
    data() {
      return {
        account: '',
        password: '',
        email:'',
      }
    },
    methods: {
      register() {
        this.$store.dispatch('register', {
          account: this.account,
          password: this.password,
          email: this.email,
        }).then(() => {
            this.$router.push({ name: 'login' })
          }).catch(err => {
            console.log(err)
        })
      }
    }
  }
</script>

