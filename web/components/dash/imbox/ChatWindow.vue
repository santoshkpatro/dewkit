<script setup>
import { ref, computed } from 'vue'
import { Layout, Avatar, Input, Button } from 'ant-design-vue'
import { Send, User } from 'lucide-vue-next'

const props = defineProps({
  conversationId: {
    type: Number,
    required: true,
  },
})

const messagesByConversation = {
  1: [
    { from: 'user', text: 'I need help with my subscription' },
    { from: 'agent', text: 'Sure — what plan are you on?' },
  ],
  2: [{ from: 'user', text: 'Thanks for resolving this!' }],
  3: [{ from: 'user', text: 'Can you reset my password?' }],
}

const messages = ref(messagesByConversation[props.conversationId] || [])
const reply = ref('')

const sendReply = () => {
  if (!reply.value.trim()) return
  messages.value.push({ from: 'agent', text: reply.value })
  reply.value = ''
}
</script>

<template>
  <Layout style="height: 100%">
    <!-- Header -->
    <Layout.Header
      style="
        background: none;
        border-bottom: 1px solid #f0f0f0;
        display: flex;
        align-items: center;
        gap: 8px;
      "
    >
      <Avatar>
        <User size="16" />
      </Avatar>
      <strong>Conversation #{{ conversationId }}</strong>
    </Layout.Header>

    <!-- Messages -->
    <Layout.Content style="padding: 16px; overflow-y: auto">
      <div
        v-for="(msg, index) in messages"
        :key="index"
        :style="{
          display: 'flex',
          justifyContent: msg.from === 'agent' ? 'flex-end' : 'flex-start',
          marginBottom: '8px',
        }"
      >
        <div
          style="max-width: 70%; padding: 8px 12px; border: 1px solid #f0f0f0; border-radius: 8px"
        >
          {{ msg.text }}
        </div>
      </div>
    </Layout.Content>

    <!-- Reply box -->
    <Layout.Footer style="border-top: 1px solid #f0f0f0; padding: 12px; display: flex; gap: 8px">
      <Input v-model:value="reply" placeholder="Reply as support…" @pressEnter="sendReply" />
      <Button type="primary" @click="sendReply">
        <Send size="16" />
      </Button>
    </Layout.Footer>
  </Layout>
</template>
