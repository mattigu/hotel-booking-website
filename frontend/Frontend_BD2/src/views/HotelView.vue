<script setup>
import { ref, inject } from 'vue'
import ReviewForm from '@/components/ReviewForm.vue'
import ReviewBox from '@/components/ReviewBox.vue'
import ReservationWindow from '../components/ReservationWindow.vue'
import vue3StarRatings from "vue3-star-ratings"

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

let resWind = ref(false)

</script>

<template>
  <div v-if="hotel" class="hotel_view">
    <div class="hotel_top">
      <div class="hotel_title">
        <h2>{{ hotel.name }}, {{ hotel.address['city'] }}</h2>
        <p>{{ hotel.address['street'] }} {{ hotel.address["house_number"] }}, {{ hotel.address["city"] }}, {{ hotel.address["country"] }}</p>
      </div>
      <div class="stars">
        <vue3StarRatings
          v-model="hotel.star_standard"
          :disable-click="true" />
      </div>
    </div>
    <div class="hotel_img">
      <img :src="`{{ hotel.photo_url }}`" onerror="this.onerror=null;src=`https://u.profitroom.com/2018-focushotels-pl/thumb/1920x1080/uploads/DJI_0372_MID.jpg`">
    </div>
    <hr style="height: 10px; color: rgb(00, 77, 85); background-color: rgb(00, 77, 85); border-radius: 5%;">
    <div class="hotel_description">
      <h2>About this hotel</h2>
      <p>{{ hotel.description }}</p>
    </div>
    <div class="amenities_list">
      <h2>Amenities</h2>
      <template v-for="amenity in hotel.amenities" :key="amenity.id">
        <h3>{{ amenity.name }}</h3>
        <p>- {{ amenity.description }}</p>
      </template>
    </div>
    <!-- very temporary will change later-->
    <div class="res_but">
      <button @click="resWind=true">
        Reserve
      </button>
      <button @click="resWind=false">
        colapse
      </button>
      <div v-if="resWind" class="test">
        <ReservationWindow />
      </div>
      <hr style="height: 10px; color: rgb(00, 77, 85); background-color: rgb(00, 77, 85); border-radius: 5%;">
    </div>
    <h2>Reviews</h2>
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
  </div>
</template>

<style>

.hotel_img img {
	object-fit: cover;
    height: 480px;
    width: 480px;
    border-radius: 5%;
	aspect-ratio: 1;
}

.hotel_top {
	display: flex;
}

.stars {
	padding-left: 5%;
	padding-top: 1.75%;
}
</style>


