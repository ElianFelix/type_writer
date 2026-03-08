// Utilities
import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vuetify/lib/composables/router.mjs'
import { processInputText } from '@/helpers/textProcessing'

const DEFAULT_TIME = 60
const DEFAULT_TEXT = "It is a long established fact that a reader will be distracted by the readable content of a page when " +
    "looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of " +
    "letters, as opposed to using 'Content here, content here', making it look like readable English. " +
    "Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a " +
    "search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, " +
    "sometimes by accident, sometimes on purpose (injected humour and the like)."

export const useAppStore = defineStore('app', () => {
  const router = useRouter()
  const route = useRoute()

  // State
  const activities = ref({
    'typingTest': {
      description: 'type the text as fast as you can under the time limit',
      settings: {},
      items: [
        {
          title: 'default-text',
          difficulty: 'normal',
          text: DEFAULT_TEXT,
          tags: ['testing', 'default'],
        },
        {
          title: 'typing-place-holder',
          difficulty: 'normal',
          text: 'some lorem ipsum testing placeholder some lorem ipsum testing placeholder  some lorem ipsum testing placeholder',
          tags: ['testing'],
        },
        {
          title: 'typing-place-holder-2',
          difficulty: 'hard',
          text: 'some lorem ipsum testing placeholder',
          tags: ['testing'],
        },
        {
          title: 'typing-place-holder-3',
          difficulty: 'hard',
          text: 'some lorem ipsum testing placeholder',
          tags: ['testing'],
        },
        {
          title: 'typing-place-holder-4',
          difficulty: 'hard',
          text: 'some lorem ipsum testing placeholder',
          tags: ['testing'],
        },
        {
          title: 'typing-place-holder-5',
          difficulty: 'hard',
          text: 'some lorem ipsum testing placeholder',
          tags: ['testing'],
        },
      ],
    },
    'drill': {
      description: 'hit the keys as the appear on screen to practice your accuracy',
      settings: {},
      items: [
        {
          title: 'drill-place-holder',
          difficulty: 'normal',
          text: 'some lorem ipsum testing placeholder',
          tags: ['testing'],
        },
      ],
    },
  })
  const selectedActivity = ref('typingTest')
  // const selectedText = ref('default-text')
  const selectedText = ref(3)


  //
  // In-game settings
  //
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
  //
  //
  //

  // Getters
  const getSelectedActivity = computed(() => selectedActivity.value)
  const getSelectedText = computed(() => selectedText.value)
  const getDefaultTime = computed(() => DEFAULT_TIME)

  // Actions
  function incrementTestTime(increment = 15) {
    const increasedTime = testTime.value + increment
    if (increasedTime > 180) {
      return
    }
    startGame(increasedTime)
  }

  function decrementTestTime(decrement = 15) {
    const decreasedTime = testTime.value + decrement
    if (decreasedTime < 15) {
      return
    }
    startGame(decreasedTime)
  }

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

  function startGame(newTime = testTime.value) {
    testTime.value = newTime
    timerSeconds.value = newTime
    stats.value = null
    started.value = false
    completed.value = false
    cursor.value = 0
    gameText.value = activities.value[selectedActivity.value].items[selectedText.value] ?? DEFAULT_TEXT
    processedGameText.value = processInputText(gameText.value.text)
    if (timerId.value) {
      clearInterval(timerId.value)
      timerId.value = null
    }
  }

  function restartGame(newTime = testTime.value) {
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

  function pauseGame() {
    started.value = false
    if (timerId.value) {
      clearInterval(timerId.value)
      timerId.value = null
    }
  }

  function endGame() {
    clearInterval(timerId.value)
    timerId.value = null
    started.value = false
    completed.value = true
    if (cursor < processedGameText.value.length) {
      processedGameText.value[cursor.value].status = ''
    }
  }

  function startActivity() {
    startGame()
    router.push({ name: '/board' })
  }

  function endActivity() {
    router.push({ name: '/' })
  }

  return {
    // State
    activities, selectedActivity, selectedText, gameText, processedGameText,
    stats, timerId, started, completed, testTime, timerSeconds, fontSize, cursor,
    // Getters
    getSelectedActivity, getSelectedText, getDefaultTime,
    // Actions
    incrementTestTime, decrementTestTime, incrementFontSize, decrementFontSize,
    startGame, restartGame, pauseGame, endGame, startActivity, endActivity,
  }
})
