<script setup>
import { ref, inject } from 'vue'

const API_URL = inject('API_URL')

async function testBackendGet() {
	const url = API_URL + "/hotels/getall"
	try {
		const response = await fetch(url)
		if (!response.ok) {
			throw new Error(`Response status: ${response.status}`)
		}
		const message = await response.json()
		return message
	} catch (error) {
		console.error(error.message)
	}
}

// async function testBackendPost() {
// 	const url = API_URL + "/test/postTest"
// 	const request = new Request(url, {
// 		method: "POST",
// 		body: JSON.stringify({
// 			"id": 123,
// 			"name": "jajo"
// 		})
// 	})
// 	try {
// 		const response = await fetch(request)
// 		if (!response.ok) {
// 			throw new Error(`Response status: ${response.status}`)
// 		}
// 		const message = await response.text()
// 		return message
// 	} catch (error) {
// 		console.error(error.message)
// 	}
// }


var getMessage = ref()
const testGet = async () => getMessage.value = await testBackendGet()
</script>

<template>
  <button @click="testGet() ">
    Get
  </button>
  <h1>Message from db: {{ getMessage }}</h1>
</template>
