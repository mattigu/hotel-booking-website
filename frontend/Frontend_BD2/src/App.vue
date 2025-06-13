<script setup>

</script>

<template>
  <nav>
    <li><RouterLink to="/">Go to Home</RouterLink></li>
  </nav>
  <main>
    <!-- Caches hotel list when navigating to a hotel.
	 Unfortunalely this is the only way to do this with script setup
	 https://stackoverflow.com/questions/65619181/how-to-make-certain-component-keep-alive-with-router-view-in-vue-3 -->
    <router-view v-slot="{ Component }">
      <keep-alive include="HotelBrowseView">
        <component :is="Component" />
      </keep-alive>
    </router-view>
  </main>
</template>
