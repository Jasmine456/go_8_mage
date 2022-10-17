<template>
  <div class="login-form">

    <a-form :model="formData" @submit="handleSubmit" :style="{ width: '450px' }">
      <!--      登录说明-->
      <a-form-item class="logo">登录后台进行博客管理</a-form-item>
      <!--      用户名输入框-->
      <a-form-item field="username" label="" hide-asterisk :rules="[{required:true,message:'请输入用户名'}]">
        <a-input v-model="formData.username" placeholder="请输入用户名">
          <template #prepend>
            <icon-user />
          </template>

        </a-input>
      </a-form-item>
      <!--      用户密码输入框-->
      <a-form-item field="password" label="" hide-asterisk :rules="[{required:true,message:'请输入用户密码'}]">
        <a-input type="password" v-model="formData.password" placeholder="请输入用户密码">
          <template #prepend>
            <icon-lock />
          </template>
        </a-input>
      </a-form-item>
      <a-form-item>
        <a-button type="primary" class="login-btn" html-type="submit">登录</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup>
import {reactive} from "vue";
import {useRouter} from "vue-router";
import {Message} from "@arco-design/web-vue";
import {LOGIN} from "../../api/user";

const router = useRouter()

const formData = reactive({username: "", password: ""})
const handleSubmit = async (data) => {
  if (data.errors == undefined) {
    //校验提交的数据 data.values
    data = data.values
    try {
      const resp = await LOGIN(data)
      console.log(resp)
      //  登陆成功，保存状态
      localStorage.setItem("username", data.username)
      localStorage.setItem("password", data.password)
      //  需要跳转到后台管理页面
      // router.push('/backend');

      //  如何获取当前Location的URL参数
      //  登陆成功直接跳转到后台页，如果URL页面带有重定向参数，则直接路由到重定向的页面
      const {redirect, ...othersQuery} = router.currentRoute.value.query;
      router.push({
        name: redirect || "BlogList",
        query: {...othersQuery},
      });
    } catch (error) {
      if (error.response.data){
        Message.error(error.response.data.message)
      }else{
        Message.error(error.message)
      }
    }
  }


  //   if (data.username === "admin" && data.password === "123456") {
  //     //  登陆成功，保存状态
  //     localStorage.setItem("username", data.username)
  //     localStorage.setItem("password", data.password)
  //     //  需要跳转到后台管理页面
  //     // router.push('/backend');
  //
  //     //  如何获取当前Location的URL参数
  //     //  登陆成功直接跳转到后台页，如果URL页面带有重定向参数，则直接路由到重定向的页面
  //     const {redirect, ...othersQuery} = router.currentRoute.value.query;
  //     router.push({
  //       name: redirect || "BlogList",
  //       query: {...othersQuery},
  //     });
  //   } else {
  //     //  弹窗消息 提醒用户
  //     Message.error("用户名或者密码有误")
  //   }
  // }
};

</script>

<style scoped>
.login-form {
  height: 100%;
  display: flex;
  align-content: center;
  justify-content: center;
  align-items: center;

}

.login-btn {
  width: 100%;
}

.logo {
  display: flex;
  font-size: 18px;
  font-weight: 500;
  align-items: center;
  justify-content: center;
  align-content: center;
  width: 100%;
}


</style>