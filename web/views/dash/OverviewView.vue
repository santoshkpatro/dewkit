<script setup>
import { onMounted } from 'vue'
const emit = defineEmits(['route-change'])

onMounted(() => {
  emit('route-change', 'overview')
})
</script>

<template>
  <div>Overview</div>
</template>
