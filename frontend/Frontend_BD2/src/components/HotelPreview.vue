<script setup>
const props = defineProps({
	name: {
		type: String,
		required: true
	},
	id: {
		type: Number,
		required: true
	}
})

</script>
<template>
  <div class="hotel">
    <h1> {{ props.name }} HOTEL ID: {{ props.id }} </h1>
    <RouterLink :id="id" :to="`hotel/${ id }`">Check availability</RouterLink>
  </div>
</template>

<style>
.hotel {
	background-color: rgb(172, 221, 221);
	width: 70%;
}

</style>
