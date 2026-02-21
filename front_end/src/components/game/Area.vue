<template>
  <div class="d-flex flex-column flex-grow-1 align-self-center justify-center align-center mx-8 area-container">
    <v-sheet class="px-10 py-4 fill-height font-mono">
      <span v-for="elem, index in computedText" :key="index" :class="elem.status">{{ elem.letter }}</span>
    </v-sheet>
    <div class="d-flex ga-4 py-5">
      <v-btn @click="Restart">Retry</v-btn>
      <v-btn>Done</v-btn>
    </div>
  </div>
</template>

<script setup>
  import { onMounted, onUnmounted, ref } from 'vue'

  function ProcessInputText(text) {
    const output = text.map((l, idx) => { return { letter: l, status: idx == 0 ? 'cursor' : '' } })
    return output
  }

  function HandleKeyPress(e) {
    // ignore if f-keys, tab, esc or modifiers other than shift are active
    if (/^(F\d{1,2}|Tab|Escape)$/.test(e.key) || e.ctrlKey || e.altKey || e.metaKey) {
      return
    }
    e.stopPropagation()
    e.preventDefault()
    console.log(`Key ${e.key} was pressed; cursor is at ${currentIdx}`, e)
    switch (true) {
      // is printable character
      case /^.$/.test(e.key): {
        if (running == false && currentIdx == 0) {
          running = true
          timerCountDownId = setInterval(() => {
            if (timerSeconds.value == 0) {
              clearInterval(timerCountDownId)
              stats.value = ComputeStats(computedText.value)
              timerCountDownId = null
              running = false
            } else timerSeconds.value--
          }, 1000, timerSeconds)
        }
        if (currentIdx >= computedText.value.length) return
        if (running) {
          computedText.value[currentIdx].status = computedText.value[currentIdx].letter == e.key ? 'right' : 'wrong'
          currentIdx++
          if (currentIdx >= computedText.value.length) return
          computedText.value[currentIdx].status = 'cursor'
        }
        break
      }
      // is deleting prev input
      case /^Backspace$/.test(e.key) && running: {
        if (currentIdx <= 0) return
        if (currentIdx < computedText.value.length) computedText.value[currentIdx].status = ''
        currentIdx--
        computedText.value[currentIdx].status = 'cursor'
        break
      }
    // default: {
    // }
    }
  }

  function ComputeStats(values) {
    let letters = 0, words = 0, errors = 0, wpm = 0, lpm = 0
    let MODIFIER =  testTime / 60
    values.forEach((cur, idx) => {
      if (idx < currentIdx) {
        letters++
        words += /^\s$/.test(cur.letter)
        errors += cur.status == 'wrong'
      }
    })
    wpm = (words / MODIFIER).toString().match(/\d+(.\d{1,2})?/).at(0)
    lpm = (letters / MODIFIER).toString().match(/\d+(.\d{1,2})?/).at(0)
    console.log(`letters: ${letters}, words: ${words}, errors: ${errors}, time: ${testTime}, wpm: ${wpm}, lpm: ${lpm}`)
    console.log({letters: letters, words: words, errors: errors, time: testTime, wpm: wpm, lpm: lpm})
    return {letters: letters, words: words, errors: errors, time: testTime, wpm: wpm, lpm: lpm}
  }

  function Restart() {
    computedText.value = ProcessInputText(SampleText.split(''))
    currentIdx = 0
    timerSeconds.value = 5
    testTime = 5
    running = false
    if (timerCountDownId) {
      clearInterval(timerCountDownId)
      timerCountDownId = null
    }
  }

  const timerSeconds = defineModel('timerSeconds')
  const stats = defineModel('stats')

  const SampleText = "It is a long established fact that a reader will be distracted by the readable content of a page when " +
    "looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of " +
    "letters, as opposed to using 'Content here, content here', making it look like readable English. " +
    "Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a " +
    "search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, " +
    "sometimes by accident, sometimes on purpose (injected humour and the like)."

  const computedText = ref(ProcessInputText(SampleText.split('')))

  let currentIdx = 0
  let running = false
  let timerCountDownId
  let testTime

  onMounted(() => {
    testTime = timerSeconds.value
    document.addEventListener('keydown', HandleKeyPress, { capture: true })
  })

  onUnmounted(() => {
    document.removeEventListener('keydown', HandleKeyPress)
  })
</script>

<style lang="scss" scoped></style>
