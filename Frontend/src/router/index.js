// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '@/views/LoginPage.vue'
import ProfilePage from '@/views/ProfilePage.vue'
import CoursesPage from '@/views/CoursesPage.vue'
import TasksPage from '@/views/TasksPage.vue'
import ProgressPage from '@/views/ProgressPage.vue'

export const routes = [
  { path: '/',   component: LoginPage },
  { path: '/profile', component: ProfilePage },
  { path: '/courses', component: CoursesPage },
  { path: '/tasks',   component: TasksPage },
  { path: '/progress',component: ProgressPage },
]

export default createRouter({
  history: createWebHistory(),
  routes
})