<template>
  <div class="d-flex position-absolute top-30 justify-center align-center ma-auto timer-container">
    <v-progress-circular v-model="timerProgress" color="secondary" :size="60" :width="8">
      <template v-slot:default>{{ model }}</template>
    </v-progress-circular>
  </div>
</template>

<script setup>
  import { computed } from 'vue';

  const model = defineModel()
  const testTime = defineModel('testTime')
  const timerProgress = computed(() => model.value / testTime.value * 100)
</script>

<style lang="scss" scoped>
</style>
