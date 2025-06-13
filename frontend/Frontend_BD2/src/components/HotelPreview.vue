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
    <h1> {{ name }} HOTEL ID: {{ id }} </h1>
    <vue3StarRatings
      v-model="star_standard"
      :disable-click="true"
    />
    <RouterLink :id="id" :to="`hotel/${ id }`">Check availability</RouterLink>
  </div>
</template>

<style>
.hotel {
	background-color: rgb(172, 221, 221);
	width: 70%;
}

</style>
