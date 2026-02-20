<template>
  <div class="d-flex position-absolute top-30 justify-center align-center ma-auto timer-container">
    <v-progress-circular v-model="model" color="secondary" :size="60" :width="8">
      <template v-slot:default>{{ model }}</template>
    </v-progress-circular>
  </div>
</template>

<script setup>
  const model = defineModel()
</script>

<style lang="scss" scoped>
</style>
