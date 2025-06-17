<script setup>
import { inject, ref } from 'vue'
import { useFetch } from '@/composables/useFetch'

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

function addReviewToHotel(newReview) {
	hotel.value.reviews.unshift(newReview)
}

async function loadHotel() {
	const request = new Request(`${API_URL}/get/hotel/${props.id}`, { method: 'GET' })
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
