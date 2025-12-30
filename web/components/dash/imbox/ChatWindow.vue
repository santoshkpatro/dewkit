<script setup>
import { ref, onMounted, watch, computed } from 'vue'
import { Send, User } from 'lucide-vue-next'
import { conversationMessageListAPI, conversationMessageCreateAPI } from '@/transport'
import { useProjectStore } from '@/stores/project'
import { storeToRefs } from 'pinia'

const props = defineProps({
  conversationId: Number,
  projectId: Number,
})

const projectStore = useProjectStore()
const { activeChatMessages, activeConversationId } = storeToRefs(projectStore)

const reply = ref('')

const loadMessages = async () => {
  const { data } = await conversationMessageListAPI(props.projectId, props.conversationId, {})

  projectStore.setActiveChatMessages(data)
}

const conversation = computed(() => {
  return projectStore.conversations.find((c) => c.id === props.conversationId)
})

const sendReply = async () => {
  if (!reply.value.trim()) return

  const postData = {
    senderType: 'staff',
    body: reply.value,
  }

  reply.value = ''
  await conversationMessageCreateAPI(props.projectId, props.conversationId, postData)
}

onMounted(() => {
  projectStore.setActiveConversation(props.conversationId)
  loadMessages()
})

watch(() => props.conversationId, loadMessages)
</script>

<template>
  <a-layout style="height: 100%">
    <!-- Header -->
    <a-layout-header
      style="
        background: none;
        border-bottom: 1px solid #f0f0f0;
        display: flex;
        align-items: center;
        gap: 8px;
      "
    >
      <a-avatar>
        <User size="16" />
      </a-avatar>
      <strong>{{ conversation.customerFullName }}</strong>
    </a-layout-header>

    <!-- Messages -->
    <a-layout-content style="padding: 16px; overflow-y: auto">
      <div
        v-for="msg in activeChatMessages"
        :key="msg.id"
        :style="{
          display: 'flex',
          justifyContent: msg.senderType === 'staff' ? 'flex-end' : 'flex-start',
          marginBottom: '8px',
        }"
      >
        <div
          :style="{
            maxWidth: '70%',
            padding: '8px 12px',
            borderRadius: '8px',
            background: msg.senderType === 'staff' ? '#e6f4ff' : '#fafafa',
            border: '1px solid #f0f0f0',
          }"
        >
          {{ msg.body }}
        </div>
      </div>
    </a-layout-content>

    <!-- Reply box -->
    <a-layout-footer style="border-top: 1px solid #f0f0f0; padding: 12px; display: flex; gap: 8px">
      <a-input v-model:value="reply" placeholder="Reply as support…" @pressEnter="sendReply" />
      <a-button type="primary" @click="sendReply">
        <Send size="16" />
      </a-button>
    </a-layout-footer>
  </a-layout>
</template>
