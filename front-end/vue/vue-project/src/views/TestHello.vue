<template>
<div>
  <button  @click="count++" type="text"> You click me {{ count }} </button>
  <span>{{ count }}</span>
</div>
</template>

<script setup>
import { ref } from "vue";
let count = ref(0)
</script>

<style lang="scss" scoped></style>
