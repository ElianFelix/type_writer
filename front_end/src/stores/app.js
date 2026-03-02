// Utilities
import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { processInputText } from '@/helpers/textProcessing'

const DEFAULT_TIME = 60
const DEFAULT_TEXT = "It is a long established fact that a reader will be distracted by the readable content of a page when " +
    "looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of " +
    "letters, as opposed to using 'Content here, content here', making it look like readable English. " +
    "Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a " +
    "search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, " +
    "sometimes by accident, sometimes on purpose (injected humour and the like)."

export const useAppStore = defineStore('app', () => {
  // Hardcoded for now TODO
  const gameText = ref(DEFAULT_TEXT)
  const processedGameText = ref(processInputText(gameText.value))

  const stats = ref(null)
  const timerId = ref(null)

  const started = ref(false)
  const completed = ref(false)

  const testTime = ref(DEFAULT_TIME)
  const timerSeconds = ref(DEFAULT_TIME)
  const fontSize = ref(2)
  const cursor = ref(0)

  const getDefaultTime = computed(() => DEFAULT_TIME)

  function incrementFontSize() {
    if (fontSize.value < 3) {
      fontSize.value += 1
    }
  }

  function decrementFontSize() {
    if (fontSize.value > 1) {
      fontSize.value -= 1
    }
  }

  function restartGame(newTime = DEFAULT_TIME) {
    testTime.value = newTime
    timerSeconds.value = newTime
    stats.value = null
    started.value = false
    completed.value = false
    cursor.value = 0
    gameText.value = DEFAULT_TEXT
    processedGameText.value = processInputText(gameText.value)
    if (timerId.value) {
      clearInterval(timerId.value)
      timerId.value = null
    }
  }

  function endGame() {
    //TODO
    return
  }

  return {
    // State
    gameText, processedGameText, stats, timerId, started, completed, testTime, timerSeconds, fontSize, cursor,
    // Getters
    getDefaultTime,
    // Actions
    incrementFontSize, decrementFontSize, restartGame, endGame,
  }
})
