<template>
  <div class="d-flex flex-column ga-4 fill-height main-container">
    <v-expansion-panels v-model="activeSections">
      <v-expansion-panel title="Profile Details" value="profile-details">
        <template v-slot:text>
          <v-form ref="form" @submit.prevent="handleFormSubmit">
            <div class="d-flex justify-space-between align-end mb-6">
              <div>Account type: {{ userInfo.user_type }} user</div>
              <v-avatar color="secondary" size="large" icon="mdi-account-circle"></v-avatar>
            </div>
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
            <div class="d-flex justify-end">
              <v-btn type="submit">Update Details</v-btn>
            </div>
          </v-form>
        </template>
      </v-expansion-panel>
      <v-expansion-panel title="Settings Prefrences" value="settings-preferences">
        <template v-slot:text>
          settings preferences inputs
        </template>
      </v-expansion-panel>
    </v-expansion-panels>
  </div>
</template>

<script setup>
  import { onMounted, ref, watchEffect } from 'vue';
  import { useAppStore } from '@/stores/app';
  import { usernameRules, passwordRules, nameRules, emailRules } from '@/helpers/input_rules'
  import { useRouter } from 'vuetify/lib/composables/router.mjs';

  const appStore = useAppStore()
  const router = useRouter()

  const form = ref()
  const userInfo = ref({
    id: 0,
    username: '',
    password: '',
    name: '',
    email: '',
    user_type: 'regular',
  })
  const activeSections = ref(['profile-details'])

  watchEffect(() => {if (!appStore.activeUser) router.push('/login')})

  onMounted(() => {
    userInfo.value = { ...appStore.activeUser }
  })

  async function handleFormSubmit() {
    if (form.value.isValid) {
      await appStore.updateUserDetails(userInfo.value)
    }
  }
</script>
