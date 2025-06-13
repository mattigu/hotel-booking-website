<script setup>
</script>

<template>
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