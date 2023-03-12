<template>
  <n-message-provider>
    <MessageApi />
  </n-message-provider>

  <n-notification-provider>
    <notification-api />
  </n-notification-provider>

  <n-layout content-style="padding: 24px">
    <n-grid cols="12" item-responsive responsive="screen">
      <n-grid-item span="0 m:2"></n-grid-item>
      <n-grid-item span="12 m:8">
        <router-view />
      </n-grid-item>
    </n-grid>
  </n-layout>
</template>

<script setup>
import MessageApi from "@/components/MessageApi";
import NotificationApi from "@/components/NotificationApi";
</script>
<style></style>
