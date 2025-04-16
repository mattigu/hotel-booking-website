import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'


const app = createApp(App)
app.provide('API_URL', 'http://localhost:3000')
app.mount('#app')
