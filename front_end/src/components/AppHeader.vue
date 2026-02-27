<template>
  <div class="d-flex ga-3 position-absolute top-0 right-0 justify-center align-center ma-5 mr-0">
    <v-sheet class="d-flex pa-2 justify-space-between"
             :class="{ 'bg-transparent': !openMenu }"
             height="200"
             rounded
             :width="openMenu ? 250 : 180"
    >
      <v-btn icon="mdi-theme-light-dark" @click="theme.toggle()"></v-btn>
      <v-btn icon="mdi-account-circle"></v-btn>
      <v-btn icon="$menu" @click="() => {openMenu = !openMenu}"></v-btn>
    </v-sheet>
  </div>
</template>

<script setup>
  import { ref } from 'vue';
  import { useTheme } from 'vuetify'

  const theme = useTheme()
  const openMenu = ref(false)
</script>

<style lang="scss" scoped>
</style>
