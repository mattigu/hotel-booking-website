import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

const app = createApp(App)
app.provide('API_URL', 'http://localhost:3000')
app.use(router)
app.mount('#app')
