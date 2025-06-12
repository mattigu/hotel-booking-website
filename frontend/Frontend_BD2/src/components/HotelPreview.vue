<script setup>
defineProps({
	name: {
		type: String,
		required: true
	},
	id: {
		type: String,
		required: true
	}
})


</script>
<template>
  <div class="hotel">
    <h1> {{ name }} HOTEL ID: {{ id }} </h1>
    <li><RouterLink :id="id" :to="`hotel/${ id }`">Check availability</RouterLink></li>
  </div>
</template>

<style>
.hotel {
	background-color: rgb(172, 221, 221);
	width: 70%;
}

</style>
