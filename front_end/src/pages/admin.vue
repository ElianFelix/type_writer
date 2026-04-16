<template>
  <div class="d-flex flex-column ga-4 fill-height main-container">
    <v-expansion-panels v-model="activeSections">
      <v-expansion-panel title="Texts Tab" value="texts-tab">
        <template v-slot:text>
          Text management
          <div>filter bar</div>
          <v-text-field label="search/filter" variant="outlined"></v-text-field>
          <div>list</div>
          <v-list
            v-model:selected="settingsSelection"
            lines="three"
            select-strategy="leaf"
          >
            <v-list-item
              v-for="item in appStore.texts"
              :key="item.id"
              :title="`${item.title} [${item.tags}]`"
              :subtitle="item.text_body"
              :value="item.id"
            >
              <template #append="">
                <v-list-item-action start>
                  <v-btn icon="mdi-pencil-circle" @click="startEditingItem(item)"></v-btn>
                </v-list-item-action>
              </template>
            </v-list-item>
          </v-list>
          <div>add new</div>
          <v-btn @click="() => showForm = !showForm">{{ showForm ? 'Cancel' : 'Add Text' }}</v-btn>
          <v-form v-if="showForm" ref="form" @submit.prevent="handleFormSubmit">
            <v-text-field
              v-model="textInfo.text_type"
              rules=""
              label="Type"
              variant="outlined"
              density="compact"
            ></v-text-field>
            <v-text-field
              v-model="textInfo.title"
              rules=""
              label="Title"
              variant="outlined"
              density="compact"
            ></v-text-field>
            <v-text-field
              v-model="textInfo.difficulty"
              rules=""
              label="Difficulty"
              variant="outlined"
              density="compact"
            ></v-text-field>
            <v-text-field
              v-model="textInfo.tags"
              rules=""
              label="Tags"
              variant="outlined"
              density="compact"
            ></v-text-field>
            <v-text-field
              v-model="textInfo.text_body"
              rules=""
              label="Body"
              variant="outlined"
              density="compact"
            ></v-text-field>
            <div class="d-flex justify-end">
              <v-btn type="submit">Add</v-btn>
            </div>
          </v-form>
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
  import { onMounted, ref, watchEffect } from 'vue';
  import { useRouter } from 'vuetify/lib/composables/router.mjs';
  import { useAppStore } from '@/stores/app';
  import { useUserStore } from '@/stores/user';

  const appStore = useAppStore()
  const userStore = useUserStore()
  const router = useRouter()

  const form = ref()
  const textInfo = ref({
    id: 0,
    text_type: '',
    title: '',
    difficulty: '',
    tags: [],
    text_body: '',
  })
  const showForm = ref(false)

  const activeSections = ref(['texts-tab'])

  watchEffect(() => {if (!userStore.isLoggedIn && !userStore.isAdmin) router.push('/')})

  function startEditingItem(item) {
    textInfo.value = item
    showForm.value = true
  }

  async function handleFormSubmit() {
    if (true) return
    if (form.value.isValid) {
      await userStore.updateUserDetails(textInfo.value)
    }
  }
</script>
