<template>
  <div class="d-flex flex-column ga-4 main-container">
    <v-expansion-panels v-model="activeSections">
      <TextPanel v-model="appStore.texts" />
      <v-expansion-panel title="Users Tab" value="users-tab">
        <template #text>
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

  const activeSections = ref(['texts-tab'])

  watchEffect(() => {if (!userStore.isLoggedIn && !userStore.isAdmin) router.push('/')})
</script>
