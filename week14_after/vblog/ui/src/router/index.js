import { createRouter, createWebHistory } from "vue-router";
import BlogView from "../views/frontend/BlogView.vue";
import BlogList from "../views/backend/BlogList.vue";
import TagList from "../views/backend/TagList.vue";

import FrontendLayout from "../layout/FrontendLayout.vue"
import BackendLayout from "../layout/BackendLayout.vue"
import {beforeEachHandler,afterEachHandler} from "./permission";




const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      redirect:"/frontend/blogs",
    },
    {
      path: "/backend",
      name: "backend",
      component: BackendLayout,
      children:[
        {
          path: "blogs",
          name: "BlogList",
          component: BlogList,
        },
        {
          path:"blog_create",
          name:"BlogCreate",
          component: ()=> import("@/views/backend/BlogCreate.vue"),
        },
        {
          path: "tags",
          name: "TagList",
          component: TagList,
        },
      ]
    },
    {
      path: "/frontend",
      name: "frontend",
      component: FrontendLayout,
      children: [
        {
          path: "blogs",
          name: "BlogView",
          component: BlogView,
        },
        {
          path:"blogs/:id",
          name:"BlogDetail",
          component: ()=> import("@/views/frontend/BlogDetail.vue")
        }
      ],
    },
    {
      path: "/login",
      name: "LoginPage",
      component: () => import("@/views/login/LoginPage.vue"),
    },
    {
      path: "/errors/403",
      name: "PermissionDeny",
      component: () => import("@/views/errors/PermissionDeny.vue"),
    },
    {
      path:"/:pathMatch(.*)*",
      name:"NotFound",
      component:() => import("@/views/errors/NotFound.vue")
    }
  ],
});

//补充导航守卫
router.beforeEach(beforeEachHandler)
router.afterEach(afterEachHandler)

export default router;
