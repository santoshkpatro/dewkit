import { createRouter, createWebHistory } from 'vue-router'

import LoginView from '@/views/LoginView.vue'
import DashView from '@/views/DashView.vue'
import ImboxView from '@/views/dash/ImboxView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      component: LoginView,
    },
    {
      path: '/dash',
      component: DashView,
      children: [
        {
          path: 'imbox',
          component: ImboxView,
        },
      ],
    },
  ],
})

export default router
