<script setup>
import RoomConfiguration from './RoomConfiguration.vue';

defineProps({
	configurations: {
		type: Array,
		required: true
	}
})

const selected = defineModel(Object)

</script>
<!-- KLUCZ NIE UNIQUE -->
<template>
  <div>
    <template v-for="conf in configurations" :key="conf.price">
      <input
        type="radio"
        v-model="selected"
        name="conf"
        :value="conf">
      <label for="conf"> <RoomConfiguration :configuration="conf" /></label>
    </template>
  </div>
</template>
