import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useProjectStore = defineStore('project', () => {
  const projects = ref([])

  const setProjects = (projectList) => {
    projects.values = projectList
  }

  return { projects, setProjects }
})
