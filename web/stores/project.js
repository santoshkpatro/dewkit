import { ref, watch } from 'vue'
import { defineStore } from 'pinia'
import { useImboxSocket } from '@/composables/useImboxSocket'

export const useProjectStore = defineStore('project', () => {
  const projects = ref([])
  const currentProject = ref(null)
  const currentProjectId = ref(localStorage.getItem('currentProject'))

  const members = ref([])

  const conversations = ref([])
  const activeConversationId = ref(null)

  const activeChatMessages = ref([])

  const { socket, isConnected, connect, disconnect, send } = useImboxSocket()

  function handleSocketMessage(event) {
    if (event.type !== 'message.new') return

    const { conversationId, message } = event.payload

    conversations.value = conversations.value.map((c) =>
      c.id === conversationId ? { ...c, lastMessage: message } : c,
    )

    if (activeConversationId.value === conversationId) {
      activeChatMessages.value.push(message)
    }
  }

  function handleSocketClose() {
    console.log('[Store] Imbox socket closed')
  }

  const connectImbox = (projectId) => {
    if (!projectId) return
    connect(projectId, handleSocketMessage, handleSocketClose)
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
    disconnectImbox()
  }

  const setMembers = (memberList) => {
    members.value = memberList
  }

  const setConversations = (conversationsList) => {
    conversations.value = conversationsList
  }

  const setActiveConversation = (conversationId) => {
    activeConversationId.value = conversationId
  }

  const setActiveChatMessages = (messageList) => {
    activeChatMessages.value = messageList
  }

  watch(currentProjectId, (id, oldId) => {
    if (!id || id === oldId) return

    disconnectImbox()
    connectImbox(id)

    const project = projects.value.find((p) => p.id === id)
    if (project) currentProject.value = project
  })

  return {
    projects,
    currentProject,
    currentProjectId,
    members,
    conversations,
    activeConversationId,
    activeChatMessages,
    socket,
    isConnected,
    send,
    connectImbox,
    disconnectImbox,
    setProjects,
    setCurrentProject,
    setCurrentProjectId,
    removeCurrentProjectId,
    setMembers,
    setConversations,
    setActiveConversation,
    setActiveChatMessages,
  }
})
