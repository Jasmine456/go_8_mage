<template>
  <div class="app">
    <div>
      <a-breadcrumb>
        <a-breadcrumb-item>文章管理</a-breadcrumb-item>
        <a-breadcrumb-item>创建文章</a-breadcrumb-item>
      </a-breadcrumb>
<!--      文章创建表单-->
      <div class="form">
      <a-form :model="form" layout="vertical">
<!--        标题-->
        <a-form-item field="title_name" label="标题">
          <a-input
              v-model="form.title_name"
              placeholder="请输入文章的标题"/>
        </a-form-item>
<!--        概述-->
        <a-form-item field="title_name" label="概述">
          <a-textarea
              v-model="form.summary"
              placeholder="请输入文章概述"
              allow-clear/>
        </a-form-item>
<!--        内容-->
        <a-form-item field="content" label="内容">
          <md-editor
              v-model="form.content"
              @onSave="handleSubmit"
          />
        </a-form-item>
      </a-form>
      </div>
    </div>
  </div>
</template>


<script setup>
import {onMounted, reactive} from "vue";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import { CREATE_BLOG,UPDATE_BLOG } from "@/api/blog.js"
import { Message } from "@arco-design/web-vue";
import {useRoute} from "vue-router";
import {GET_BLOG} from "../../api/blog";


//判断是否有id参数，如果有，请该id的参数
const route=useRoute()
onMounted(async ()=>{
if (route.query.id){
  try{
    const resp = await GET_BLOG(route.query.id)
    form.id=resp.data.id
    form.title_name=resp.data.title_name
    form.content=resp.data.content
    form.summary=resp.data.summary
    console.log(resp)
  }catch (error){
    Message.error(`查询文章异常：${error.response.data.message}`);
  }
}
})

//表单数据
const form = reactive({
  id:null,
  title_name:"",
  summary:"",
  content:"",
})

const handleSubmit = async ()=>{
  //可以手动校验
  if (form.id !== null){
  //  需要更新
    try {
      const resp = await UPDATE_BLOG(form.id,form);
      Message.success(`文章：${resp.data.title_name} 已保存`)


    }catch (error){
      Message.error(`保存文章异常：${error.response.data.message}`)
    }
  }else{
  //需要保存
  try {
    const resp = await CREATE_BLOG(form);
    // console.log(resp)
    Message.success(`文章：${resp.data.title_name} 已保存`)
    form.id=resp.data.id
  }catch(error){
    // console.log(error)
    Message.error(`保存文章异常：${error.response.data.message}`)
  }
}
}

</script>

<style  scoped>
.app{
  width: 100%;
}

.form{
  margin-top: 10px ;
  display: flex;
}

</style>