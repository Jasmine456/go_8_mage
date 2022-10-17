<template>
  <div class="content">
    <a-space direction="vertical" style="width: 700px">
      <div  v-for="item in blogs" :key="item.id">
      <BlogItem
          v-if="item.status === 'published'"
          :id="item.id"
          :title="item.title_name"
          :publishAt="item.publish_at"
          :author="item.author"
          :summary="item.summary"
          ></BlogItem>
      </div>
    </a-space>
  </div>
</template>

<script setup>
import BlogItem from "./BlogItem.vue"
import {LIST_BLOG} from "../../api/blog";
import {onMounted, reactive, ref} from "vue";

const req = reactive({
  keywords:"",
  page_number:1,
  page_size:500,
  //前台页面，只过滤出已发布的文章
  status:"published",
})

const blogs = ref([]);

onMounted( async ()=>{
  try {
    const resp = await LIST_BLOG(req)
    blogs.value=resp.data.Items;
    // console.log(resp);
  } catch (error) {
    console.log(error);
  }
})


</script>

<style lang="scss" scoped></style>