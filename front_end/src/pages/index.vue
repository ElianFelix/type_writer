<template>
  <div class="d-flex flex-column ga-4 fill-height main-container">
    <div class="d-flex justify-center align-center bg-red histogram-container">
      <div>
        WIP
      </div>
    </div>
    <div class="d-flex justify-center align-center bg-red select-container">
      <v-sheet class="h-100 w-100">
        <v-tabs v-model="activityType" color="primary">
          <v-tab value="one">Typing Test</v-tab>
          <v-tab value="two">Drill</v-tab>
        </v-tabs>

        <v-tabs-window v-model="activityType">
          <v-tabs-window-item value="one">
            <div class="d-flex w-100 align-center justify-space-between position-absolute" style="z-index: 1;">
              <v-btn>filters</v-btn>
              <h3>description of game</h3>
              <div class="d-flex flex-column align-center">
                <v-btn>random</v-btn>
                <v-btn>start</v-btn>
              </div>
            </div>
            <v-carousel height="352" hide-delimiters>
              <v-carousel-item
                  color="red"
              >
                <div class="d-flex flex-column ga-4 align-center justify-center" style="margin-top: 140px;">
                  <div class="d-flex ga-2">
                    <h2>title</h2>
                    <v-chip>difficulty</v-chip>
                  </div>
                  <p>sample of highlighted text, probably up to the first 2 lines.</p>
                </div>
              </v-carousel-item>
              <v-carousel-item
                  color="blue"
              ></v-carousel-item>
              <v-carousel-item
                  color="green"
              ></v-carousel-item>
            </v-carousel>
          </v-tabs-window-item>
          <v-tabs-window-item value="two">
            <v-sheet class="pa-5" color="orange">Two</v-sheet>
          </v-tabs-window-item>
          <v-tabs-window-item value="three">
            <v-sheet class="pa-5" color="brown">Three</v-sheet>
          </v-tabs-window-item>
        </v-tabs-window>
      </v-sheet>
    </div>
  </div>
</template>

<script setup>
  import { ref } from 'vue';



  const activityType = ref("one")
</script>
