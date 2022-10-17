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
          <a-button size="mini" type="text" @click="JumpToFrontEnd">前台</a-button>
          <a-button size="mini" type="text" @click="Logout">退出</a-button>
        </div>
      </div>
    </div>

    <div class="main">
      <!--      侧边栏导航-->
      <div class="sidebar">
        <a-menu
            :style="{ width: '200px', height: '100%' }"
            :default-open-keys="['1']"
            :default-selected-keys="['/backend/blogs']"
            show-collapse-button
            breakpoint="xl"
            @menu-item-click="changeMenuItem"
        >
          <!--            点击菜单项时触发 @menu-item-click -->
        <a-sub-menu key="0">
          <template #icon>
            <icon-apps></icon-apps>
          </template>
          <template #title>文章管理</template>
          <a-menu-item key="/backend/blogs">文章列表</a-menu-item>
          <a-menu-item key="/backend/tags">标签管理</a-menu-item>
        </a-sub-menu>
        </a-menu>
      </div>
      <!--      内容页面-->
      <div class="content">
        <RouterView/>
      </div>
    </div>
  </div>
</template>

<script setup>
import {useRouter} from 'vue-router';

const router = useRouter()

const changeMenuItem = (key) => {
  router.push(key);
    console.log(key);
}

const JumpToFrontEnd = ()=>{
  router.push('/');
}

const Logout=()=>{
  localStorage.removeItem("username");
  localStorage.removeItem("password");
  router.push("/login")
}
</script>

<style scoped>
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


.main {
  display: flex;
  align-content: center;
  justify-content: flex-start;
  align-items: flex-start;
  height: calc(100vh - 45px);
}

.sidebar {
  height: 100%;
  border-right: 1px solid rgb(229, 230, 235);
}

.content {
  margin: 8px;
  width: 100%;
  height: 100%;
}

</style>