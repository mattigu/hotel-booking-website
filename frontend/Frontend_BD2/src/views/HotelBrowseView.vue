<script setup>
import { ref } from 'vue'
import HotelPreview from '@/components/HotelPreview.vue';

// hotels/getall
async function getHotels() {
	const request = new Request("http://localhost:3000/hotels/getall", {
		method: "GET",
	});

	try {
		const response = await fetch(request)
		if (!response.ok) {
			throw new Error(`Failed to get response`);}
		let hotels = await response.json();
		return hotels
	} catch (error) {
		console.error(error.message);
		alert("Failed")
		return;
	}
}

let allHotels = ref([]);
const loadHotels = async () => {allHotels.value = await getHotels();}
loadHotels();

</script>

<template>
  <button @click="loadHotels() "> Find hotels </button>
  <template
    v-for="hotel in allHotels"
    :key="hotel.id"
  >
    <HotelPreview
      :name="hotel.name"
      :id="hotel.id"
    />
  </template>
</template>
