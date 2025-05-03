import { createWebHistory, createRouter } from 'vue-router'

import HotelBrowse from '@/views/HotelBrowse.vue'
import HotelView from '@/views/HotelView.vue'

const routes = [
  { path: '/', component: HotelBrowse },
  { path: '/hotel/:id', component: HotelView, props: true },
// Components can be lazy loaded as well, it might be useful later
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
