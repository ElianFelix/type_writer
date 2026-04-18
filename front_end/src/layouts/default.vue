<template>
  <v-main class="d-flex flex-column app-container">
    <AppHeader />
    <router-view v-slot="{ Component }">
      <div style="margin-top: 120px; padding-bottom: 40px;">
        <transition>
          <component :is="Component" />
        </transition>
      </div>
    </router-view>
    <AppFooter />
  </v-main>
</template>

<script setup>
  //
</script>
