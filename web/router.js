import { createRouter, createWebHistory } from 'vue-router'

import IndexView from '@/views/IndexView.vue'
import LoginView from '@/views/LoginView.vue'
import DashView from '@/views/DashView.vue'
import ImboxView from '@/views/dash/ImboxView.vue'
import OverviewView from '@/views/dash/OverviewView.vue'
import CreateView from '@/views/CreateView.vue'

import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      component: LoginView,
      name: 'login',
    },
    {
      path: '/create',
      component: CreateView,
      name: 'create',
      meta: { requiresAuth: true },
    },
    {
      path: '/:projectId/dash',
      component: DashView,
      meta: { requiresAuth: true },
      children: [
        {
          path: 'overview',
          component: OverviewView,
          name: 'overview',
          meta: { requiresAuth: true },
        },
        {
          path: 'imbox',
          component: ImboxView,
          name: 'imbox',
          meta: { requiresAuth: true },
        },
      ],
    },
    {
      path: '',
      component: IndexView,
      name: 'index',
      meta: { requiresAuth: true },
    },
  ],
})

router.beforeEach((to) => {
  const auth = useAuthStore()

  // Route requires login but user is not logged in
  if (to.meta.requiresAuth && !auth.isLoggedIn) {
    return {
      name: 'login',
      query: { redirect: to.fullPath },
    }
  }
})

export default router
