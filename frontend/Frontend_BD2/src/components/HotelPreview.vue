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

function test() {
	console.log("something")
	return
}

let text = async () => {text.value = await test();}

</script>
<template>
  <div class="hotel">
    <h1> {{ name }} HOTEL ID: {{ id }} </h1>
    <button @click="text()">aaaa</button>
  </div>
</template>

<style>
.hotel {
	background-color: rgb(172, 221, 221);
	width: 70%;
}

</style>