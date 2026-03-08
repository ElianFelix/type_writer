<template>
  <div class="d-flex flex-column ga-4 fill-height main-container">
    <div class="d-flex justify-center align-center bg-red histogram-container">
      <div>
        WIP
      </div>
    </div>
    <div class="d-flex justify-center align-center bg-red select-container">
      <v-sheet class="h-100 w-100">
        <v-tabs v-model="appStore.selectedActivity" color="primary" density="compact">
          <v-tab value="typingTest">Typing Test</v-tab>
          <v-tab value="drill">Drill</v-tab>
        </v-tabs>

        <v-tabs-window v-model="appStore.getSelectedActivity">
          <v-tabs-window-item v-for="(vo, k) in appStore.activities" :key="k" :id="k" :value="k">
            <div class="d-flex ga-6 w-100 align-start justify-space-between position-absolute pa-4"
                 style="z-index: 1; background: linear-gradient(rgb(0 0 0 / 0.6) 20%, transparent);"
            >
              <div class="d-flex ga-2 align-end">
                <v-btn append-icon="mdi-cog">settings</v-btn>
                <v-chip>{{ vo.items.length }}</v-chip>
              </div>
              <div class="align-self-stretch">
                <p style="color: white;">{{ vo.description }}</p>
                <v-divider></v-divider>
              </div>
              <div class="d-flex flex-column ga-2 align-end pl-8">
                <v-btn prepend-icon="mdi-shuffle">random</v-btn>
                <v-btn prepend-icon="mdi-play" @click="appStore.startActivity">start</v-btn>
              </div>
            </div>
            <v-carousel v-model="appStore.selectedText" height="364" hide-delimiters>
              <v-carousel-item
                v-for="(vi, i) in appStore.activities[appStore.selectedActivity].items" :key="i" :id="vi.title"
                :color="bgColors[i % bgColors.length]"
              >
                <div class="d-flex align-center justify-center" style="margin-top: 160px; margin-inline: 150px;">
                  <div class="d-flex flex-column ga-4">
                    <div class="d-flex ga-2">
                      <h2>{{ vi.title }}</h2>
                      <v-divider vertical></v-divider>
                      <v-chip>{{ vi.difficulty }}</v-chip>
                      <v-chip v-for="tag in vi.tags" :key="tag">{{ tag }}</v-chip>
                    </div>
                    <p>{{ vi.text.length > 200 ? vi.text.split(' ').slice(0, 50).join(' ') + ' ...' : vi.text }}</p>
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
  import { useAppStore } from '@/stores/app';

  const appStore = useAppStore()

  const bgColors = ['red', 'blue', 'green', 'yellow']
</script>
