<script setup>
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import { Mail, Lock, Github } from 'lucide-vue-next'

import { authLoginAPI, projectListAPI } from '@/transport'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()
const form = reactive({
  email: '',
  password: '',
})

const submit = async () => {
  try {
    const loginResp = await authLoginAPI(form)
    authStore.loginUser(loginResp.data)

    const projectsResp = await projectListAPI()

    // router.push({ name: 'imbox' })
  } catch (error) {
    console.log('Error', error)
  }
}
</script>

<template>
  <div class="page">
    <!-- Brand -->
    <div class="brand">DewKit</div>

    <!-- Login Box -->
    <div class="login-box">
      <div class="heading">Log in to your account</div>

      <div class="field">
        <label>Email</label>
        <a-input v-model:value="form.email" type="email">
          <template #prefix>
            <Mail size="16" />
          </template>
        </a-input>
      </div>

      <div class="field">
        <label>Password</label>
        <a-input-password v-model:value="form.password">
          <template #prefix>
            <Lock size="16" />
          </template>
        </a-input-password>
      </div>

      <a-button type="primary" block @click="submit"> Login </a-button>
    </div>

    <!-- Footer center -->
    <div class="footer-center">support@dewkit.app</div>

    <!-- Footer right -->
    <a
      class="footer-right"
      href="https://github.com/santoshkpatro/dewkit"
      target="_blank"
      rel="noopener noreferrer"
    >
      <Github size="16" />
      <span>GitHub</span>
    </a>
  </div>
</template>

<style scoped>
.page {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
}

/* Brand */
.brand {
  position: absolute;
  top: 24px;
  left: 32px;
  font-size: 26px;
  font-weight: 500;
  color: #9aa0a6;
}

/* Login box */
.login-box {
  width: 360px;
  display: flex;
  flex-direction: column;
  gap: 18px;
}

/* Heading */
.heading {
  font-size: 24px;
  font-weight: 500;
  color: #111;
  margin-bottom: 6px;
}

/* Fields */
.field {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.field label {
  font-size: 13px;
  color: #666;
}

/* Footer */
.footer-center {
  position: absolute;
  bottom: 24px;
  font-size: 13px;
  color: #aaa;
}

.footer-right {
  position: absolute;
  bottom: 22px;
  right: 32px;
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #999;
  text-decoration: none;
}

.footer-right:hover {
  color: #555;
}
</style>
