<template>
  <div style="text-align: center; font-size: 24px">页面不存在</div>
</template>

<script setup>
import { onMounted } from "vue";
import router from "@/router";

onMounted(() => {
  window.$message.warning("页面不存在，正在跳转首页...");
  setTimeout(() => {
    router.push("/");
  }, 3000);
});
</script>

<style scoped></style>
