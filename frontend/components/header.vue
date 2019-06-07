<template>
  <header>
    <h1><nuxt-link to="/">Adventar</nuxt-link></h1>
    <div v-if="$store.state.user === null">
      <ul>
        <li><button @click="login('google')">Google でログイン</button></li>
        <li><button @click="login('github')">GitHub でログイン</button></li>
        <li><button @click="login('twitter')">Twitter でログイン</button></li>
        <li><button @click="login('facebook')">Facebook でログイン</button></li>
      </ul>
    </div>
    <div v-else-if="$store.state.user">
      <div>id: {{ $store.state.user.id }}</div>
      <div v-if="!editmode">name: {{ $store.state.user.name }} <span @click="editmode = true">Edit</span></div>
      <div v-if="editmode">name: <input type="text" :value="$store.state.user.name" @change="onChangeName" /></div>
      <img :src="$store.state.user.iconUrl" width="50" height="50" />
      <button @click="logout()">Logout</button>
      <nuxt-link to="/calendars/new">カレンダーを作る</nuxt-link>
    </div>
  </header>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { updateUser } from "~/lib/grpc/Client";
import { getToken } from "~/plugins/firebase";
import { loginWithFirebase, logoutWithFirebase } from "~/plugins/firebase";

@Component
export default class extends Vue {
  editmode = false;

  async onChangeName(e) {
    const token = await getToken();
    if (token === null) {
      throw new Error("Token is null");
    }
    const user = await updateUser(e.target.value, token);
    this.$store.commit("setUser", user);
    this.editmode = false;
  }

  login(provider) {
    loginWithFirebase(provider);
  }

  logout() {
    logoutWithFirebase();
  }
}
</script>
