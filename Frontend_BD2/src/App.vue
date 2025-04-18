<script setup>
import { ref, inject } from 'vue'

const API_URL = inject('API_URL')

async function testBackendGet() {
	const url = API_URL + "/test/getTest"
	try {
		const response = await fetch(url)
		if (!response.ok) {
			throw new Error(`Response status: ${response.status}`)
		}
		const message = await response.text()
		console.log(response)
		return message
	} catch (error) {
		console.error(error.message)
	}
}

async function testBackendPost() {
	const url = API_URL + "/test/postTest"
	const request = new Request(url, {
        method: "POST",
		body: JSON.stringify({
			"id" : 123,
			"name" : "jajo"
		})
	})
	try {
		const response = await fetch(request)
		if (!response.ok) {
			throw new Error(`Response status: ${response.status}`)
		}
		const message = await response.text()
		return message
	} catch (error) {
		console.error(error.message)
	}
}


var getMessage = ref()
var postMessage = ref()
const testGet = async () => getMessage.value = await testBackendGet()
const testPost = async () => postMessage.value = await testBackendPost()
</script>

<template>

<button @click="testGet()">Get</button>
<button @click="testPost()">Post</button>
<!-- <button @click="testBackend()">Fetch</button> -->
<h1>Message from db: {{ getMessage }}</h1>
<h1>Post response db: {{ postMessage }}</h1>

</template>

