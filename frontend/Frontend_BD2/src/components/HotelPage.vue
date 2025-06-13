<script setup>
import vue3StarRatings from "vue3-star-ratings"

const props = defineProps({
	hotel: {
		type: Object,
		required: true
	}
})

const hotel = props.hotel

</script>
<template>
  <!-- must use div here for display: flex, but maybe there's a better solution -->
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
  <hr>
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