<script setup>

defineProps({
	id: String
})

</script>

<template>
  <p>Specific hotel view with id= {{ id }}
    Here a hotel page with the id can be loaded after the user clicks it when searching</p>


  <a href="https://router.vuejs.org/guide/advanced/data-fetching.html">Exact example from docs what to do</a>

</template>
