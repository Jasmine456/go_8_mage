import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style

// 打印路由详情
// console.log(import.meta)


const router = createRouter({
  //环境变量配置
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      //直接到路由配置，该视图会直接加载
      component: HomeView,
    },
    {
      path: "/about",
      name: "about",
      // alias:"/about_page",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/AboutView.vue"),
    },
    {
      // /test/1,/test/2,params id=1,id=2
      // /test?name=a query {name=a}
      path: "/test:id",
      name: "test",
      component: () => import("../views/TestHello.vue"),
    },
    {
      path:"/blogs",
      name:"BlogList",
      component: () => import("../views/BlogList.vue"),

    },
    {
      //vue2 --> vue3
      // * --> :pathMatch(.*)
      path:"/:pathMatch(.*)*",
      name:"NotFound",
      component: ()=>import("../views/NotFound.vue")
    },
  ],
});

// // 定义Router对象的全局守卫（路由进入前）
// router.beforeEach((to,from ,next) =>{
//   // start progress bar
//   NProgress.start()
//   console.log("beforeEach to",to)
//   console.log("beforeEach from",from)
//   next()
// });
//
// // 定义Router对象的全局守卫（路由完成）
// router.afterEach((to,from) =>{
//   // finish progress bar
//   NProgress.done()
//   console.log("afterEach to",to)
//   console.log("afterEach from",from)
// });

export default router;
