<template>
  <div class="about">
   <DemoArchive
       :name="demoArchiveName"
       greeting-message="test"
       v-model:count="count"
   ></DemoArchive>
    <button @click="jumpToHome">Home</button>
    <button @click="jumpToTest">Test</button>
    <span>鼠标坐标： {{ x }}, {{ y }}</span>

    用户名：<input type="text" v-model="login.username">
    密码：<input type="password" v-model="login.password">
<!--    <button @click="submit">登录</button>-->
    <el-button type="small" @click="submit">登录</el-button>
    <div>
      <el-button>Default</el-button>
      <el-button type="primary">Primary</el-button>
      <el-button type="success">Success</el-button>
      <el-button type="info">Info</el-button>
      <el-button type="warning">Warning</el-button>
      <el-button type="danger">Danger</el-button>
    </div>
</div>
</template>

<script setup>
import DemoArchive from "@/components/DemoArchive.vue";
import {inject, onMounted} from "vue";
import { useMouse } from "@/stores/mouse.js";
// import { useMouse } from "@/vueuse/core";
import { useCount } from "@/stores/count.js";
// import { login } from "@/stores/login.js";
import { useRouter } from "vue-router";
// import { useStorage } from "@vueuse/core";
import { useUserStore } from "../stores/counter.js";
import {nextTick} from "vue";

const demoArchiveName="点击我";

//从根节点 获取注入的全局变量count
// const count = inject("count",0)
// 从全局的组合式函数中获取状态
const count = useCount()

//
// const login = useStorage('login',{username:'',password:'',expired:0})
const login = useUserStore();
console.log(login)

const submit = () =>{
  nextTick(() =>{
    // console.log("login test")
    console.log(login.value.username,login.value.password)
  })
};



//获取到Router对象
// vue2的使用方式：this.$router
const router = useRouter()

const jumpToHome = ()=>{
  router.push("/")
}

const jumpToTest = ()=>{
  // name route name
  // params 路由参数，定义路由的时候的路由变量
  // query query string , url参数
  router.push({name: "test",query:{name:"a"},params:{id:10}});
}

const { x,y } =useMouse();
// const handleCount = (value) => {
//   //接收到来自子组件的count事件
//   count.value=value
//   console.log(value);
// };
</script>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}


</style>