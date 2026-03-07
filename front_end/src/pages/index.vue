<template>
  <div class="d-flex flex-column ga-4 fill-height main-container">
    <div class="d-flex justify-center align-center bg-red histogram-container">
      <div>
        WIP
      </div>
    </div>
    <div class="d-flex justify-center align-center bg-red select-container">
      <v-sheet class="h-100 w-100">
        <v-tabs v-model="selectedActivity" color="primary" density="compact">
          <v-tab value="typingTest">Typing Test</v-tab>
          <v-tab value="drill">Drill</v-tab>
        </v-tabs>

        <v-tabs-window v-model="selectedActivity">
          <v-tabs-window-item v-for="(v, k) in activities" :key="k" :id="k" :value="k">
            <div class="d-flex ga-6 w-100 align-start justify-space-between position-absolute pa-4"
                 style="z-index: 1; background: linear-gradient(rgb(0 0 0 / 0.6) 20%, transparent);"
            >
              <div class="d-flex ga-2 align-end">
                <v-btn append-icon="mdi-cog">settings</v-btn>
                <v-chip>{{ v.items.length }}</v-chip>
              </div>
              <p style="color: white;">{{ v.description }}</p>
              <div class="d-flex flex-column ga-2 align-end pl-8">
                <v-btn prepend-icon="mdi-shuffle">random</v-btn>
                <v-btn prepend-icon="mdi-play">start</v-btn>
              </div>
            </div>
            <v-carousel height="364" hide-delimiters>
              <v-carousel-item
                v-for="(vi, i) in activities[selectedActivity].items" :key="i" :id="i"
                :color="bgColors[i%bgColors.length]"
              >
                <div class="d-flex align-center justify-center" style="margin-top: 140px; margin-inline: 150px;">
                  <div class="d-flex flex-column ga-4">
                    <div class="d-flex ga-2">
                      <h2>{{ vi.title }}</h2>
                      <v-divider vertical></v-divider>
                      <v-chip>{{ vi.difficulty }}</v-chip>
                      <v-chip v-for="tag in vi.tags" :key="tag">{{ tag }}</v-chip>
                    </div>
                    <p>{{ vi.text.repeat(5) }}</p>
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

  const bgColors = ['red', 'blue', 'green', 'yellow']
  const activities = ref({
    'typingTest': {
      description: 'type the text as fast as you can under the time limit',
      settings: {},
      items: [
        {
          title: 'typing-place-holder',
          difficulty: 'normal',
          text: 'some lorem ipsum testing placeholder',
          tags: ['testing'],
        },
        {
          title: 'typing-place-holder-2',
          difficulty: 'hard',
          text: 'some lorem ipsum testing placeholder',
          tags: ['testing'],
        },
        {
          title: 'typing-place-holder-3',
          difficulty: 'hard',
          text: 'some lorem ipsum testing placeholder',
          tags: ['testing'],
        },
        {
          title: 'typing-place-holder-4',
          difficulty: 'hard',
          text: 'some lorem ipsum testing placeholder',
          tags: ['testing'],
        },
        {
          title: 'typing-place-holder-5',
          difficulty: 'hard',
          text: 'some lorem ipsum testing placeholder',
          tags: ['testing'],
        },
      ],
    },
    'drill': {
      description: 'hit the keys as the appear on screen to practice your accuracy',
      settings: {},
      items: [
        {
          title: 'drill-place-holder',
          difficulty: 'normal',
          text: 'some lorem ipsum testing placeholder',
          tags: ['testing'],
        },
      ],
    },
  })
  const selectedActivity = ref('typingTest')

</script>
