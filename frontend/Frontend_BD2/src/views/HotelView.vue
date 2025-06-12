<script setup>
import { ref } from 'vue'

// View that displays a detailed Hotel page.

const props = defineProps({
	id: {
		type: String,
		required: true
	}
})

async function getHotel(id) {
	const request = new Request(`http://localhost:3000/hotels/getbyid/${ id }`, {
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

const hotel = ref(null)
const loadHotel = async (id) => {hotel.value = await getHotel(id);}
loadHotel(props.id)

</script>
<template>
  <div v-if="hotel" class="content">
    <h2>Hello from {{ hotel.name }} page id = {{ hotel.id }}</h2>
    <p>{{ hotel.description }}</p>
  </div>
</template>
