<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

// View that displays a detailed Hotel page.

// const props = defineProps({
// 	id: {
// 		type: String,
// 		required: true
// 	}
// })

async function getHotel(id) {
	const request = new Request(`http://localhost:3000/hotels/getbyid?id=${ id }`, {
		method: "GET",
	});

	try {
		const response = await fetch(request)
		if (!response.ok) {
			throw new Error(`Failed to get response`);}
		let hotel = await response.json();
		return hotel
	} catch (error) {
		console.error(error.message);
		alert("Failed")
		return;
	}
}

async function fetchData(id) {
	console.log(id)
	error.value = hotel.value = null
	loading.value = true

	try {
		hotel.value = await getHotel(id)
	} catch (err) {
		error.value = err.toString()
	} finally {
		loading.value = false
	}
}

const loading = ref(false)
const hotel = ref(null)
const error = ref(null)

const route = useRoute()

watch(() => route.params.id, fetchData, { immediate: true })

// const loadHotel = async () => {hotel.value = await getHotel();}
// loadHotel(props.id)

</script>
<template>
  <div
    v-if="loading"
    class="loading">
    Loading...
  </div>

  <div
    v-if="error"
    class="error">
    {{ error }}
  </div>

  <div
    v-if="hotel"
    class="content">
    <h2>Hello from {{ hotel.name }} page id = {{ hotel.id }}</h2>
    <p>{{ hotel.description }}</p>
  </div>
</template>
