<script setup>
import { ref, inject } from 'vue'

const API_URL = inject('API_URL')


async function testBackend() {
	const url = API_URL + "/"
	try {
		const response = await fetch(url)
		if (!response.ok) {
			throw new Error(`Response status: ${response.status}`)
		}
		const message = await response.text()
		return message
	} catch (error) {
		console.error(error.message)
	}
}


var dbTest = ref()
const loadTest = async () => dbTest.value = await testBackend()
</script>

<template>

<button @click="loadTest()">Fetch</button>
<!-- <button @click="testBackend()">Fetch</button> -->
<h1>Message from dfddb: {{ dbTest }}</h1>

</template>

