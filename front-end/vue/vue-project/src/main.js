import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";

// vue-router插件,配置路由条目
import router from "./router";

import "./assets/main.css";

const app = createApp(App);

//注册全局组件
import TestHello from "../src/views/TestHello.vue"
app.component("TestHello",TestHello)

//加载全局状态管理插件
// 1. 引入依赖
const pinia = createPinia()
import piniaPersist from 'pinia-plugin-persist'
// 2.配置store实例使用piniaPersist插件

pinia.use(piniaPersist)
app.use(pinia);

//加载UI插件
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
app.use(ElementPlus)

//加载路由
app.use(router);
// console.log(app)

//自定义全局指令
//v-focus
app.directive('focus',{
    mounted:(el) =>{
        el.focus();
    }
});

app.mount("#app");

