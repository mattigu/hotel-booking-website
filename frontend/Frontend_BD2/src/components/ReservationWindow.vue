<script setup>
import { inject } from 'vue';

const searchDetails = inject('search_details') // !!!
const { city, startdate, enddate, guests} = searchDetails.value

</script>

<template>
  <!-- Ja bym tu imo wywalił te inputy i wstawił tylko info jako potwierdzenie. i zrobił to wszystko w Reservation panel już
   żeby się nie męczyć  -->
  date
  <div class="res_data_selector">
    <VueDatePicker
      v-model="formDateRange"
      range
    />
  </div>
  number of people
  <input
    type="number"
    min="0"
    max="10"
    required
  >
</template>

<style>
.res_data_selector {
  width: 30%;
}
</style>
