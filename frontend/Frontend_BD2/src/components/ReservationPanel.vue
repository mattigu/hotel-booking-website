<script setup>
import { ref, inject, computed, watch} from 'vue';
import AddonSelector from './AddonSelector.vue';
import ConfigurationSelector from './ConfigurationSelector.vue';
import { useFetch } from '@/composables/useFetch';

// I interpreted the prices as daily prices so they are multiplied by the number of days

const API_URL = inject('API_URL')
const searchDetails = inject('search_details') // INFO FROM SEARCH BAR

const { city, startdate, enddate, guests} = searchDetails.value
const emit = defineEmits(['reservation_confirmed'])


const props = defineProps({
	hotel: {
		type: Object,
		required: true
	}
})

const { id: hotel_id, addons} = props.hotel

const configurations = ref([])

async function reserveRoom() {
	// console.log(typeof startdate)
	if (startdate) {
		const reservation = {
			room_id: selectedConfiguration.value["id"],
			hotel_id: hotel_id,
			start_date: startdate,
			end_date: enddate,
			addons: selectedAddons.value.map(addon => addon.id),
			customer: {
				name: formName.value,
				surname: formSurname.value,
				phone_number: formPhone_number.value
			},
			payment_info: {
				payment_type: "przelew",
				payment_data: formPayment_data.value,
				amount: totalPrice.value
			}
		}

		const url = API_URL + '/reserve/room'
		const request = new Request(url, {
			method: "POST",
			body: JSON.stringify(reservation),
			headers: {
				"Content-Type": "application/json"
			}})
		const { error, execute } = useFetch(request)
		await execute()

		if (error.value) {
			alert(`Error: ${error.value}`)
		} else {
			emit('reservation_confirmed', reservation)
			console.log('res suc')
		}
	}
}

async function fetchConfigurations() {
	// if (typeof startdate == typeof "") {
	if (startdate) {
		const query = {
			hotel_id: hotel_id,
			guests: 2,
			start_date: startdate,
			end_date: enddate
		}
		const url = API_URL + '/get/configuration?'+ new URLSearchParams(query)
		const request = new Request(url)

		const { data, execute } = useFetch(request)
		await execute()
		configurations.value = data.value || []

	}

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


function tryOpenModal() {
	if (!selectedConfiguration.value.id) {
		alert('Select a room configuration before reserving.');
		return;
	}

	showModal.value = true;
}

const showModal = ref(false)
const formName = ref('')
const formSurname = ref('')
const formPhone_number = ref('')
const formPayment_data = ref('')

</script>

<template>
  <div class="res">
    <h1>Reservation</h1>
    <span>Current total: {{ totalPrice }}$</span>
    <AddonSelector
      :addons="addons"
      v-model="selectedAddons" />

    <ConfigurationSelector
      :configurations="configurations"
      v-model="selectedConfiguration" />
    <br>

    <!-- Button to open the modal -->
    <!-- <button
      @click="showModal = true"
      :disabled="!selectedConfiguration.id">
      Reserve
    </button> -->

    <button @click="tryOpenModal">
      Reserve
    </button>

    <!-- Modal -->
    <div v-if="showModal">
      <h2>Reservation ditails</h2>
      <br>
      single beds: {{ selectedConfiguration.single_beds }}
      <br>
      double beds: {{ selectedConfiguration.double_beds }}
      <br>
      <div v-if="selectedAddons.length">
        Choosen additional amenities
        <li
          v-for="addon in selectedAddons"
          :key="addon.id">
          {{ addon.name }} - {{ addon.price }} zł/dzień
        </li>
      </div>
      Total price
      {{ totalPrice }}$

      <h2>Enter your details</h2>
      <form @submit.prevent="reserveRoom">
        <label>
          First name
        </label>
        <input
          type="text"
          name="name"
          v-model="formName"
          maxlength="25"
          required
        >
        <label>
          Last name
        </label>
        <input
          type="text"
          name="surname"
          v-model="formSurname"
          maxlength="25"
          required
        >
        <label>
          Phone number
        </label>
        <input
          type="text"
          name="phone_number"
          v-model="formPhone_number"
          maxlength="25"
          required
        >
        <h3>Payment info</h3>

        <label>
          Bank account number
        </label>
        <input
          type="text"
          name="payment_data"
          v-model="formPayment_data"
          maxlength="25"
          required
        >

        <div style="padding-top: 10px;">
          <span style="padding: 5px;">
            <button type="submit">Confirm</button>
          </span>
          <span>
            <button type="button" @click="showModal = false">Cancel</button>
          </span>
        </div>
      </form>
    </div>
  </div>
</template>



<style>
.res button {
	padding: 8px 12px;
	font-size: 1rem;
	border: none;
	border-radius: 4px;
	background-color: rgb(00, 77, 85);
	color: white;
	cursor: pointer;
}

.res button:hover {
  background-color: rgb(0, 54, 60);
}

.res button {
	transition: background-color 0.1s ease;
}
</style>
