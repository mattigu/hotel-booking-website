<script setup>
import vue3StarRatings from "vue3-star-ratings"

const props = defineProps({
	review: {
		type: Object,
		required: true
	}
})

const { review_text, username, rating } = props.review

const date = "11/12/2025"

</script>


<template>
  <div class="review">
    <div class="username_date">
      <span class="review_username">{{ username }}</span>
      <span class="review_date">{{ date }}</span>
    </div>

    <vue3StarRatings
      v-model="rating"
      :disable-click="true" />

    <p class="review_text">
      {{ review_text }}
    </p>
  </div>
</template>

<style>

.review {
	padding: 12px;
	border: 1px solid #ddd;
	border-radius: 8px;
	background: #fff;
	word-wrap: break-word;
	width: 100%;            /* make sure it fills grid cell */
	box-sizing: border-box;
	transition: box-shadow 0.3s ease;
}

.username_date {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 6px;
}

.review_date {
	font-size: 0.875rem;
	color: #666;
}

.review_username {
	font-weight: bold;
}

.review_text {
	font-size: 0.95rem;
	color: #333;
	margin: 0;
	white-space: pre-wrap;
	overflow-wrap: break-word;
	word-break: break-word;
}

.review:hover {
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.06);
}

</style>
