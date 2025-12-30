import { ref } from 'vue'
import { projectImboxWS } from '@/transport'

const socket = ref(null)
const isConnected = ref(false)

export function useImboxSocket() {
  function connect(projectId, onMessage, onClose) {
    if (!projectId || socket.value) return

    socket.value = projectImboxWS(projectId)

    socket.value.onopen = () => {
      isConnected.value = true
      console.log('[WS] Connected:', projectId)
    }

    socket.value.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data)
        onMessage?.(data)
      } catch (err) {
        console.error('[WS] Invalid message', err)
      }
    }

    socket.value.onclose = () => {
      isConnected.value = false
      socket.value = null
      onClose?.()
      console.log('[WS] Disconnected')
    }

    socket.value.onerror = (err) => {
      console.error('[WS] Error', err)
    }
  }

  function send(payload) {
    if (!socket.value || socket.value.readyState !== WebSocket.OPEN) return
    socket.value.send(JSON.stringify(payload))
  }

  function disconnect() {
    socket.value?.close()
    socket.value = null
  }

  return {
    socket,
    isConnected,
    connect,
    send,
    disconnect,
  }
}
