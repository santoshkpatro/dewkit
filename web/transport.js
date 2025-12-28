import axios from 'axios'

const http = axios.create({
  baseURL: '/api',
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
})

export const authMetaAPI = () => http.get('/auth/meta')
export const authProfileAPI = () => http.get('/auth/profile')
export const authLoginAPI = (data) => http.post('/auth/login', data)

export const projectListAPI = () => http.get('/projects')
export const projectMembersAPI = (projectId) => http.get(`/projects/${projectId}/members`)

export const conversationListAPI = (projectId, params) =>
  http.get(`/projects/${projectId}/conversations`, {
    params,
  })

function createWebSocket(path) {
  const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws'
  const host = window.location.host

  const normalizedPath = path.startsWith('/') ? path : `/${path}`

  return new WebSocket(`${protocol}://${host}/ws${normalizedPath}`)
}

export function projectImboxWS(projectId) {
  return createWebSocket(`/projects/${projectId}/imbox`)
}
