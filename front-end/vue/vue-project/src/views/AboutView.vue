<template>
  <div class="about">
   <DemoArchive name="demoArchiveName" greeting-message="test"></DemoArchive>
</div>
</template>

<script setup>
import DemoArchive from "../components/DemoArchive.vue"

const demoArchiveName="点击我"
</script>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}


</style>