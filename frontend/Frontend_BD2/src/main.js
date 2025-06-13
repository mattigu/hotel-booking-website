import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css'

import vue3StarRatings from "vue3-star-ratings"

const app = createApp(App)

app.component('VueDatePicker', VueDatePicker);
app.component("Vue3StarRatings", vue3StarRatings)

app.provide('API_URL', 'http://localhost:3000')

app.use(router)

app.mount('#app')
