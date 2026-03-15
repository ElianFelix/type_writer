<template>
  <v-main class="d-flex flex-column app-container position-relative">
    <AppHeader />
    <router-view v-slot="{ Component }">
      <transition>
        <component :is="Component" />
      </transition>
    </router-view>
    <AppFooter />
  </v-main>
</template>

<script setup>
  //
</script>
