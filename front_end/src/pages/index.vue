<template>
  <div class="fill-height board-container">
    <Timer v-model="timerSeconds" v-model:test-time="testTime" />
    <div class="d-flex flex-column fill-height ga-3 align-center">
      <Stats v-model="stats" />
      <Area v-model:stats="stats" v-model:test-time="testTime" v-model:timer-seconds="timerSeconds" />
    </div>
  </div>
</template>

<script setup>
  import { ref } from 'vue';

  const DEFAULT_TIME = 60

  let testTime = ref(DEFAULT_TIME)
  let timerSeconds = ref(DEFAULT_TIME)
  let stats = ref(null)
</script>
