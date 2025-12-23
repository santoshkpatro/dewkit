<script setup>
import { useRouter } from 'vue-router'
import { onMounted } from 'vue'

import { projectListAPI } from '@/http'
import { useProjectStore } from '@/stores/project'

const router = useRouter()
const projectStore = useProjectStore()

const loadProjects = async () => {
  const { data } = await projectListAPI()
  projectStore.setProjects(data)
}

onMounted(async () => {
  await loadProjects()
  if (projectStore.projects.length == 0) {
    router.push({ name: 'create' })
  }
})
</script>

<template></template>
