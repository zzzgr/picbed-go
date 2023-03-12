import { createRouter, createWebHashHistory } from "vue-router";
import NProgress from "nprogress";

const routes = [
  {
    path: "/",
    name: "index",
    component: () => import("@/views/Index.vue"),
    meta: {
      title: "pic-go",
    },
  },
  {
    path: "/404",
    name: "404",
    component: () => import("@/views/404"),
    meta: {
      title: "404",
    },
  },
  {
    path: "/:pathMatch(.*)",
    redirect: "/404",
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  NProgress.start();
  if (to.meta.title) {
    document.title = `${to.meta.title} - ${process.env.VUE_APP_TITLE}`;
  } else {
    document.title = process.env.VUE_APP_TITLE;
  }
  next();
});

router.afterEach(() => {
  NProgress.done();
});

export default router;
