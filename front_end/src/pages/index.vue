<template>
  <div class="d-flex flex-column ga-4 fill-height main-container">
    <div class="d-flex justify-center align-center bg-red histogram-container">
      <v-table
        class="w-100 h-100"
        fixed-header
      >
        <thead>
          <tr>
            <th class="text-left">
              User
            </th>
            <th class="text-left">
              Activitity
            </th>
            <th class="text-left">
              Title
            </th>
            <th class="text-left">
              WPM
            </th>
            <th class="text-left">
              LPM
            </th>
            <th class="text-left">
              Time
            </th>
            <th class="text-left">
              Letters
            </th>
            <th class="text-left">
              Words
            </th>
            <th class="text-left">
              Errors
            </th>
            <th class="text-left">
              Corrected
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="item in appStore.activityResults.sort(sortResultsIdDescending)" :key="item.id"
          >
            <td>{{ appStore.activeUser.username }}</td>
            <td>{{ appStore.activities.find((a) => a.id == item.activity_id).name }}</td>
            <td>{{ appStore.texts.find((t) => t.id == item.text_id).title }}</td>
            <td>{{ item.points }}</td>
            <td>{{ 0 }}</td>
            <td>{{ item.duration }} secs</td>
            <td>{{ 0 }}</td>
            <td>{{ 0 }}</td>
            <td>{{ item.errors }}</td>
            <td>{{ 0 }}</td>
          </tr>
        </tbody>
      </v-table>
    </div>
    <div class="d-flex justify-center align-center bg-red select-container">
      <v-sheet class="h-100 w-100">
        <v-tabs v-model="appStore.selectedActivity" color="primary" density="compact">
          <v-tab value="typing-test">Typing Test</v-tab>
          <v-tab value="drill">Drill</v-tab>
        </v-tabs>
        <v-tabs-window v-model="appStore.getSelectedActivity">
          <v-tabs-window-item v-for="vo in appStore.activities" :key="vo.name" :id="vo.name" :value="vo.name">
            <v-sheet id="settingsDrawer" class="d-flex flex-column px-2 justify-start" :class="{ 'open-drawer': openSettingsDrawer }">
              <v-list-item>Font Size</v-list-item>
              <div class="d-flex align-center justify-center">
                <div>lv. {{ appStore.fontSize }}</div>
                <v-btn :disabled="appStore.fontSize == 3" size="small" icon="mdi-plus-circle" @click="appStore.incrementFontSize"></v-btn>
                <v-btn :disabled="appStore.fontSize == 1" size="small" icon="mdi-minus-circle" @click="appStore.decrementFontSize"></v-btn>
              </div>
              <v-list-item>Timer Limit</v-list-item>
              <v-number-input
                v-model="appStore.testTime"
                class="align-self-center flex-grow-0"
                persistent-placeholder
                placeholder="secs"
                width="150"
                :min="15" :max="180"
                density="compact"
                variant="solo-filled"
                control-variant="split"
                hide-details
                :step="15"
                @update:model-value="appStore.startGame(appStore.testTime)"
              ></v-number-input>
            </v-sheet>
            <div v-if="openSettingsDrawer" class="menu-backdrop" @click="handleSettingsClick"></div>
            <div class="d-flex ga-6 w-100 align-start justify-space-between position-absolute pa-4"
                 style="z-index: 1; background: linear-gradient(rgb(0 0 0 / 0.6) 20%, transparent);"
            >
              <div class="d-flex ga-2 align-end">
                <v-btn append-icon="mdi-cog" @click="handleSettingsClick">settings</v-btn>
                <v-chip>{{ appStore.texts.length }}</v-chip>
              </div>
              <div class="align-self-stretch">
                <p style="color: white;">{{ vo.description }}</p>
                <v-divider></v-divider>
              </div>
              <div class="d-flex flex-column ga-2 align-end pl-8">
                <v-btn prepend-icon="mdi-shuffle" @click="pickRandomActivityText">random</v-btn>
                <v-btn prepend-icon="mdi-play" @click="appStore.startActivity">start</v-btn>
              </div>
            </div>
            <v-carousel v-model="appStore.selectedText" height="364" hide-delimiters>
              <v-carousel-item
                v-for="(vi, i) in appStore.texts" :key="i" :id="vi.title"
                :color="bgColors[i % bgColors.length]"
              >
                <div class="d-flex align-center justify-center" style="margin-top: 160px; margin-inline: 150px;">
                  <div class="d-flex flex-column ga-4">
                    <div class="d-flex ga-2">
                      <h2>{{ vi.title }}</h2>
                      <v-divider vertical></v-divider>
                      <v-chip>{{ vi.difficulty }}</v-chip>
                      <!-- <v-chip v-for="tag in vi.tags" :key="tag">{{ tag }}</v-chip> -->
                    </div>
                    <p>{{ vi.text_body.length > 200 ? vi.text_body.split(' ').slice(0, 50).join(' ') + ' ...' : vi.text_body }}</p>
                  </div>
                </div>
              </v-carousel-item>
            </v-carousel>
          </v-tabs-window-item>
        </v-tabs-window>
      </v-sheet>
    </div>
  </div>
</template>

<script setup>
  import { ref } from 'vue';
  import { useAppStore } from '@/stores/app';

  const bgColors = ['red', 'blue', 'green', 'yellow']
  const appStore = useAppStore()

  const openSettingsDrawer = ref(false)

  function sortResultsIdDescending(a, b) {
    if (a.id > b.id) { return -1 }
    else if (a.id < b.id) { return 1 }
    return 0
  }

  function pickRandomActivityText() {
    const pickIndex = Math.floor(Math.random() * appStore.texts.length)
    console.log('random pick index ->', pickIndex)
    appStore.selectedText = pickIndex
    appStore.startActivity()
  }

  function handleSettingsClick() {
    openSettingsDrawer.value = !openSettingsDrawer.value
  }
</script>
