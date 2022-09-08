import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";

import "./assets/main.css";

const app = createApp(App);

//注册全局组件
import TestHello from "../src/views/TestHello.vue"
app.component("TestHello",TestHello)

app.use(createPinia());
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

