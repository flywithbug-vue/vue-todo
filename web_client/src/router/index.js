import Vue from 'vue';
import Router from 'vue-router';
import App from "../App";
import LandingPage from "../components/marketing/LandingPage";
import About from "../components/marketing/About";
import Login from "../components/auth/Login";
import Register from "../components/auth/Register";
import TestTodosVariable from "../components/marketing/TestTodosVariable";
import Logout from "../components/auth/Logout";
import {store} from "../store/store";


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
      component:App,
      meta: {
        requiresAuth: true,
      }
    },
    {
      path:'/about',
      name:'about',
      component:About
    },
    {
      path:'/login',
      name:'login',
      component:Login,
      meta: {
        requiresVisitor: true,
      }
    },
    {
      path:'/register',
      name:'register',
      component:Register,
      meta: {
        requiresVisitor: true,
      }
    },
    {
      path:'/logout',
      name:'logout',
      component:Logout
    },
    {
      path:'/todo/item/:id',
      name:'item',
      component:TestTodosVariable
    }
  ],
  mode:'history'
})
