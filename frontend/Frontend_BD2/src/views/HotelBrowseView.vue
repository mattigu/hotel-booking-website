<script>
export default { name: 'HotelBrowseView' }
</script>

<script setup>
import { ref } from 'vue'
import HotelPreview from '@/components/HotelPreview.vue';
import SearchBar from '@/components/SearchBar.vue';


const hotelList = ref([])

function handleSearch(hotels) {
	hotelList.value = hotels
}

</script>

<template>
  <SearchBar @search_complete="handleSearch" />

  <template
    v-for="hotel in hotelList"
    :key="hotel.id"
  >
    <HotelPreview
      :hotel="hotel"
    />
  </template>
</template>
