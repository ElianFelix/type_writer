<template>
  <div class="d-flex ga-3 position-absolute top-0 right-0 justify-center align-center my-5"
       id="mainMenu"
       tabindex="9"
  >
    <v-sheet class="d-flex flex-column pa-2 elevaltion-2 justify-space-between"
             :class="{ 'bg-transparent': !openMenu, 'open-menu': openMenu, 'closed-menu': !openMenu }"
             rounded
    >
      <div class="d-flex ga-2 justify-space-between">
        <v-btn v-if="!openMenu" icon="mdi-account-circle"></v-btn>
        <v-btn v-if="!openMenu" :icon="activeThemeIcon" @click="changeActiveTheme"></v-btn>
        <v-btn :icon="openMenu ? 'mdi-close' : 'mdi-menu'" @click="toggleMenu"></v-btn>
      </div>
      <div class="flex-column ga-1 menu-section ma-5 mt-0" :class="openMenu ? 'd-flex' : 'd-none'">
        <div v-if="showSettings" class="d-flex flex-column ga-1 mx-4 mb-1 menu-item">
          <div class="menu-sub-heading">Settings</div>
          <div class="d-flex ga-2 justify-start">
            <v-btn :disabled="fontSize == 3" size="small" icon="mdi-plus-circle" @click="appStore.incrementFontSize"></v-btn>
            <v-btn :disabled="fontSize == 1" size="small" icon="mdi-minus-circle" @click="appStore.decrementFontSize"></v-btn>
            <v-btn size="small" icon="mdi-restore" @click="appStore.restartGame()"></v-btn>
            <v-btn size="small" icon="mdi-stop"></v-btn>
          </div>
          <div class="mx-auto">
            <v-number-input
              v-model="timeInput"
              persistent-placeholder
              placeholder="secs"
              :min="15" :max="180"
              width="150"
              density="compact"
              variant="solo-filled"
              control-variant="split"
              hide-details
              :step="15"
              @update:model-value="appStore.restartGame(timeInput)"
            ></v-number-input>
            <!--<v-btn-group border density="compact" variant="tonal">
              <v-btn icon="mdi-minus"></v-btn>
              <v-text-field hide-details></v-text-field>
              <v-btn icon="mdi-plus"></v-btn>
              <v-btn icon="mdi-check"></v-btn>
            </v-btn-group> -->
            <!-- <v-btn size="small" icon="mdi-check"></v-btn> -->
          </div>
        </div>
        <div class="menu-item mx-4">
          <div class="menu-sub-heading">Theme</div>
          <div class="ml-4">
            <v-switch v-model="activeTheme"
                      @update:modelValue="() => {theme.change(activeTheme)}"
                      false-value="light"
                      :false-icon="activeThemeIcon"
                      true-value="dark"
                      :true-icon="activeThemeIcon"
                      density="compact" hide-details
                      color="secondary"
            >
              <template v-slot:label>
                {{ activeTheme }}
              </template>
            </v-switch>
          </div>
        </div>
        <div class="menu-item"><v-list-item to="/">Home</v-list-item></div>
        <div class="menu-item"><v-list-item to="/board">Board</v-list-item></div>
        <div class="menu-item"><v-list-item>Profile</v-list-item></div>
        <div class="menu-item"><v-list-item to="/login">Login</v-list-item></div>
        <div class="menu-item"><v-list-item to="/signup">Signup</v-list-item></div>
      </div>
    </v-sheet>
    <div v-if="openMenu" class="menu-backdrop" @click="toggleMenu"></div>
  </div>
</template>

<script setup>
  import { computed, onMounted, onUnmounted, ref } from 'vue';
  import { useTheme } from 'vuetify'
  import { useAppStore } from '@/stores/app';
  import { useRoute } from 'vuetify/lib/composables/router.mjs';

  function changeActiveTheme() {
    activeTheme.value = activeTheme.value == 'dark' ? 'light' : 'dark'
    theme.change(activeTheme.value)
  }

  function toggleMenu(e) {
    const mainMenu = document.querySelector('#mainMenu')
    if (openMenu.value) {
      mainMenu.blur()
    } else {
      mainMenu.focus()
      if (appStore.started) {
        appStore.pauseGame()
      }
    }
    openMenu.value = !openMenu.value
  }

  function handleKeyPress(e) {
    if (/^Escape$/.test(e.key)) {
      toggleMenu()
    }
  }

  const appStore = useAppStore()
  const theme = useTheme()
  const route = useRoute()

  const openMenu = ref(false)
  const activeTheme = ref('dark')
  const timeInput = ref(appStore.testTime)

  const activeThemeIcon = computed(() => {
    console.log('current theme here ->', theme.current.value)
    return theme.current.value.dark ? 'mdi-weather-night' : 'mdi-white-balance-sunny'
  })
  const fontSize = computed(() => appStore.fontSize)
  const showSettings = computed(() => route.value.name.match(/board/))

  onMounted(() => {
    document.addEventListener('keydown', handleKeyPress)
  })

  onUnmounted(() => {
    document.removeEventListener('keydown', handleKeyPress)
  })
</script>

<style lang="scss" scoped>
</style>
