<script setup>
import { inject, ref } from 'vue'
import { useFetch } from '@/composables/useFetch'

import HotelPage from '@/components/HotelPage.vue'

const API_URL = inject('API_URL')
// View that displays a detailed Hotel page.

const props = defineProps({
	id: {
		type: String,
		required: true
	}
})


async function loadHotel() {
	const request = new Request(`${API_URL}/get/hotel/${props.id}?guests=${1}`, { method: 'GET' })
	const { data, execute } = useFetch(request)
	await execute()
	hotel.value = data.value
}

const hotel = ref()
loadHotel()

</script>

<template>
  <div v-if="hotel" class="hotel_view">
    <HotelPage :hotel="hotel" />
  </div>
</template>
