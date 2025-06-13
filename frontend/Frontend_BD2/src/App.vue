<script setup>

</script>

<template>
  <div class="background">
    <nav>
      <li><RouterLink to="/">Go to Home</RouterLink></li>
    </nav>
    <div class="main_view">
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
    </div>
  </div>
</template>

<style>
.background {
	background-color: rgb(00, 77, 85);
}

.main_view {
	background-color: white;
	width: 70%;
  height: 100vmax;
	margin: auto;
	padding-left: 5%;
	padding-right: 5%;
  /* position:absolute; */
}

</style>