<script setup>
import { ref, inject, computed, watch} from 'vue';
import AddonSelector from './AddonSelector.vue';
import ConfigurationSelector from './ConfigurationSelector.vue';
import { useFetch } from '@/composables/useFetch';

// I interpreted the prices as daily prices so they are multiplied by the number of days

const API_URL = inject('API_URL')
const searchDetails = inject('search_details') // INFO FROM SEARCH BAR

const { city, startdate, enddate, guests} = searchDetails.value


const props = defineProps({
	hotel: {
		type: Object,
		required: true
	}
})

const { id: hotel_id, addons} = props.hotel

const configurations = ref([])

async function fetchConfigurations() {
	const query = {
		hotel_id: hotel_id,
		guests: 2
	}
	const url = API_URL + '/get/configuration?'+ new URLSearchParams(query)
	const request = new Request(url)

	const { data, execute } = useFetch(request)
	await execute()
	configurations.value = data.value || []
}

fetchConfigurations()

function daysBetween(start, end) {
	const startDate = new Date(start);
	const endDate = new Date(end);
	return Math.floor((endDate - startDate) / (1000 * 60 * 60 * 24));
}

const selectedAddons = ref([])
const selectedConfiguration = ref({})

const addonsPrice = ref(0)
const configurationPrice = ref(0)

watch(selectedAddons, () => {
	addonsPrice.value = selectedAddons.value.reduce((total, item) => total + (item.price || 0), 0);
})

watch(selectedConfiguration, () => {
	configurationPrice.value = (selectedConfiguration.value.price || 0);
})

const totalPrice = computed(() => {
	console.log(daysBetween(startdate, enddate))
	return (addonsPrice.value + configurationPrice.value) * daysBetween(startdate, enddate)
})

</script>

<template>
  <h1>Reservation</h1>
  <span>Current total: {{ totalPrice }}</span>
  <AddonSelector
    :addons="addons"
    v-model="selectedAddons" />

  <ConfigurationSelector
    :configurations="configurations"
    v-model="selectedConfiguration" />
</template>



<style>

</style>
