<template>
  <div class="fill-height board-container">
    <Timer v-model="timerSeconds" />
    <div class="d-flex flex-column fill-height ga-3 align-center">
      <Stats v-model="stats" />
      <Area v-model:stats="stats" v-model:timer-seconds="timerSeconds" />
    </div>
  </div>
</template>

<script setup>
  import { ref } from 'vue';

  //
  let timerSeconds = ref(60)
  let stats = ref(null)
</script>
