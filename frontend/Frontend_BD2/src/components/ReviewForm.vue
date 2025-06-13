<script setup>

import { ref, inject } from 'vue'
import { useFetch } from '@/composables/useFetch'
import vue3StarRatings from "vue3-star-ratings"

const API_URL = inject('API_URL')

const emit = defineEmits(['review_posted'])

async function postReview() {
	const review = {
		hotel_id: 		1,
		username: 		formUsername.value,
		review_text: 	formText.value,
		rating: 		formStars.value
	}
	const url = API_URL + '/new/review'
	const request = new Request(url, {
		method: "POST",
		body: JSON.stringify(review),
		headers: {
			"Content-Type": "application/json"
		}})
	const { error, execute } = useFetch(request)
	await execute()

	if (error.value) {
		alert(`Error: ${error.value}`)
	} else {
		emit('review_posted', review)
	}
}

function handleRating(value) { // Rounds to nearest integer
	formStars.value = Math.floor(value + 1)
}

const formText = ref()
const formUsername = ref()
const formStars = ref(4)

</script>

<template>
  <form @submit.prevent="postReview" id="review_form">
    <vue3StarRatings
      :model-value="formStars"
      @update:model-value="handleRating"
    />
    <input
      type="text"
      name="username"
      v-model="formUsername"
      required>
    <label for="username">Username</label>

    <textarea
      name="review_text"
      v-model="formText" />

    <button type="submit">Post</button>
  </form>
</template>

<style>

#review_form {
	outline: 2px solid black;
}
</style>
