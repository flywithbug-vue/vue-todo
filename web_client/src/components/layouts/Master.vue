<template>
    <div id="app">
      <ul class="nav">
        <li><a href="/">Home</a></li>
        <li><a href="/">About</a> </li>
        <li><a href="/">Login</a></li>
        <li><a href="/">Register</a></li>
      </ul>
      <router-view></router-view>
    </div>
</template>

<script>
export default {
}
</script>

<style>
  *{
    box-sizing: border-box;
    margin: 0px;
    padding: 0px;
  }

  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #2c3e50;
    font-size: 24px;
    height: 100vh;
  }

  .flex-center {
    display: flex;
    justify-content: center;
  }

  .nav {
    display: flex;
    list-style: none;
    padding: 15px 0;
    justify-content: flex-end;
    background: #F5F8FA;
    border-bottom: 1px solid lightgrey;
    margin: 0 0 24px;
  }

  .nav a {
    color: #636b6f;
    padding: 0 25px;
    font-size: 14px;
    font-weight: 600;
    letter-spacing: .1rem;
    text-decoration: none;
    text-transform: uppercase;
  }


</style>
