<template>
  <div class="d-flex flex-column ga-4 fill-height main-container">
    <v-expansion-panels v-model="activeSections">
      <v-expansion-panel title="Texts Tab" value="texts-tab">
        <template v-slot:text>
          <v-text-field
            v-model="searchFilter"
            label="search"
            variant="outlined"
            density="comfortable"
          ></v-text-field>
          <v-list
            lines="three"
            select-strategy="leaf"
          >
            <v-list-item
              v-for="item in filteredTexts"
              :key="item.id"
              :title="`${item.title} [${item.tags}]`"
              :subtitle="item.text_body"
              :value="item.id"
            >
              <template #append="">
                <v-list-item-action start>
                  <v-btn icon="mdi-pencil-circle" @click="() => startEditingItem(item)"></v-btn>
                </v-list-item-action>
              </template>
            </v-list-item>
          </v-list>
          <div class="d-flex flex-column justify-end">
            <v-btn v-if="!showForm" @click="toggleForm()">Add New Text</v-btn>
            <v-form v-else ref="form" @submit.prevent="handleFormSubmit">
              <div class="py-5">{{ textInfo.id == 0 ? 'New' : 'Updating' }} Text Info</div>
              <v-select
                v-model="textInfo.text_type"
                :items="['full-text', 'drill']"
                :rules="genericTextRules"
                label="Type"
                density="compact"
                variant="outlined"
              ></v-select>
              <v-text-field
                v-model="textInfo.title"
                :rules="textTitleRules"
                label="Title"
                variant="outlined"
                density="compact"
              ></v-text-field>
              <v-select
                v-model="textInfo.difficulty"
                :items="['normal', 'easy', 'hard']"
                :rules="genericTextRules"
                label="Difficulty"
                density="compact"
                variant="outlined"
              ></v-select>
              <v-text-field
                v-model="textInfo.tags"
                :rules="genericTextRules"
                label="Tags"
                variant="outlined"
                density="compact"
              ></v-text-field>
              <v-textarea
                v-model="textInfo.text_body"
                :rules="genericTextRules"
                label="Body"
                variant="outlined"
                density="compact"
              ></v-textarea>
              <div class="d-flex justify-end">
                <v-btn @click="toggleForm()">Cancel</v-btn>
                <v-btn type="submit">{{ textInfo.id == 0 ? 'Add' : 'Udpate' }}</v-btn>
              </div>
            </v-form>
          </div>
        </template>
      </v-expansion-panel>
      <v-expansion-panel title="Users Tab" value="users-tab">
        <template v-slot:text>
          User management
        </template>
      </v-expansion-panel>
    </v-expansion-panels>
  </div>
</template>

<script setup>
  import { computed, onMounted, ref, watchEffect } from 'vue';
  import { useRouter } from 'vuetify/lib/composables/router.mjs';
  import { genericTextRules, textTitleRules } from '@/helpers/input_rules';
  import { useAppStore } from '@/stores/app';
  import { useUserStore } from '@/stores/user';

  const appStore = useAppStore()
  const userStore = useUserStore()
  const router = useRouter()

  const form = ref()
  const emptyText = ref({
    id: 0,
    text_type: '',
    title: '',
    difficulty: '',
    tags: [],
    text_body: '',
  })
  const textInfo = ref(emptyText.value)
  const showForm = ref(false)
  const searchFilter = ref('')
  const activeSections = ref(['texts-tab'])

  const filteredTexts = computed(() => appStore.texts?.filter((t) => t.title.includes(searchFilter.value)))

  watchEffect(() => {if (!userStore.isLoggedIn && !userStore.isAdmin) router.push('/')})

  function startEditingItem(item) {
    textInfo.value = item
    showForm.value = true
  }

  function toggleForm() {
    if (showForm.value) {
      textInfo.value = emptyText.value
    }
    showForm.value = !showForm.value
  }

  function filterItems() {
    pass
  }

  async function handleFormSubmit() {
    if (form.value.isValid) {
      const resp = await appStore.createOrUpdateText(textInfo.value)
      if (resp !== 0) {
        toggleForm()
      }
    }
  }
</script>
