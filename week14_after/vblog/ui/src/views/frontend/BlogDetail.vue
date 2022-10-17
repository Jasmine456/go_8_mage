<template>
  <div>
    <h2 class="title">
      {{ data.title_name }}
    </h2>
    <div class="describe">
      <div>发表于: {{ data.publish_at }}</div>
      <a-divider direction="vertical" />
      <div>作者: {{ data.author }}</div>
    </div>
    <div class="summary">
      {{ data.summary }}
    </div>
    <div style="width: 700px">
      <md-editor
          v-model="data.content"
          preview-only
      />
    </div>
  </div>
</template>

<script setup>
import {onMounted, ref} from "vue";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import {GET_BLOG} from "../../api/blog";
import {useRoute} from "vue-router";
const data = ref({});

onMounted(async ()=>{
  try {
    // id从路由里面获取,读取的路由对象
    const route = useRoute()
    console.log(route)
    const resp = await GET_BLOG(route.params.id);
    data.value = resp.data
  }catch (error){
    console.log(error)
  }
});


// const title = "测试";
// const publishAt = 10000;
// const author = "测试";
// const text = ref(`## 参考
//
// + [博客样式参考](http://www.arccode.net/)
// + [md-editor-v3 官方页面](https://imzbf.github.io/md-editor-v3/index)
// + [md-editor-v3 Github](https://github.com/imzbf/md-editor-v3)`);
</script>

<style lang="less" scoped>
.title {
  display: flex;
  justify-content: center;
  align-items: center;
  align-content: center;
  color: #555;
}

.describe {
  display: flex;
  justify-content: center;
  align-items: center;
  align-content: center;
  color: #999;
  font-size: 12px;
}

.summary{
  margin:8px;
  color: #555;
}
</style>