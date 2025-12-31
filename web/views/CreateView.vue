<script setup>
import { ref } from 'vue'
import { User, FileText } from 'lucide-vue-next'
import { projectCreateAPI } from '@/transport'
import { useRouter } from 'vue-router'

const router = useRouter()

const formState = ref({
  name: '',
  description: '',
})

const rules = {
  name: [{ required: true, message: 'Name is required' }],
  description: [],
}

const submitForm = async () => {
  const { data } = await projectCreateAPI(formState.value)
  router.push({ name: 'imbox', params: { projectId: data.id } })
}
</script>

<template>
  <div class="container">
    <a-card title="Create Project" class="card">
      <a-form layout="vertical" :model="formState" :rules="rules" @finish="submitForm">
        <a-form-item label="Project Name" name="name">
          <a-input v-model:value="formState.name" placeholder="Enter project name">
            <template #prefix>
              <User size="16" />
            </template>
          </a-input>
        </a-form-item>

        <a-form-item label="Description" name="description">
          <a-textarea
            v-model:value="formState.description"
            placeholder="Enter description"
            :rows="4"
          >
            <template #prefix>
              <FileText size="16" />
            </template>
          </a-textarea>
        </a-form-item>

        <a-button type="primary" html-type="submit" block> Create </a-button>
      </a-form>
    </a-card>
  </div>
</template>

<style scoped>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.card {
  width: 360px;
}
</style>
