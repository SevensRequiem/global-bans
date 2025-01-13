import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/bans',
      name: 'bans',
      component: () => import('../views/BansView.vue'),
    },
    {
      path: '/appeal',
      name: 'appeal',
      component: () => import('../views/AppealView.vue'),
    },
    {
      path: '/servers',
      name: 'servers',
      component: () => import('../views/ServersView.vue'),
    }
  ],
})

export default router
