<template>
  <div>
    <!--  顶部导航-->
    <div class="header">
      <!--    左侧LOGO区域-->
      <div class="logo">
      我的博客
      </div>
      <!--    右侧 操作区域-->
      <div class="right-header">
        <div>
<!--          登录后台进行博客管理-->
          <a-button size="mini" type="text" @click="jumpToBackend">后台管理</a-button>
        </div>
      </div>
    </div>

    <!--  内容页面-->
    <div class="content">
      <RouterView />
    </div>
  </div>
</template>

<script setup>
import {useRouter} from "vue-router";

const router=useRouter()
//跳转到后台管理系统
const jumpToBackend = ()=>{
//  1. 检查用户是否登录
  const username = localStorage.getItem("username")
  const password = localStorage.getItem("password")
  if (
      username !== null && password !== null && username !== "" && password !==""
  ){
    //  用户已登录 ---》直接跳转后台管理页面
    router.push("/backend/blogs")
  }else{
    //  用户未登录 ---》重定向到Login页面
    router.push("/login")
  }


}
</script>

<style  scoped>
.header {
  display: flex;
  align-content: center;
  justify-content: flex-start;
  align-items: center;
  border-bottom: 1px solid rgb(229, 230, 235);
  height: 45px;
}

.logo {
  margin-left: 8px;
  font-size: 14px;
  font-weight: 500;
}

.right-header {
  margin-left: auto;
}

.content {
  margin: 20px;
  display: flex;
  align-content: center;
  justify-content: center;
  align-items: center;
}
</style>