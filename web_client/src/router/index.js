import Vue from 'vue';
import Router from 'vue-router';
import App from "../App";
import LandingPage from "../components/marketing/LandingPage";
import About from "../components/marketing/About";
import Login from "../components/auth/Login";
import Register from "../components/auth/Register";
import TestTodosVariable from "../components/marketing/TestTodosVariable";


Vue.use(Router)

export default new Router({
  routes: [
    {
      path:'/',
      name:'home',
      component:LandingPage
    },
    {
      path:'/todo',
      name:'todo',
      component:App
    },
    {
      path:'/about',
      name:'about',
      component:About
    },
    {
      path:'/login',
      name:'login',
      component:Login
    },
    {
      path:'/register',
      name:'register',
      component:Register
    },
    {
      path:'/todo/:id',
      name:'todo',
      component:TestTodosVariable
    }
  ],
  mode:'history'
})
