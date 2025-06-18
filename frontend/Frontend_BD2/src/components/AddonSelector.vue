<script setup>

defineProps({
	addons: {
		type: Array,
		required: true
	}
})

const selected = defineModel({ type: Array, default: () => [] })

</script>

<template>
  <div>
    <template
      v-for="addon in addons"
      :key="addon.id">
      <input
        type="checkbox"
        v-model="selected"
        :value="addon"
        name="addon"
      >
      <label for="addon"> {{ addon.name }} - ${{ addon.price }} </label>
    </template>
  </div>
</template>
