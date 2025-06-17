<script setup>

import vue3StarRatings from "vue3-star-ratings"

const props = defineProps({
	hotel: {
		type: Object,
		required: true
	}
})

// For now only these since the request is about to change anyway
const { id, name, star_standard, sigle_beds, double_beds, photo_url, price } = props.hotel

</script>
<template>
  <div class="hotel">
    <div class="preview_img">
      <img :src="photo_url" onerror="this.onerror=null;this.src=`https://u.profitroom.com/2018-focushotels-pl/thumb/1920x1080/uploads/DJI_0372_MID.jpg`">
    </div>
    <div class="preview_info">
      <span>
        <vue3StarRatings
          v-model="star_standard"
          :disable-click="true"
        />
      </span>

      <h1> {{ name }} </h1>

      <RouterLink :id="id" :to="`/hotel/${ id }`">Check availability</RouterLink>
    </div>

    <div class="best_offer_info">
      <p>Best offer</p>
      {{ price }}
      <p>single beds:{{ sigle_beds }}</p>
      <p>double beds:{{ double_beds }}</p>
    </div>
  </div>
</template>

<style>
.hotel {
  border-radius: 8px;
  border: 2px solid #ddd;

	width: 70%;
	display: flex;
  margin: auto;
}

.preview_img img {
	object-fit: cover;
  height: 240px;
  width: 240px;
  border-radius: 5%;
	aspect-ratio: 1;
  margin: 2%;
}

.preview_info {
	margin: 2%;
}

</style>
