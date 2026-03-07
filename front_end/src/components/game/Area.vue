<template>
  <div class="d-flex flex-column flex-grow-1 align-self-center justify-top align-center mx-8 area-container">
    <v-sheet class="px-10 py-4 text-pre-wrap" :class="`font-mono-${fontSize ?? 2}`">
      <div class="start-tooltip"
           :style="`visibility: ${appStore.started || appStore.completed ? 'hidden' : 'visible'};left: ${tooltipOffsetx}px; top: ${tooltipOffsety}px`"
      >Start typing</div>
      <span v-for="elem, index in processedText" :key="index" :id="index" :class="elem.status">{{ elem.letter }}</span>
    </v-sheet>
    <div v-if="appStore.completed" class="d-flex ga-4 py-5">
      <v-btn @click="appStore.restartGame(15)">Retry</v-btn>
      <v-btn @click="appStore.restartGame()">Done</v-btn>
    </div>
  </div>
</template>

<script setup>
  import { useAppStore } from '@/stores/app'
  import { onMounted, onUnmounted, watch } from 'vue'

  function handleKeyPress(e) {
    console.log(`Key ${e.key} was pressed; cursor is at ${appStore.cursor}`, e)
    switch (true) {
      // ignore if f-keys, tab, esc or modifiers other than shift are active
      case (/^(F\d{1,2}|Tab|Escape|Enter)$/.test(e.key) ||
        e.ctrlKey ||
        e.altKey ||
        e.metaKey ||
        e.target.matches('#mainMenu *')
      ): {
        return
      }
      // is printable character
      case /^.$/.test(e.key): {
        if (appStore.started == false) {
          appStore.started = true
          appStore.timerId = setInterval(() => {
            if (timerSeconds.value == 0) {
              clearInterval(appStore.timerId)
              stats.value = computeStats(processedText.value)
              appStore.timerId = null
              appStore.started = false
              appStore.completed = true
              if (appStore.cursor < processedText.value.length) {
                processedText.value[appStore.cursor].status = ''
              }
            } else {
              timerSeconds.value--
            }
          }, 1000, timerSeconds)
        }
        if (appStore.cursor >= processedText.value.length) return
        if (appStore.started && appStore.timerSeconds > 0) {
          processedText.value[appStore.cursor].status = processedText.value[appStore.cursor].letter == e.key ? 'right' : 'wrong'
          appStore.cursor++
          if (appStore.cursor >= processedText.value.length) return
          processedText.value[appStore.cursor].status = 'cursor'
        }
        break
      }
      // is deleting prev input
      case /^Backspace$/.test(e.key) && appStore.started: {
        if (appStore.cursor <= 0) return
        if (appStore.cursor < processedText.value.length) processedText.value[appStore.cursor].status = ''
        appStore.cursor--
        processedText.value[appStore.cursor].status = 'cursor'
        break
      }
    // default: {
    // }
    }
    e.stopPropagation()
    e.preventDefault()
  }

  function computeStats(values) {
    let letters = 0, words = 0, errors = 0, wpm = 0, lpm = 0
    const MODIFIER =  testTime.value / 60
    for (const [idx, val] of values.entries()) {
      if (idx < appStore.cursor) {
        letters++
        words += idx == 0 || /^\s$/.test(values[idx - 1]?.letter)
        errors += val.status == 'wrong'
      } else {
        break
      }
    }
    wpm = Number((words / MODIFIER).toString().match(/\d+(.\d{1,2})?/).at(0))
    lpm = Number((letters / MODIFIER).toString().match(/\d+(.\d{1,2})?/).at(0))
    console.log(`letters: ${letters}, words: ${words}, errors: ${errors}, time: ${testTime.value}, wpm: ${wpm}, lpm: ${lpm}`)
    console.log({letters: letters, words: words, errors: errors, time: testTime.value, wpm: wpm, lpm: lpm})
    return {letters: letters, words: words, errors: errors, time: testTime.value, wpm: wpm, lpm: lpm}
  }

  const appStore = useAppStore()

  const timerSeconds = defineModel('timerSeconds')
  const testTime = defineModel('testTime')
  const stats = defineModel('stats')
  const fontSize = defineModel('fontSize')

  const tooltipOffsetx = ref(7)
  const tooltipOffsety = ref(0)

  const processedText = computed({
    get() {
      return appStore.processedGameText
    },
    set(newValue) {
      appStore.processedGameText.value = newValue
    },
  })

  watch([() => appStore.started, () => appStore.fontSize, () => appStore.cursor], ([started, _, cursor]) => {
    if (!started || cursor == 0) {
      const cursorElmnt = document.querySelector('.cursor')
      if (cursorElmnt) {
        const startTooltip = document.querySelector('.start-tooltip')
        // console.log(cursor)
        // console.log(cursor.offsetLeft, cursor.offsetWidth, startTooltip.offsetWidth)
        tooltipOffsetx.value = cursorElmnt.offsetLeft + cursorElmnt.offsetWidth/2 - startTooltip.offsetWidth/2
        tooltipOffsety.value = cursorElmnt.offsetTop - cursorElmnt.offsetHeight/4 - startTooltip.offsetHeight
        // console.log(tooltipOffsetx.value, tooltipOffsety.value)
      }

    }
  }, { flush: 'post' })

  onMounted(() => {
    document.addEventListener('keydown', handleKeyPress)

    const cursor = document.querySelector('.cursor')
    if (cursor) {
      const startTooltip = document.querySelector('.start-tooltip')
      tooltipOffsetx.value = cursor.offsetLeft + cursor.offsetWidth/2 - startTooltip.offsetWidth/2
      tooltipOffsety.value = cursor.offsetTop - cursor.offsetHeight/4 - startTooltip.offsetHeight
    }
  })

  onUnmounted(() => {
    document.removeEventListener('keydown', handleKeyPress)
  })
</script>

<style lang="scss" scoped></style>
