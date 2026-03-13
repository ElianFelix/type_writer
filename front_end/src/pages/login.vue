<template>
  <div class="d-flex flex-column ga-4 fill-height main-container">
    <div class="d-flex justify-center align-center bg-red histogram-container">
      <v-sheet class="d-flex flex-column ga-4" width="280">
        <v-form @submit.prevent>
          <v-text-field
            v-model="username"
            :rules="rules"
            label="Username"
            variant="outlined"
          ></v-text-field>
          <v-text-field
            v-model="password"
            :type="false ? 'text' : 'password'"
            hint="At least 8 characters"
            :rules="rules"
            label="Password"
            variant="outlined"
          ></v-text-field>
          <div class="d-flex justify-end">
            <v-btn class="mt-2">Sign Up</v-btn>
            <v-btn class="mt-2" type="submit">Submit</v-btn>
          </div>
        </v-form>
      </v-sheet>
    </div>
  </div>
</template>

<script setup>
</script>
