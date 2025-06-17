import { createWebHistory, createRouter } from 'vue-router'

import HotelView from '@/views/HotelView.vue'
import HotelBrowseView from '@/views/HotelBrowseView.vue'

const routes = [
	{ path: '/', name:"HotelBrowseView" ,component: HotelBrowseView },
	{ path: '/hotel/:id', name:"HotelView", component: HotelView, props: true },
// Components can be lazy loaded as well, it might be useful later
]

const router = createRouter({
	scrollBehavior(_to, _from, savedPosition) {
		if (savedPosition) {
			return savedPosition
		} else {
			return { top: 0 }
		}
	},
	history: createWebHistory(),
	routes,
})

export default router
