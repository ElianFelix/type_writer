<template>
  <div class="fill-height board-container">
    <Timer v-model="appStore.timerSeconds" v-model:test-time="appStore.testTime" />
    <div class="d-flex flex-column fill-height ga-3 align-center">
      <Stats v-model="appStore.stats" />
      <Area
        v-model:stats="appStore.stats"
        v-model:test-time="appStore.testTime"
        v-model:timer-seconds="appStore.timerSeconds"
        v-model:font-size="appStore.fontSize"
      />
    </div>
  </div>
</template>

<script setup>
  import { useAppStore } from '@/stores/app';

  const appStore = useAppStore()
</script>
