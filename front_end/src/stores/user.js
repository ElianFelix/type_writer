import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vuetify/lib/composables/router.mjs'
import * as api from '@/api/api'
import { useAppStore } from './app'

export const useUserStore = defineStore('user', () => {
  const appStore = useAppStore()

  // [
  //   {
  //     id: 1,
  //     user_type: "regular",
  //     username: "testivo",
  //     name: "testivo",
  //     email: "test@user.com",
  //     created_at: "2026-03-11T04:00:02.518314Z",
  //     updated_at: "2026-03-11T18:35:11.425319164Z"
  //   },
  // ]
  const users = ref(null)

  // {
  //   id: 1,
  //   user_type: "regular",
  //   username: "testivo",
  //   name: "testivo",
  //   email: "test@user.com",
  //   created_at: "2026-03-11T04:00:02.518314Z",
  //   updated_at: "2026-03-11T18:35:11.425319164Z"
  // }
  const activeUser = ref(JSON.parse(localStorage.getItem('user.active_user')))
  const authToken = ref(JSON.parse(localStorage.getItem('user.auth_token')))

  const getActiveUser = computed(() => activeUser.value)
  const getAuthToken = computed(() => authToken.value)
  const isLoggedIn = computed(() => activeUser.value !== null && authToken.value !== null)
  const isAdmin = computed(() => activeUser.value !== null && activeUser.value.user_type == 'admin')

  function logoutUser() {
    activeUser.value = null
    authToken.value = null
    localStorage.removeItem('user.active_user')
    localStorage.removeItem('user.auth_token')
  }

  async function refreshUsers() {
    const usersResp = await api.getUsers()
    console.log('users resp ->', usersResp)
    if (usersResp) {
      users.value = usersResp
    }
  }

  async function loginUser(loginInfo = {}) {
    const loginResp = await api.login(loginInfo)
    if (loginResp) {
      activeUser.value = loginResp.active_user
      authToken.value = loginResp.token
      localStorage.setItem('user.active_user', JSON.stringify(activeUser.value))
      localStorage.setItem('user.auth_token', JSON.stringify(authToken.value))
      return true
    }
    appStore.addMessage("Failed login", "error")
    return false
  }

  async function signUpUser(userInfo = {}) {
    const signUpResp = await api.createUser(userInfo)
    if (signUpResp) {
      activeUser.value = signUpResp
      const loggedIn = await loginUser({ username: userInfo.username, password: userInfo.password})
      return loggedIn
    }
    appStore.addMessage("Failed sign up", "error")
    return false
  }

  async function updateUserDetails(userInfo = {}) {
    const updateResp = await api.updateUser(userInfo, authToken.value)
    if (updateResp) {
      activeUser.value = updateResp
      localStorage.setItem('user.active_user', JSON.stringify(activeUser.value))
      appStore.addMessage("User details updated")
      return true
    }
    appStore.addMessage("Failed user details update", "error")
    return false
  }

  refreshUsers()

  return {
    // State
    users, activeUser, authToken,
    // Getters
    getActiveUser, getAuthToken, isLoggedIn, isAdmin,
    // Actions
    logoutUser, loginUser, refreshUsers, signUpUser, updateUserDetails,
  }
})
