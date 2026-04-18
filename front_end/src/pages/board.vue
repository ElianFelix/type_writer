<template>
  <div class="board-container">
    <Timer v-model="appStore.timerSeconds" v-model:test-time="appStore.testTime" />
    <div class="d-flex flex-column fill-height ga-3 align-stretch">
      <Stats v-if="appStore.completed" v-model="appStore.stats" />
      <Area
        v-model:font-size="appStore.fontSize"
        v-model:stats="appStore.stats"
        v-model:timer-seconds="appStore.timerSeconds"
      />
    </div>
  </div>
</template>

<script setup>
  import { useAppStore } from '@/stores/app';

  const appStore = useAppStore()
</script>
