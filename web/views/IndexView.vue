<script setup>
import { useRouter } from 'vue-router'
import { onMounted } from 'vue'

import { projectListAPI } from '@/http'
import { useProjectStore } from '@/stores/project'

const router = useRouter()
const projectStore = useProjectStore()

const loadProjects = async () => {
  const { data } = await projectListAPI()
  return data
}

onMounted(async () => {
  const projects = await loadProjects()

  if (projects.length == 0) {
    router.push({ name: 'create' })
    return
  }

  const currentProjectId = parseInt(localStorage.getItem('currentProject'))
  const currentProject = projects.find((p) => p.id === currentProjectId)

  if (!currentProject) {
    projectStore.removeActiveProject()
  }

  router.push({ name: 'imbox', params: { projectId: projects[0].id } })
})
</script>

<template></template>
