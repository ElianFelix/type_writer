import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vuetify/lib/composables/router.mjs'
import * as api from '@/api/api'

export const useUserStore = defineStore('user', () => {
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
  const users = ref()

  // {
  //   id: 1,
  //   user_type: "regular",
  //   username: "testivo",
  //   name: "testivo",
  //   email: "test@user.com",
  //   created_at: "2026-03-11T04:00:02.518314Z",
  //   updated_at: "2026-03-11T18:35:11.425319164Z"
  // }
  const activeUser = ref()
  const authToken = ref("")

  const getActiveUser = computed(() => activeUser.value)

  async function loginUser(loginInfo = {}) {
    const loginResp = await api.login(loginInfo)
    if (loginResp) {
      activeUser.value = loginResp.active_user
      authToken.value = loginResp.token
      return true
    }
    return false
  }

  async function refreshUsers() {
    const usersResp = await api.getUsers()
    console.log('users resp ->', usersResp)
    if (usersResp) {
      users.value = usersResp
    }
  }

  async function signUpUser(userInfo = {}) {
    const signUpResp = await api.createUser(userInfo)
    if (signUpResp) {
      activeUser.value = signUpResp
      const loggedIn = await loginUser({ ...userInfo.username, ...userInfo.password})
      return loggedIn
    }
    return false
  }

  async function updateUserDetails(userInfo = {}) {
    const updateResp = await api.updateUser(userInfo)
    if (updateResp) {
      activeUser.value = updateResp
      return true
    }
    return false
  }

  refreshUsers()

  return {
    // State
    users, activeUser, authToken,
    // Getters
    getActiveUser,
    // Actions
    loginUser, refreshUsers, signUpUser, updateUserDetails,
  }
})
