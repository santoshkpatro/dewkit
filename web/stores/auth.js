import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

import { authMetaAPI } from '@/transport'

export const useAuthStore = defineStore('auth', () => {
  const settings = ref(null)
  const loggedInUser = ref(null)

  const isLoggedIn = computed(() => !!loggedInUser.value)

  const loginUser = (loginData) => {
    loggedInUser.value = loginData
  }

  const setupAuth = async () => {
    try {
      const { data } = await authMetaAPI()
      settings.value = data.settings
      loggedInUser.value = data.loggedInUser
    } catch (error) {
      console.log('Error', error)
    }
  }

  return { settings, loggedInUser, isLoggedIn, setupAuth, loginUser }
})
