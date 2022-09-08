<template>
  <div>
      Demo Archive
    <span> {{ greetingMessage }} </span>
    <button @click="count++">{{ name }}</button>
    <span>{{ count }}</span>
  </div>
</template>

<script setup>
import { ref } from "vue";
//定义一个属性
const pros = defineProps({
  name:{type:String,required:true},
  greetingMessage:String,
});

let count=ref(0)
</script>

<style lang="scss" scoped></style>

