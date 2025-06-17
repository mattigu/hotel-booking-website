<script setup>
import vue3StarRatings from "vue3-star-ratings"
import ReviewForm from '@/components/ReviewForm.vue'
import ReviewBox from '@/components/ReviewBox.vue'

const props = defineProps({
	hotel: {
		type: Object,
		required: true
	}
})

function addReviewToHotel(newReview) {
	reviews.unshift(newReview)
}

const {name, description, star_standard, address, amenities, photo_url, reviews} = props.hotel
const {street, city, house_number, country} = address
</script>
<template>
  <!-- must use div here for display: flex, but maybe there's a better solution -->
  <div class="hotel_top">
    <div class="hotel_title">
      <h2>{{ name }}, {{ city }}</h2>
      <p>{{ street }} {{ house_number }}, {{ city }}, {{ country }}</p>
    </div>

    <div class="stars">
      <vue3StarRatings
        v-model="star_standard"
        :disable-click="true" />
    </div>
  </div>

  <div class="hotel_img">
    <img :src="photo_url" onerror="this.onerror=null;src=`https://u.profitroom.com/2018-focushotels-pl/thumb/1920x1080/uploads/DJI_0372_MID.jpg`">
  </div>

  <hr>
  <h2>About this hotel</h2>
  <p>{{ description }}</p>
  <h2>Amenities</h2>

  <template v-for="amenity in amenities" :key="amenity.id">
    <h3>{{ amenity.name }}</h3>
    <p>{{ amenity.description }}</p>
  </template>

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
