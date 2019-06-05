<template>
  <div>
    <h1><nuxt-link to="/">Adventar</nuxt-link></h1>
    <h2>Login</h2>
    <ul>
      <li><button @click="login('google')">Google でログイン</button></li>
      <li><button @click="login('github')">GitHub でログイン</button></li>
      <li><button @click="login('twitter')">Twitter でログイン</button></li>
      <li><button @click="login('facebook')">Facebook でログイン</button></li>
    </ul>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { initializeFirebaseApp, loginWithFirebase } from "~/lib/Firebase";

initializeFirebaseApp();

@Component
export default class extends Vue {
  login(provider) {
    loginWithFirebase(provider);
  }
}
</script>
