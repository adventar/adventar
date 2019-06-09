<template>
  <header>
    <h1><nuxt-link to="/">Adventar</nuxt-link></h1>
    <div v-if="$store.state.user">
      <img :src="$store.state.user.iconUrl" width="50" height="50" />
      <button @click="logout()">Logout</button>
      <nuxt-link to="/calendars/new">カレンダーを作る</nuxt-link>
      <nuxt-link to="/setting">設定</nuxt-link>
    </div>
    <div v-else>
      <ul>
        <li><button @click="login('google')">Google でログイン</button></li>
        <li><button @click="login('github')">GitHub でログイン</button></li>
        <li><button @click="login('twitter')">Twitter でログイン</button></li>
        <li><button @click="login('facebook')">Facebook でログイン</button></li>
      </ul>
    </div>
  </header>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { loginWithFirebase, logoutWithFirebase } from "~/plugins/firebase";

@Component
export default class extends Vue {
  login(provider) {
    loginWithFirebase(provider);
  }

  logout() {
    this.$router.push("/");
    logoutWithFirebase();
  }
}
</script>
