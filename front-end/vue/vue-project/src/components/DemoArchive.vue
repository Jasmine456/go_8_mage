<template>
  <div>
    <span> {{ greetingMessage }} </span>
    <button @click="clickCount">{{ name }}</button>
    <span>{{ count }}</span>
  </div>
</template>

<script setup>
//定义一个属性,defineProps() 是一个宏由编译器负责处理，无需引入
const pros = defineProps({
  name:{type:String,default:"默认值"},
  greetingMessage:String,
  count:Number,
});

//定义组件的事件
const emit = defineEmits(["update:count"])

const clickCount = () =>{
  // count.value++;
//  子组件向父组件发送事件
  emit("update:count",pros.count+1);
};

//此时count这个变量，维护再子组件内部的，count变成一个属性，由父组件传递过来
//如果count变量有变更，然后通过事件通知 父组件，从而实现单向数据流


console.log(pros);

// let count=ref(0)
</script>

<style lang="scss" scoped></style>

