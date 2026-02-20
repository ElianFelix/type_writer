<template>
  <div class="d-flex align-self-center pa-auto justify-center align-center stats-container" style="visibility: visible;">
    <v-sheet class="d-flex ga-12 w-100 py-8 justify-center align-center inset-shadow">
      <div>99 WPM</div>
      <div>99 LPM</div>
      <div>time: {{ model?.time ?? '' }}</div>
      <div>letters: {{ model?.letters ?? '' }}</div>
      <div>words: {{ model?.words ?? '' }}</div>
      <div>errors: {{ model?.errors ?? '' }}</div>
      <div>corrected</div>
    </v-sheet>
  </div>
</template>

<script setup>
  const model = defineModel()
</script>

<style lang="scss" scoped>
</style>
