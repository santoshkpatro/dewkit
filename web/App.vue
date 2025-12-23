<script setup>
import { RouterView } from 'vue-router'

const theme = {
  token: {
    colorPrimary: '#104084',
    fontSize: 13,
    borderRadius: 4,
    colorBgLayout: '#f7f7f5',
    colorBgContainer: '#ffffff',
  },
}
</script>

<template>
  <a-config-provider :theme="theme">
    <router-view></router-view>
  </a-config-provider>
</template>

<style scoped></style>
