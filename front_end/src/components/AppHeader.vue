<template>
  <div class="d-flex ga-3 position-absolute top-0 right-0 justify-center align-center ma-5">
    <v-sheet class="d-flex flex-column pa-2 justify-space-between"
             :class="{ 'bg-transparent': !openMenu, 'open-menu': openMenu, 'closed-menu': !openMenu }"
             tabindex="0"
             rounded
    >
      <div class="d-flex ga-2 justify-space-between">
        <v-btn v-if="!openMenu" icon="mdi-account-circle"></v-btn>
        <v-btn v-if="!openMenu" :icon="activeThemeIcon" @click="changeActiveTheme"></v-btn>
        <v-btn :icon="openMenu ? 'mdi-close' : 'mdi-menu'" @click="() => {openMenu = !openMenu}"></v-btn>
      </div>
      <div v-if="openMenu" class="d-flex flex-column ga-1 menu-section ma-5 mt-0">
        <div class="menu-item mx-4">
          <div class="menu-sub-heading">Settings</div>
          <div class="d-flex ga-2 justify-start">
            <v-btn size="small" icon="mdi-plus-circle"></v-btn>
            <v-btn size="small" icon="mdi-minus-circle"></v-btn>
            <v-btn size="small" icon="mdi-restore"></v-btn>
            <v-btn size="small" icon="mdi-stop"></v-btn>
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
                      color="blue"
            >
              <template v-slot:label>
                {{ activeTheme }}
              </template>
            </v-switch>
          </div>
        </div>
        <div class="menu-item"><v-list-item>Home</v-list-item></div>
        <div class="menu-item"><v-list-item>Profile</v-list-item></div>
        <div class="menu-item"><v-list-item>Login/logout</v-list-item></div>
      </div>
    </v-sheet>
  </div>
</template>

<script setup>
  import { computed, ref } from 'vue';
  import { useTheme } from 'vuetify'

  function changeActiveTheme() {
    activeTheme.value = activeTheme.value == 'dark' ? 'light' : 'dark'
    theme.change(activeTheme.value)
  }

  const theme = useTheme()
  const openMenu = ref(false)
  const activeTheme = ref('dark')
  const activeThemeIcon = computed(() => {
    console.log('current theme here -> ', theme.current.value)
    return theme.current.value.dark ? 'mdi-weather-night' : 'mdi-white-balance-sunny'
  })
</script>

<style lang="scss" scoped>
</style>
