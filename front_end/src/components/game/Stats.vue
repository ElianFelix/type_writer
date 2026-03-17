<template>
  <div class="d-flex align-self-center pa-auto justify-center align-center stats-container"
       :class="{ 'hidden': false }"
  >
    <v-sheet class="d-flex ga-12 w-100 px-4 py-8 justify-center align-center inset-shadow">
      <div>{{ model[1].wpm ?? '' }} WPM</div>
      <div>{{ model[1].lpm ?? '' }} LPM</div>
      <div>time: {{ model[0] ?? '' }} secs</div>
      <div>letters: {{ model[1].letters ?? '' }}</div>
      <div>words: {{ model[1].words ?? '' }}</div>
      <div>errors: {{ model[1].errors ?? '' }}</div>
      <div>corrected</div>
    </v-sheet>
  </div>
</template>

<script setup>
  const model = defineModel()
</script>

<style lang="scss" scoped>
</style>
