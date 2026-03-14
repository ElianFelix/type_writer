import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vuetify/lib/composables/router.mjs'
import { processInputText } from '@/helpers/textProcessing'
import * as api from '@/api/api'

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

  // hardcoded user to build up functionality
  const activeUser = ref({
    id: 1,
    user_type: "regular",
    username: "testivo",
    name: "testivo",
    email: "test@user.com",
    created_at: "2026-03-11T04:00:02.518314Z",
    updated_at: "2026-03-11T18:35:11.425319164Z"
  })

  const users = ref([
    {
      id: 1,
      user_type: "regular",
      username: "testivo",
      name: "testivo",
      email: "test@user.com",
      created_at: "2026-03-11T04:00:02.518314Z",
      updated_at: "2026-03-11T18:35:11.425319164Z"
    },
  ])

  const activities = ref([
    {
      "id": 1,
      "name": "typing-test",
      "description": "type the text as fast as you can under the time limit",
      "created_at": "2026-03-10T23:14:08.10223Z",
      "updated_at": "2026-03-10T23:14:08.10223Z"
    },
    {
      "id": 2,
      "name": "drill",
      "description": "type the text as fast as you can under the time limit",
      "created_at": "2026-03-10T23:14:13.622774Z",
      "updated_at": "2026-03-10T23:14:13.622774Z"
    },
  ])
  const texts = ref([
    {
        "id": 1,
        "text_type": "full-text",
        "title": "db-in-default-test",
        "difficulty": "normal",
        "text_body": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).",
        "text_length": 613,
        "created_at": "2026-03-11T01:16:31.708976Z",
        "updated_at": "2026-03-11T01:16:31.708976Z"
    },
    {
        "id": 2,
        "text_type": "full-text",
        "title": "db-in-tinto-talk",
        "difficulty": "normal",
        "text_body": "Here, your friendly Content Design Lead Pavia will be your host today, as today @Johan has been hijacked by the Swedes invited to the Swedish Game Awards (as Europa Universalis V has been nominated for 3 categories, Best Technology, Best Design, and Game of the Year; it can also be voted by the community as Player’s Game of the Year); while @SaintDaveUK is busy with a super secret project working on the game’s first DLC, Fate of the Phoenix.",
        "text_length": 450,
        "created_at": "2026-03-11T01:39:31.739875Z",
        "updated_at": "2026-03-11T01:39:31.739875Z"
    }
  ])
  const selectedActivity = ref('typing-test')
  // const selectedText = ref('default-text')
  const selectedText = ref(0)

  const activityResults = ref([
    // later model
    // {
    //   id: 0,
    //   user_id: 0,
    //   activity_id: 0,
    //   text_id: 0,
    //   type: 'typing-test',
    //   title: 'place-holder-title',
    //   time: 300,
    //   points: 300,
    //   errors: 300,
    //   created_at: "timestamp",
    //   updated_at: "timestamp",
    //   result: { // later implementation will use result as a sub struct
    //     wpm: 300,
    //     lpm: 300,
    //     letters: 300,
    //     words: 300,
    //     errors: 300,
    //     corrected: 300,
    //   },
    // },
      {
        id: 0,
        user_id: 1,
        activity_id: 1,
        text_id: 1,
        points: 300,
        duration: 60,
        errors: 300
      },
  ])


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
    gameText.value = texts.value[selectedText.value] ?? DEFAULT_TEXT
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

  function endActivity() {
    addActivityResult()
    router.push({ name: '/' })
  }

  async function addActivityResult() {
    if (!stats.value) {
      return
    }
    const resultText = texts.value[selectedText.value]
    // activityResults.value.unshift({
    //   id: activityResults.value.length,
    //   user_id: 0,
    //   activity_id: 0,
    //   text_id: 0,
    //   type: selectedActivity.value,
    //   title: resultText.title,
    //   time: stats.value.time,
    //   points: stats.value.wpm,
    //   errors: stats.value.errors,
    //   created_at: "timestamp",
    //   updated_at: "timestamp",
    //   result: {
    //     wpm: stats.value.wpm,
    //     lpm: stats.value.lpm,
    //     letters: stats.value.letters,
    //     words: stats.value.words,
    //     errors: stats.value.errors,
    //     corrected: stats.value.corrected,
    //   }
    // })

    const newScoreResp = await api.createScore({
      user_id: activeUser.value.id,
      activity_id: activities.value.find((a) => a.name == selectedActivity.value).id,
      text_id: resultText.id,
      points: stats.value.wpm,
      duration: stats.value.time,
      errors: stats.value.errors
    })
    if (newScoreResp) {
      console.log('new score ->', newScoreResp)
      activityResults.value.push(newScoreResp)
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

  async function refreshActivityResults() {
    const scoresResp = await api.getScores()
    console.log('scores resp ->', scoresResp)
    if (scoresResp) {
      activityResults.value = scoresResp
    }
  }

  async function refreshUsers() {
    const usersResp = await api.getUsers()
    console.log('texts resp ->', usersResp)
    if (usersResp) {
      users.value = usersResp
    }
  }

  async function signUpUser(userInfo = {}) {
    const signUpResp = await api.createUser(userInfo)
    if (signUpResp) {
      activeUser.value = signUpResp
      return true
    }
    return false
  }

  async function loginUser(loginInfo = {}) {
    const loginResp = await api.getUserByIdOrUsername(loginInfo.username)
    if (loginResp) {
      activeUser.value = loginResp
      return true
    }
    return false
  }

  refreshActivities()
  refreshTexts()
  refreshActivityResults()
  refreshUsers()

  return {
    // State
    users, activities, texts, selectedActivity, selectedText, gameText, processedGameText,
    stats, timerId, started, completed, testTime, timerSeconds, fontSize, cursor,
    activityResults, activeUser,
    // Getters
    getSelectedActivity, getSelectedText, getDefaultTime,
    // Actions
    incrementTestTime, decrementTestTime, incrementFontSize, decrementFontSize,
    startGame, pauseGame, endGame, startActivity, endActivity, addActivityResult,
    refreshActivities, refreshTexts, refreshActivityResults, signUpUser, loginUser,
    refreshUsers,
  }
})
