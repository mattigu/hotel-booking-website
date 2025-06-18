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



      <div class="room_info">
        <li>single beds:{{ sigle_beds }}</li>
        <li>double beds:{{ double_beds }}</li>
      </div>
    </div>

    <span>{{ price }}</span>

    <div class="av_check">
      <RouterLink
        class="av_button"
        :id="id"
        :to="`/hotel/${ id }`">
        Check availability
      </RouterLink>
    </div>
  </div>
</template>

<style>
.hotel {
  border-radius: 8px;
  border: 2px solid #ddd;

	width: 60%;
	display: flex;
  margin: auto;
	transition: box-shadow 0.3s ease;
}

.hotel:hover {
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.06);
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
  width: 50%;
}

.room_info {
  padding-left: 15px;
}

.av_check {
  align-self: last baseline;
  padding-bottom: 5%;
}

/* .av_button {
	background-color: rgb(00, 77, 85);
} */

</style>
