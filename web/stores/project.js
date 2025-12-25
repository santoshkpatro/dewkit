import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useProjectStore = defineStore('project', () => {
  const projects = ref([])

  const setProjects = (projectList) => {
    projects.values = projectList
  }

  const setActiveProject = (projectId) => {
    localStorage.setItem('currentProject', projectId)
  }

  const removeActiveProject = () => {
    localStorage.removeItem('currentProject')
  }

  return { projects, setProjects, setActiveProject, removeActiveProject }
})
