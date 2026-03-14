<template>
  <div class="d-flex flex-column align-center fill-height main-container">
    <v-sheet class="d-flex flex-column ga-4 mt-16 px-16 py-12" width="500">
      <v-form @submit.prevent="handleFormSubmit">
        <v-text-field
          v-model="loginInfo.username"
          :rules="usernameRules"
          label="Username"
          variant="outlined"
          density="compact"
        ></v-text-field>
        <v-text-field
          v-model="loginInfo.password"
          :type="false ? 'text' : 'password'"
          hint="At least 8 characters"
          :rules="passwordRules"
          label="Password"
          variant="outlined"
          density="compact"
        ></v-text-field>
        <div class="d-flex flex-column justify-end align-stretch mt-4">
          <v-label class="text-label-small">Don't have an account?</v-label>
          <div class="d-flex justify-space-between align-end">
            <v-btn @click="() => router.push('/signup')">Sign Up</v-btn>
            <v-btn-group density="compact">
              <v-btn @click="() => router.push('/')">Cancel</v-btn>
              <v-btn type="submit">Login</v-btn>
            </v-btn-group>
          </div>
        </div>
      </v-form>
    </v-sheet>
  </div>
</template>

<script setup>
  import { ref, watchEffect } from 'vue';
  import { useRouter } from 'vuetify/lib/composables/router.mjs';
  import { useAppStore } from '@/stores/app';

  const appStore = useAppStore()
  const router = useRouter()

  const loggedIn = ref(false)
  const loginInfo = ref({
    username: '',
    password: '',
  })

  const usernameRules = [
    (v) => /^\w+$/.test(v)
  ]
  const passwordRules = [
    (v) => /^.{8,}$/.test(v)
  ]

  watchEffect(() => {if (loggedIn.value) router.push('/')})

  function handleFormSubmit() {
    loggedIn.value = appStore.loginUser(loginInfo.value)
  }
</script>
