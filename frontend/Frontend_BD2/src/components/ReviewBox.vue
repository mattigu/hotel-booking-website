<script setup>
import vue3StarRatings from "vue3-star-ratings"

const props = defineProps({
	review: {
		type: Object,
		required: true
	}
})

const { review_text, username, rating } = props.review

</script>


<template>
  <div class="review">
    <vue3StarRatings
      v-model="rating"
      :disable-click="true" />
    <h1> {{ username }} </h1>
    <p>{{ review_text }}</p>
  </div>
</template>

<style>

.review {
	outline: 2px solid black;
}

</style>
