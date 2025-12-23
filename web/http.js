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
