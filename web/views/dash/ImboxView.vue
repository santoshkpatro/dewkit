<script setup>
import { ref, onMounted, watch } from 'vue'
import { Layout, List, Avatar, Dropdown, Menu, Empty } from 'ant-design-vue'
import { Filter, Circle, Clock, CheckCircle, Archive } from 'lucide-vue-next'
import ChatWindow from '@/components/dash/imbox/ChatWindow.vue'
import { conversationListAPI } from '@/transport'
import { useRoute } from 'vue-router'
import { useProjectStore } from '@/stores/project'
import { storeToRefs } from 'pinia'

const route = useRoute()
const emit = defineEmits(['route-change'])

const projectStore = useProjectStore()
const { conversations } = storeToRefs(projectStore)

const activeConversationId = ref(null)
const activeFilter = ref('open')

const loadConversations = async () => {
  const { data } = await conversationListAPI(route.params.projectId, { status: activeFilter.value })
  projectStore.setConversations(data)
}

onMounted(() => {
  emit('route-change', 'inbox')
  loadConversations()
})

/* ✅ watch filter change */
watch(activeFilter, () => {
  activeConversationId.value = null
  loadConversations()
})

const filterConfig = {
  open: { label: 'Open', icon: Circle },
  pending: { label: 'Pending', icon: Clock },
  resolved: { label: 'Resolved', icon: CheckCircle },
  archived: { label: 'Archived', icon: Archive },
}
</script>

<template>
  <Layout class="imbox-layout">
    <Layout.Sider width="320" theme="light" class="imbox-sider">
      <div class="imbox-filter">
        <strong>Inbox</strong>

        <Dropdown trigger="click">
          <div class="filter-trigger">
            <component :is="filterConfig[activeFilter].icon" size="16" />
            <span>{{ filterConfig[activeFilter].label }}</span>
            <Filter size="14" />
          </div>

          <template #overlay>
            <Menu :selectedKeys="[activeFilter]" @click="({ key }) => (activeFilter = key)">
              <Menu.Item key="open">
                <div class="menu-item"><Circle size="14" /> Open</div>
              </Menu.Item>

              <Menu.Item key="pending">
                <div class="menu-item"><Clock size="14" /> Pending</div>
              </Menu.Item>

              <Menu.Item key="resolved">
                <div class="menu-item"><CheckCircle size="14" /> Resolved</div>
              </Menu.Item>

              <Menu.Item key="archived">
                <div class="menu-item"><Archive size="14" /> Archived</div>
              </Menu.Item>
            </Menu>
          </template>
        </Dropdown>
      </div>

      <List :data-source="conversations" item-layout="horizontal" class="imbox-list">
        <template #renderItem="{ item }">
          <List.Item class="imbox-list-item" @click="activeConversationId = item.id">
            <List.Item.Meta :title="item.customerFullName" :description="item.lastMessage.body">
              <template #avatar>
                <Avatar>{{ item.customerFullName.charAt(0) }}</Avatar>
              </template>
            </List.Item.Meta>

            <!-- <Badge v-if="item.unread" :count="item.unread" /> -->
          </List.Item>
        </template>
      </List>
    </Layout.Sider>

    <Layout.Content class="imbox-content">
      <ChatWindow v-if="activeConversationId" :conversation-id="activeConversationId" />
      <Empty v-else description="Select a conversation" />
    </Layout.Content>
  </Layout>
</template>

<style scoped>
.imbox-layout {
  height: 100%;
}

.imbox-sider {
  border-right: 1px solid #f0f0f0;
}

.imbox-filter {
  padding: 12px 16px;
  border-bottom: 1px solid #f0f0f0;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.filter-trigger {
  display: flex;
  align-items: center;
  gap: 6px;
  cursor: pointer;
}

.menu-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.imbox-list {
  overflow-y: auto;
}

.imbox-list-item {
  cursor: pointer;
}

.imbox-content {
  height: 100%;
}
</style>
