import { ref, watch } from 'vue'
import { defineStore } from 'pinia'
import { useImboxSocket } from '@/composables/useImboxSocket'

export const useProjectStore = defineStore('project', () => {
  const projects = ref([])
  const currentProject = ref(null)
  const members = ref([])

  const currentProjectId = ref(localStorage.getItem('currentProject'))

  const { socket, isConnected, connect, disconnect } = useImboxSocket()

  const connectImbox = (projectId) => {
    connect(projectId)
  }

  const disconnectImbox = () => {
    disconnect()
  }

  const setProjects = (projectList) => {
    projects.value = projectList
  }

  const setCurrentProject = (projectData) => {
    currentProject.value = projectData
  }

  const setCurrentProjectId = (projectId) => {
    currentProjectId.value = projectId
    localStorage.setItem('currentProject', projectId)
  }

  const removeCurrentProjectId = () => {
    currentProjectId.value = null
    localStorage.removeItem('currentProject')
  }

  const setMembers = (memberList) => {
    members.value = memberList
  }

  watch(currentProjectId, (id) => {
    if (!id || !projects.value.length) return
    const project = projects.value.find((p) => p.id == id)
    if (project) {
      currentProject.value = project
    }
  })

  return {
    projects,
    currentProject,
    currentProjectId,
    members,
    socket,
    isConnected,
    connectImbox,
    disconnectImbox,
    setProjects,
    setCurrentProject,
    setCurrentProjectId,
    removeCurrentProjectId,
    setMembers,
  }
})
