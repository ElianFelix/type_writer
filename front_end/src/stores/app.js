import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vuetify/lib/composables/router.mjs'
import * as api from '@/api/api'
import { processInputText } from '@/helpers/text_processing'
import { useUserStore } from './user'

const DEFAULT_TIME = 60
const DEFAULT_TEXT = {
  "id": 1,
  "text_type": "full-text",
  "title": "db-in-default-test",
  "difficulty": "normal",
  "tags": ["testing"],
  "text_body": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).",
  "text_length": 613,
  "created_at": "2026-03-11T01:16:31.708976Z",
  "updated_at": "2026-03-11T01:16:31.708976Z"
}

export const useAppStore = defineStore('app', () => {
  const router = useRouter()
  const route = useRoute()
  const userStore = useUserStore()

  // [
  //   {
  //     "id": 1,
  //     "name": "typing-test",
  //     "description": "type the text as fast as you can under the time limit",
  //     "created_at": "2026-03-10T23:14:08.10223Z",
  //     "updated_at": "2026-03-10T23:14:08.10223Z"
  //   },
  //   {
  //     "id": 2,
  //     "name": "drill",
  //     "description": "type the text as fast as you can under the time limit",
  //     "created_at": "2026-03-10T23:14:13.622774Z",
  //     "updated_at": "2026-03-10T23:14:13.622774Z"
  //   },
  // ]
  const activities = ref()
  // [
  //   {
  //     "id": 1,
  //     "text_type": "full-text",
  //     "title": "db-in-default-test",
  //     "difficulty": "normal",
  //     "tags": ["testing"],
  //     "text_body": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).",
  //     "text_length": 613,
  //     "created_at": "2026-03-11T01:16:31.708976Z",
  //     "updated_at": "2026-03-11T01:16:31.708976Z"
  //   },
  //   {
  //     "id": 2,
  //     "text_type": "full-text",
  //     "title": "db-in-tinto-talk",
  //     "difficulty": "normal",
  //     "tags": ["testing"],
  //     "text_body": "Here, your friendly Content Design Lead Pavia will be your host today, as today @Johan has been hijacked by the Swedes invited to the Swedish Game Awards (as Europa Universalis V has been nominated for 3 categories, Best Technology, Best Design, and Game of the Year; it can also be voted by the community as Player’s Game of the Year); while @SaintDaveUK is busy with a super secret project working on the game’s first DLC, Fate of the Phoenix.",
  //     "text_length": 450,
  //     "created_at": "2026-03-11T01:39:31.739875Z",
  //     "updated_at": "2026-03-11T01:39:31.739875Z"
  //   },
  // ]
  const texts = ref(null)

  // [
  //   {
  //     id: 0,
  //     user_id: 1,
  //     activity_id: 1,
  //     text_id: 1,
  //     duration: 60,
  //     result: {
  //       wpm: 300,
  //       lpm: 300,
  //       letters: 100,
  //       words: 100,
  //       errors: 100,
  //       corrected: 0
  //     }
  //   },
  // ]
  const scores = ref(null)

  const selectedActivity = ref('typing-test')
  const selectedText = ref(0)
  const messages = ref([])

  //
  // In-game settings
  //
  const gameText = ref(DEFAULT_TEXT)
  const processedGameText = ref(processInputText(gameText.value.text_body))

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
    gameText.value = texts.value?.find((t) => t.id == selectedText.value) ?? DEFAULT_TEXT
    processedGameText.value = processInputText(gameText.value.text_body)
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
    started.value = false
    completed.value = true
    if (cursor.value < processedGameText.value.length) {
      processedGameText.value[cursor.value].status = ''
    }
    if (timerId.value) {
      clearInterval(timerId.value)
      timerId.value = null
    }
  }

  function startActivity() {
    startGame()
    router.push({ name: '/board' })
  }

  function endActivity(retry = false) {
    if (userStore.isLoggedIn) {
      addScore()
    } else {
      addMessage("Result not saved. Log in to save results", "warning")
    }
    startGame()
    if (!retry) {
      router.push({ name: '/' })
    }
  }

  function computeStats() {
    const actualTime = timerSeconds.value == 0 ? testTime.value :  testTime.value - timerSeconds.value
    const MODIFIER =  actualTime / 60
    let letters = 0, words = 0, errors = 0, wpm = 0, lpm = 0
    for (const [idx, val] of processedGameText.value.entries()) {
      if (idx < cursor.value) {
        letters++
        words += idx == 0 || /^\s$/.test(processedGameText.value[idx - 1]?.letter)
        errors += val.status == 'wrong'
      } else {
        break
      }
    }
    wpm = Number((words / MODIFIER).toString().match(/\d+(.\d{1,2})?/).at(0))
    lpm = Number((letters / MODIFIER).toString().match(/\d+(.\d{1,2})?/).at(0))
    const computedStats = [actualTime, {wpm, lpm, letters, words, errors, corrected: 0}]
    // console.log(computedStats)
    return computedStats
  }

  function addMessage(text, color='info') {
    messages.value.push({ text, color })
  }

  async function addScore() {
    if (!stats.value) {
      return
    }

    const newScoreResp = await api.createScore({
      user_id: userStore.activeUser.id,
      activity_id: activities.value.find((a) => a.name == selectedActivity.value).id,
      text_id: gameText.value.id,
      duration: stats.value[0],
      result: stats.value[1],
    }, userStore.getAuthToken)
    if (newScoreResp) {
      console.log('new score ->', newScoreResp)
      scores.value.push(newScoreResp)
    }
  }

  async function refreshActivities() {
    const activitiesResp = await api.getActivities()
    console.log('activities resp ->', activitiesResp)
    if (activitiesResp) {
      activities.value = activitiesResp
    }
  }

  async function refreshTexts() {
    const textsResp = await api.getTexts()
    console.log('texts resp ->', textsResp)
    if (textsResp) {
      texts.value = textsResp
    }
  }

  async function createOrUpdateText(text) {
    const textReq = {
      text_type: text.text_type,
      title: text.title,
      difficulty: text.difficulty,
      tags: text.tags.split(","),
      text_body: text.text_body
    }

    const textResp =
      text.id == 0
      ? await api.createText(textReq, userStore.getAuthToken)
      : await api.updateText(text.id, textReq, userStore.getAuthToken)

    if (textResp) {
      await refreshTexts()
      addMessage("Successful text creation/update", "success")
      return textResp.id
    } else {
      addMessage("Failed text creation/update", "error")
      return 0
    }
  }

  async function refreshScores() {
    const scoresResp = await api.getScores()
    console.log('scores resp ->', scoresResp)
    if (scoresResp) {
      scores.value = scoresResp
    }
  }

  refreshActivities()
  refreshTexts()
  refreshScores()

  return {
    // State
    activities, texts, selectedActivity, selectedText, gameText, processedGameText,
    stats, timerId, started, completed, testTime, timerSeconds, fontSize, cursor,
    scores, messages,
    // Getters
    getSelectedActivity, getSelectedText, getDefaultTime,
    // Actions
    incrementTestTime, decrementTestTime, incrementFontSize, decrementFontSize,
    startGame, pauseGame, endGame, startActivity, endActivity, computeStats, addScore,
    refreshActivities, refreshTexts, refreshScores, addMessage, createOrUpdateText,
  }
})
