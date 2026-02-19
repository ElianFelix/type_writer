<template>
  <div class="d-flex flex-column flex-grow-1 align-self-center justify-center align-center mx-8 area-container">
    <v-sheet class="px-10 py-4 fill-height font-mono">
<!--      <span class="cursor">
        <v-tooltip activator="parent" location="top" model-value text="Start Typing" @update:model-value="true" />
      I</span>t is a long established fact that a reader will be distracted by the readable content of a page when
      looking at its layout. The point of  using Lorem Ipsum is that it has a more-or-less normal distribution of
      letters, as opposed to using 'Content here, content here', making it  look like readable English.
      Many desktop publishing packages and web  page editors now use Lorem Ipsum as their default model text, and a
      search for 'lorem ipsum' will uncover many web sites still in their  infancy. Various versions have evolved over the years,
      sometimes by  accident, sometimes on purpose (injected humour and the like). -->
      <span v-for="elem, index in computedText" :key="index" :class="elem.status">{{ elem.letter }}</span>
    </v-sheet>
    <div class="d-flex ga-4 py-5">
      <v-btn @click="Restart">Retry</v-btn>
      <v-btn>Done</v-btn>
    </div>
  </div>
</template>

<script setup>
  import { onMounted, ref } from 'vue'

  function ProcessInputText(text) {
    const output = text.map((l, idx) => { return {letter: l, status: idx == 0 ? 'cursor' : ''} })
    return output
  }

  function HandleKeyPress(e) {
    console.log(`key ${e.key} was pressed; also cursor is at ${currentIdx}`)
    if (currentIdx >= computedText.value.length) return
    computedText.value[currentIdx].status = computedText.value[currentIdx].letter == e.key ? 'right' : 'wrong'
    currentIdx++
    if (currentIdx >= computedText.value.length) return
    computedText.value[currentIdx].status = 'cursor'
  }

  function Restart() {
    computedText.value = ProcessInputText(SampleText.split(''))
    currentIdx = 0
  }

  const SampleText = "It is a long established fact that a reader will be distracted by the readable content of a page when " +
    "looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of " +
    "letters, as opposed to using 'Content here, content here', making it look like readable English. " +
    "Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a " +
    "search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, " +
    "sometimes by accident, sometimes on purpose (injected humour and the like)."

  const computedText = ref(ProcessInputText(SampleText.split('')))

  let currentIdx = 0

  onMounted(() => {
    window.addEventListener('keypress', HandleKeyPress)
  })


</script>

<style lang="scss" scoped>
</style>
