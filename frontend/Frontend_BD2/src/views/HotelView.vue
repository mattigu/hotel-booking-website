<script setup>
import { ref, inject } from 'vue'
import ReviewForm from '@/components/ReviewForm.vue'
import ReviewBox from '@/components/ReviewBox.vue'
import HotelPage from '@/components/HotelPage.vue'

const API_URL = inject('API_URL')
// View that displays a detailed Hotel page.

const props = defineProps({
	id: {
		type: String,
		required: true
	}
})

async function fetchHotel(id) {
	const url = API_URL + `/get/hotel/${ id }`
	const request = new Request(url, {
		method: "GET",
	});

	try {
		const response = await fetch(request)
		if (!response.ok) {
			throw new Error(`Failed to get response`);}
		const hotel = await response.json();
		return hotel
	} catch (error) {
		console.error(error.message);
		alert("Failed")
		return;
	}
}

function addReviewToHotel(newReview) {
	hotel.value.reviews.unshift(newReview)
}

const hotel = ref(null)
const loadHotel = async (id) => {hotel.value = await fetchHotel(id)}
loadHotel(props.id)

</script>

<template>
  <div v-if="hotel" class="hotel_view">
    <HotelPage :hotel="hotel" />
  </div>

  <div class="reviews" v-if="hotel">
    <hr>
    <h2>Reviews</h2>
    <ReviewForm @review_posted="addReviewToHotel" />

    <template
      v-for="review in hotel.reviews"
      :key="review.username"
    >
      <!-- Use a better key for this loop later -->

      <ReviewBox :review="review" />
    </template>
  </div>
</template>
