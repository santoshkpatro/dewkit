<script setup>
import { ref } from 'vue'
import { RouterView, useRouter, useRoute } from 'vue-router'
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

const router = useRouter()
const route = useRoute()

const { settings } = useAuthStore()

const activeRouteName = ref(null)

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

const teammates = ref([
  { name: 'Alice Johnson', avatar: 'https://i.pravatar.cc/32?img=5', online: true },
  { name: 'Bob Robertson The Third', avatar: 'https://i.pravatar.cc/32?img=8', online: false },
  { name: 'Eve', avatar: 'https://i.pravatar.cc/32?img=12', online: true },
])

function navigate(item) {
  activeRouteName.value = item.name
  router.push({ name: item.name, params: { projectId: route.params.projectId } })
}

function onChildRouteChange(routeName) {
  activeRouteName.value = routeName
}
</script>

<template>
  <div class="layout">
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
            <span v-if="item.count" class="count">
              {{ item.count }}
            </span>
          </div>
        </nav>

        <div class="teammates">
          <p class="section-title">Teammates</p>

          <div v-for="mate in teammates" :key="mate.name" class="teammate">
            <img :src="mate.avatar" />
            <span class="name">{{ mate.name }}</span>
            <Dot :size="24" :class="mate.online ? 'online' : 'offline'" />
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
            <span class="project-name">DWAPP</span>
            <span class="project-id">YSAH344JSF32</span>
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
        <RouterView @route-change="onChildRouteChange" />
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
  color: #333;
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

.name {
  flex: 1;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.teammate img {
  width: 24px;
  height: 24px;
  border-radius: 50%;
}

.online {
  color: #22c55e;
}

.offline {
  color: #bdbdbd;
}

.sidebar-bottom {
  display: flex;
  flex-direction: column;
  gap: 4px;
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
  align-items: center;
  justify-content: space-between;
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
  font-size: 14px;
  font-weight: 500;
}

.project-id {
  font-size: 12px;
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
  padding: 4px;
}

.profile {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
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
