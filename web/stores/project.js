import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useProjectStore = defineStore('project', () => {
  const projects = ref([])

  const members = ref([])

  const setProjects = (projectList) => {
    projects.value = projectList
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

  return { projects, setProjects, setActiveProject, removeActiveProject, setMembers, members }
})
