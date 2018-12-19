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
            this.$store.dispatch('loginAuth',{
              account:this.account,
              password:this.password,
            })
          }
      }
    }
</script>

<style >

  label {
    display: block;
    margin-bottom: 4px;
  }

  .login-heading {
    margin-bottom: 16px;
  }

  .form-control {
    margin-bottom: 24px;
  }

  .mb-more {
    margin-bottom: 42px;
  }

  .login-form {
    max-width: 500px;
    margin: auto;
  }

  .login-input {
    width: 100%;
    font-size: 16px;
    padding: 12px 16px;
    outline: 0;
    border-radius: 3px;
    border: 1px solid lightgrey;
  }

  .btn-submit {
    width: 100%;
    padding: 14px 12px;
    font-size: 18px;
    font-weight: bold;
    background: #60BD4F;
    color: white;
    border-radius: 3px;
    cursor: pointer;

  /*&:hover {*/
     /*background: darken(#60BD4F, 10%);*/
   /*}*/

  }

</style>
