import Vue from 'vue';
import Router from 'vue-router';
import App from "../App";
import LandingPage from "../components/marketing/LandingPage";


Vue.use(Router)

export default new Router({
  routes: [
    {
      path:'/',
      name:'landing',
      component:LandingPage
    },
    {
      path:'/todo',
      name:'todo',
      component:App
    }
  ]
})
