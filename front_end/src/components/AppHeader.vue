<template>
  <div class="d-flex ga-3 position-fixed top-0 right-0 justify-center align-center ma-5">
    <v-btn icon="$menu"></v-btn>
    <v-btn icon="mdi-account-circle"></v-btn>
    <v-btn icon="mdi-theme-light-dark" @click="theme.toggle()"></v-btn>
  </div>
</template>

<script setup>
  import { useTheme } from 'vuetify'

  const theme = useTheme()
</script>

<style lang="scss" scoped>
</style>
