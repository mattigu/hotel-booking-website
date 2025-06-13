<script setup>
import { ref, inject } from 'vue'
import ReviewForm from '@/components/ReviewForm.vue'
import ReviewBox from '@/components/ReviewBox.vue'

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
  <div v-if="hotel" class="content">
    <div class="hotel_top">
      <p>name: {{ hotel.name }} page id = {{ hotel.id }}</p>
      <!-- will change this to img later -->
      <p>Hotel star standard: {{ hotel.star_standard }}</p>
      <h2>{{ hotel.name }}, {{ hotel.description }}, {{ hotel.address['city'] }}</h2>
      <p>{{ hotel.address['street'] }} {{ hotel.address["house_number"] }}, {{ hotel.address["city"] }}, {{ hotel.address["country"] }}</p>
    </div>
    <div class="hotel_img">
      <img :src="`{{ hotel.photo_url }}`" onerror="this.onerror=null;src=`https://u.profitroom.com/2018-focushotels-pl/thumb/1920x1080/uploads/DJI_0372_MID.jpg`">
    </div>
    <div class="amenities_list">
      <h2>Amenities</h2>
      <template v-for="amenity in hotel.amenities" :key="amenity.id">
        <p>{{ amenity.name }}</p>
        <p>- {{ amenity.description }}</p>
      </template>
    </div>
  </div>

  <ReviewForm @review_posted="addReviewToHotel" />

  <div class="reviews" v-if="hotel">
    <template
      v-for="review in hotel.reviews"
      :key="review.username"
    >
      <!-- Use a better key for this loop later -->

      <ReviewBox :review="review" />
    </template>
  </div>
</template>

<style>
.hotel_img img{
	object-fit: cover;
    height: 240px;
    width: 240px;
    border-radius: 5%;
	aspect-ratio: 1;
}
</style>
