<template>
  <div class="d-flex position-absolute top-30 justify-center align-center ma-auto timer-container">
    <v-progress-circular color="secondary" model-value="50" :size="60" :width="8">
      <template v-slot:default>00:30</template>
    </v-progress-circular>
  </div>
</template>

<script setup>
</script>

<style lang="scss" scoped>
</style>
