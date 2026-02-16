<template>
  <div class="d-flex flex-column flex-grow-1 align-self-center justify-center align-center mx-8 area-container">
    <v-sheet class="px-10 py-4 fill-height font-mono">
      <span class="cursor">
        <v-tooltip class="red" model-value text="Start Typing" @update:model-value="true" activator="parent" location="top"/>
      I</span>t is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of  using Lorem Ipsum is that it has a more-or-less normal distribution of  letters, as opposed to using 'Content here, content here', making it  look like readable English. Many desktop publishing packages and web  page editors now use Lorem Ipsum as their default model text, and a  search for 'lorem ipsum' will uncover many web sites still in their  infancy. Various versions have evolved over the years, sometimes by  accident, sometimes on purpose (injected humour and the like).
    </v-sheet>
    <div class="d-flex ga-4 py-5">
      <v-btn>Retry</v-btn>
      <v-btn>Done</v-btn>
    </div>
  </div>
</template>

<script setup>
</script>

<style lang="scss" scoped>
</style>
