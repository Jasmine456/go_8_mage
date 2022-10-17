<template>
  <div class="content">
    <a-result class="result" status="404" :subtitle="'你访问的页面不存在'">
      <template #extra>
        <a-button key="back" type="primary" @click="back"> 返回主页 </a-button>
      </template>
    </a-result>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router"
const router = useRouter()
const back = () =>{
  router.push('/')
}

</script>

<style lang="less" scoped>
.content{
  position: absolute;
  top: 50%;
  left: 50%;
  margin-top: -121px;
  margin-left: -95px;
  text-align: center;
}
</style>