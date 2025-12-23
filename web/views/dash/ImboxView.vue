<script setup>
import { onMounted } from 'vue'
const emit = defineEmits(['route-change'])

onMounted(() => {
  emit('route-change', 'imbox')
})
</script>

<template>
  <div>Imbox</div>
</template>
