import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";

import "./assets/main.css";

const app = createApp(App);

// 引入UI组件 arco-design
import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';
// 额外引入图标库
import ArcoVueIcon from '@arco-design/web-vue/es/icon';

app.use(ArcoVueIcon);
app.use(ArcoVue);
app.use(createPinia());
app.use(router);

app.mount("#app");
