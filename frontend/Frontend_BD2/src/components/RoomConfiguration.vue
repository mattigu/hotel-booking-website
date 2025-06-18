<script setup>

const props = defineProps({
	configuration: {
		type: Object,
		required: true
	}
})

const { single_beds, double_beds, price } = props.configuration

</script>


<template>
  <span> {{ single_beds }} single beds, {{ double_beds }} double beds, {{ price }}$</span>
</template>
