<script setup>
import { ref, computed, onMounted } from 'vue'
import { RouterView, useRouter, useRoute } from 'vue-router'
import { storeToRefs } from 'pinia'
import {
  LayoutDashboard,
  Inbox,
  AtSign,
  Ticket,
  LifeBuoy,
  Sliders,
  Star,
  Bug,
  Dot,
  Bell,
  Boxes,
} from 'lucide-vue-next'

import { useAuthStore } from '@/stores/auth'
import { useProjectStore } from '@/stores/project'
import { projectListAPI, projectMembersAPI } from '@/http'

const router = useRouter()
const route = useRoute()

const authStore = useAuthStore()
const projectStore = useProjectStore()

const { settings } = storeToRefs(authStore)
const { currentProject, members } = storeToRefs(projectStore)

const user = ref({
  name: 'John Doe',
  avatar: 'https://i.pravatar.cc/40?img=3',
})

const navItems = [
  { name: 'overview', title: 'Overview', icon: LayoutDashboard },
  { name: 'imbox', title: 'Imbox', icon: Inbox, count: 12 },
  { name: 'mentions', title: 'Mentions', icon: AtSign, count: 3 },
  { name: 'tickets', title: 'Tickets', icon: Ticket },
  { name: 'support', title: 'Help & Support', icon: LifeBuoy },
  { name: 'configure', title: 'Configure', icon: Sliders },
]

const activeRouteName = computed(() => route.name)

function navigate(item) {
  router.push({
    name: item.name,
    params: { projectId: route.params.projectId },
  })
}

const loadProjects = async () => {
  const { data } = await projectListAPI()
  projectStore.setProjects(data)

  const project = data.find((p) => p.id == route.params.projectId)
  projectStore.setCurrentProject(project)
}

const loadMembers = async () => {
  const { data } = await projectMembersAPI(route.params.projectId)

  projectStore.setMembers(
    data.map((m) => ({
      ...m,
      avatar: `https://i.pravatar.cc/32?img=${m.id}`,
      online: false,
    })),
  )
}

onMounted(async () => {
  projectStore.setActiveProject(route.params.projectId)
  await loadProjects()
  await loadMembers()
})
</script>

<template>
  <div v-if="!currentProject" class="layout">
    <aside class="sidebar">
      <a-skeleton active :paragraph="{ rows: 8 }" />
    </aside>

    <div class="main">
      <div class="topbar">
        <a-skeleton-input style="width: 240px" active />
      </div>

      <div class="content">
        <a-skeleton active :paragraph="{ rows: 6 }" />
      </div>
    </div>
  </div>

  <div v-else class="layout">
    <aside class="sidebar">
      <div class="sidebar-top">
        <h1 class="title">{{ settings.app.orgName }}</h1>

        <nav class="nav">
          <div
            v-for="item in navItems"
            :key="item.name"
            class="nav-item"
            :class="{ active: activeRouteName === item.name }"
            @click="navigate(item)"
          >
            <component :is="item.icon" :size="18" />
            <span class="label">{{ item.title }}</span>
            <span v-if="item.count" class="count">{{ item.count }}</span>
          </div>
        </nav>

        <div class="teammates">
          <p class="section-title">Teammates</p>

          <div v-for="member in members" :key="member.id" class="teammate">
            <img :src="member.avatar" />
            <span class="name">{{ member.fullName }}</span>
            <Dot :size="22" :class="member.online ? 'online' : 'offline'" />
          </div>
        </div>
      </div>

      <div class="sidebar-bottom">
        <div class="nav-item subtle">
          <Bug :size="16" />
          <span>Report an Issue</span>
        </div>

        <a href="https://github.com" target="_blank" class="nav-item subtle">
          <Star :size="16" />
          <span>Star on GitHub</span>
        </a>
      </div>
    </aside>

    <div class="main">
      <header class="topbar">
        <div class="top-left">
          <Boxes :size="18" />
          <div class="project">
            <span class="project-name">{{ currentProject.name }}</span>
            <span class="project-id">{{ currentProject.code }}</span>
          </div>
        </div>

        <div class="top-right">
          <a-input-search placeholder="Search" />
          <button class="icon-btn">
            <Bell :size="18" />
          </button>
          <div class="profile">
            <img :src="user.avatar" />
            <span>{{ user.name }}</span>
          </div>
        </div>
      </header>

      <main class="content">
        <RouterView />
      </main>
    </div>
  </div>
</template>

<style scoped>
.layout {
  display: flex;
  height: 100vh;
  font-family: system-ui, sans-serif;
}

.sidebar {
  width: 220px;
  background: #fafafa;
  border-right: 1px solid #eee;
  padding: 16px 12px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.sidebar-top {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.title {
  font-size: 16px;
  font-weight: 600;
  padding-left: 4px;
}

.nav {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px;
  font-size: 14px;
  border-radius: 6px;
  cursor: pointer;
}

.nav-item:hover {
  background: #f0f0f0;
}

.nav-item.active {
  background: #ececec;
  font-weight: 500;
}

.nav-item.subtle {
  color: #666;
}

.label {
  flex: 1;
}

.count {
  font-size: 12px;
  color: #999;
}

.teammates {
  margin-top: 6px;
}

.section-title {
  font-size: 12px;
  color: #888;
  padding-left: 8px;
  margin-bottom: 4px;
}

.teammate {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 6px 8px;
  font-size: 13px;
}

.teammate img {
  width: 24px;
  height: 24px;
  border-radius: 50%;
}

.name {
  flex: 1;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.online {
  color: #22c55e;
}

.offline {
  color: #bdbdbd;
}

.main {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.topbar {
  height: 56px;
  border-bottom: 1px solid #eee;
  padding: 8px 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.top-left {
  display: flex;
  align-items: center;
  gap: 10px;
}

.project {
  display: flex;
  flex-direction: column;
  line-height: 1.1;
}

.project-name {
  font-size: 16px;
  font-weight: 500;
}

.project-id {
  font-size: 10px;
  color: #888;
}

.top-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.icon-btn {
  background: none;
  border: none;
  cursor: pointer;
}

.profile {
  display: flex;
  align-items: center;
  gap: 8px;
}

.profile img {
  width: 28px;
  height: 28px;
  border-radius: 50%;
}

.content {
  flex: 1;
  padding: 16px;
}
</style>
