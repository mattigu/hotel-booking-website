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
      placeholder="Your name"
      maxlength="25"
      required>

    <textarea
      name="review_text"
      v-model="formText"
      placeholder="Write your review..."
      rows="5" />

    <button type="submit">Post</button>
  </form>
</template>

<style>

#review_form {
	display: flex;
	flex-direction: column;
	margin-bottom: 20px;
	gap: 10px;
	padding: 12px;
	border: 1px solid #ccc;
	border-radius: 8px;
	background: #fff;
	max-width: 600px;
	width: 100%;
	box-sizing: border-box;
}

#review_form input {
	max-width: 30%;
	padding: 4px;
}

#review_form textarea {
	padding: 4px;
}

#review_form button {
	padding: 8px 12px;
	font-size: 1rem;
	border: none;
	border-radius: 4px;
	background-color: rgb(00, 77, 85);
	color: white;
	cursor: pointer;
}

#review_form button:hover {
  background-color: rgb(0, 54, 60);
}

#review_form button {
	transition: background-color 0.1s ease;
}
</style>
