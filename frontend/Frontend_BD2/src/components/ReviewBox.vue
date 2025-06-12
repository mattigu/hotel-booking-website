<script setup>

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
    <h1>{{ rating }} stars: {{ username }} </h1>
    <p>{{ review_text }}</p>
  </div>
</template>

<style>

.review {
	outline: 2px solid black;
}

</style>
