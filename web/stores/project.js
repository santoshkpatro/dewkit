import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useProjectStore = defineStore('project', () => {
  const projects = ref([])

  const currentProject = ref(null)
  const members = ref([])

  const setProjects = (projectList) => {
    projects.value = projectList
  }

  const setCurrentProject = (projectData) => {
    currentProject.value = projectData
  }

  const setActiveProject = (projectId) => {
    localStorage.setItem('currentProject', projectId)
  }

  const removeActiveProject = () => {
    localStorage.removeItem('currentProject')
  }

  const setMembers = (memberList) => {
    members.value = memberList
  }

  return {
    projects,
    currentProject,
    setCurrentProject,
    setProjects,
    setActiveProject,
    removeActiveProject,
    setMembers,
    members,
  }
})
