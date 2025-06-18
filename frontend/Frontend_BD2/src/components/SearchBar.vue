<script setup>
import { ref, inject } from 'vue'
import { useFetch } from '@/composables/useFetch'

import VueDatePicker from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'

const API_URL = inject('API_URL')
const searchDetails = inject('search_details')

const emit = defineEmits(['search_complete'])

function constructRequest(query) {
	const url = API_URL + '/get/hotels?' + new URLSearchParams(query);
	return new Request(url, { method: 'GET' });
}

async function search() {
	const parsedDateRange = parseDateRange(formDateRange.value)
	console.log(parsedDateRange)

	const query = {
		city: formCity.value,
		startdate: parsedDateRange[0],
		enddate: parsedDateRange[1],
		guests: formNumPeople.value
	}
	searchDetails.value = query
	const request = constructRequest(query)

	const { data: hotels, execute } = useFetch(request);
	await execute();

	emit('search_complete', hotels.value)
}

function formatDate(date) {
	return date?.toISOString().split('T')[0] // yyyy-mm-dd
}

function parseDateRange(dateRange) {
	return [formatDate(dateRange[0]), formatDate(dateRange[1])]
}


let debugDates = [new Date('2025-07-18'), new Date('2025-07-21')]

const formCity = ref('Warszawa')
const formNumPeople = ref(3)
const formDateRange = ref(debugDates)


</script>

<template>
  <form @submit.prevent="search" id="search">
    <div class="top">
      <input
        type="text"
        v-model="formCity"
        required

        style="width: 80%;"
      >

      <VueDatePicker
        v-model="formDateRange"
        range
        :enable-time-picker="false"
        style="padding-top: 1px; padding-left:2px; padding-right: 2px;"
      />


      <!-- :min-date="new Date()" -->

      <input
        type="number"
        v-model="formNumPeople"
        min="0"
        max="10"
        required

        style="width: 15%;"
      >
      <button type="submit">Search</button>
    </div>
  </form>
</template>


<style>
#search {
	max-width: 60%;
  margin: auto;
  padding: 1%;
}


.top {
  display: flex;
  height: 40px;
  padding: 0.5%;
	background-color: rgb(00, 77, 85);
  border-radius: 5px;
}

.top button {
  width: 10%;
}
</style>


