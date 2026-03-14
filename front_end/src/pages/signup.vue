<template>
  <div class="d-flex flex-column align-center fill-height main-container">
    <v-sheet class="d-flex flex-column ga-4 mt-16 px-16 py-12" width="500">
      <v-form @submit.prevent="handleFormSubmit">
        <v-text-field
          v-model="userInfo.username"
          :rules="usernameRules"
          label="Username"
          variant="outlined"
          density="compact"
        ></v-text-field>
        <v-text-field
          v-model="userInfo.password"
          :type="false ? 'text' : 'password'"
          hint="At least 8 characters"
          :rules="passwordRules"
          label="Password"
          variant="outlined"
          density="compact"
        ></v-text-field>
        <v-text-field
          v-model="userInfo.name"
          :rules="nameRules"
          label="Name"
          variant="outlined"
          density="compact"
        ></v-text-field>
        <v-text-field
          v-model="userInfo.email"
          :rules="emailRules"
          label="Email"
          variant="outlined"
          density="compact"
        ></v-text-field>
        <div class="d-flex flex-column justify-end align-stretch mt-4">
          <v-label class="text-label-small">Already have an account?</v-label>
          <div class="d-flex justify-space-between align-end">
            <v-btn @click="() => router.push('/login')">Login</v-btn>
            <v-btn-group density="compact">
              <v-btn @click="() => router.push('/')">Cancel</v-btn>
              <v-btn type="submit">Sign Up</v-btn>
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

  const signedUp = ref(false)
  const userInfo = ref({
    username: '',
    password: '',
    name: '',
    email: '',
    user_type: 'regular',
  })

  const usernameRules = [
    (v) => /^\w+$/.test(v)
  ]
  const passwordRules = [
    (v) => /^.{8,}$/.test(v)
  ]
  const nameRules = [
    (v) => v.length === 0 || /^(\w+)$/.test(v)
  ]
  const emailRules = [
    (v) => /^\w+@\w+\.\w{3}$/.test(v)
  ]

  watchEffect(() => {if (signedUp.value || appStore.activeUser) router.push('/')})

  function handleFormSubmit() {
    signedUp.value = appStore.signUpUser(userInfo.value)
  }
</script>
