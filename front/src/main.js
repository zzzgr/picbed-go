import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import NProgress from "nprogress";
import "nprogress/nprogress.css";

import naive from "naive-ui";

NProgress.configure({
  easing: "ease", // 动画方式
  speed: 500, // 递增进度条的速度
  showSpinner: false, // 是否显示加载ico
  trickleSpeed: 200, // 自动递增间隔
  minimum: 0.3, // 初始化时的最小百分比
});

createApp(App).use(router).use(naive).mount("#app");
