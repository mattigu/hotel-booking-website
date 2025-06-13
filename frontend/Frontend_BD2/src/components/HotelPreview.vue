<script setup>

import vue3StarRatings from "vue3-star-ratings"

const props = defineProps({
	hotel: {
		type: Object,
		required: true
	}
})

// For now only these since the request is about to change anyway
const { id, name, star_standard } = props.hotel

</script>
<template>
  <div class="hotel">
    <div class="preview_img">
      <img :src="`{{ hotel.photo_url }}`" onerror="this.onerror=null;src=`https://u.profitroom.com/2018-focushotels-pl/thumb/1920x1080/uploads/DJI_0372_MID.jpg`">
    </div>
    <div class="preview_text">
      <h1> {{ name }} HOTEL ID: {{ id }} </h1>
      <vue3StarRatings
        v-model="star_standard"
        :disable-click="true"
      />
      <RouterLink :id="id" :to="`hotel/${ id }`">Check availability</RouterLink>
    </div>
  </div>
</template>

<style>
.hotel {
	background-color: rgb(172, 221, 221);
	width: 70%;
	display: flex;
}

.preview_img img {
	object-fit: cover;
    height: 240px;
    width: 240px;
    border-radius: 5%;
	aspect-ratio: 1;
}

.preview_img {
	margin: 2%;
}

</style>
