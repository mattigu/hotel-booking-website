<script setup>
import { ref, inject } from 'vue'
import { useFetch } from '@/composables/useFetch'

import VueDatePicker from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'

const API_URL = inject('API_URL')
const emit = defineEmits(['search_complete'])

function constructRequest(query) {
	const url = API_URL + '/get/hotels?' + new URLSearchParams(query);
	return new Request(url, { method: 'GET' });
}

async function search() {
	const parsedDateRange = parseDateRange(formDateRange.value)
	const query = {
		city: formCity.value,
		startdate: parsedDateRange[0],
		enddate: parsedDateRange[1],
		guests: formNumPeople.value
	}
	const request = constructRequest(query)

	const { data: hotels, execute } = useFetch(request, { immediate: false });
	await execute();

	emit('search_complete', hotels.value)
}

function formatDate(date) {
	return date?.toISOString().split('T')[0] // yyyy-mm-dd
}

function parseDateRange(dateRange) {
	return [formatDate(dateRange[0]), formatDate(dateRange[1])]
}


let debugDates = [new Date('2025-05-24'), new Date('2025-05-24')]

const formCity = ref('Warszawa')
const formNumPeople = ref(3)
const formDateRange = ref(debugDates)


</script>

<template>
  <form @submit.prevent="search" id="search">
    <button type="submit">Search</button>
    <input
      type="text"
      v-model="formCity"
      required
    >
    <VueDatePicker
      v-model="formDateRange"
      range
    /> <!--  :min-date="new Date()" -->

    <input
      type="number"
      v-model="formNumPeople"
      min="0"
      max="10"
      required
    >
  </form>
</template>


<style>
#search {
	max-width: 50%;
}
</style>


