import { ref } from 'vue'
import { projectImboxWS } from '@/transport'

export function useImboxSocket() {
  const socket = ref(null)
  const isConnected = ref(false)

  function connect(projectId) {
    if (!projectId) return
    if (socket.value) return

    socket.value = projectImboxWS(projectId)

    socket.value.onopen = () => {
      isConnected.value = true
      console.log('[WS] Connected:', projectId)
    }

    socket.value.onclose = () => {
      isConnected.value = false
      socket.value = null
      console.log('[WS] Disconnected')
    }

    socket.value.onerror = (err) => {
      console.error('[WS] Error', err)
    }
  }

  function disconnect() {
    if (!socket.value) return
    socket.value.close()
    socket.value = null
  }

  return {
    socket,
    isConnected,
    connect,
    disconnect,
  }
}
