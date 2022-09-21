<template>
  <div>
    Blog List Page
  </div>
  <div>
    <ul v-for="item in data" :key="item.id">
      <li>{{ item.title_name }}</li>
    </ul>
  </div>
</template>

<script setup>
// 通过axios这个http的客户端，获取来自server提供的数据
import axios from "axios";
import { ref,onMounted } from "vue";
import { LIST_BLOG } from "../api/blog.js";



let data = ref([])

const fetchData = async  () => {
  console.log("fetch data ...");
  try{
    const resp = await LIST_BLOG()
  data.value = resp.data.Items;
  }catch (error){
    console.log(error);
  }finally {
    console.log("fetch data complete")
  }

  // axios.get("/vblog/api/v1/blog").then((resp) => {
  //   data.value = resp.data.Items;
  //   console.log(resp.data.Items);
  // }).catch((error) => {
  //   console.log(error);
  // }).finally(() => {
  //   console.log("fetch data complete");
  // });
};

//比如需要把resp的数据展示的页面上，需要页面已经加载完成
onMounted(() => {
  fetchData();
});
</script>

<style lang="scss" scoped>

</style>