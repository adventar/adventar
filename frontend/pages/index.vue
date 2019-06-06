<template>
  <div>
    <h1><nuxt-link to="/">Adventar</nuxt-link></h1>
    <div v-if="$store.state.user !== null">
      <div>id: {{ $store.state.user.id }}</div>
      <div v-if="!editmode">name: {{ $store.state.user.name }} <span @click="editmode = true">Edit</span></div>
      <div v-if="editmode">name: <input type="text" :value="$store.state.user.name" @change="onChangeName" /></div>
      <img :src="$store.state.user.iconUrl" width="50" height="50" />
    </div>
    <nuxt-link to="/login">Login</nuxt-link>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { updateUser } from "~/lib/grpc/Client";
import { getToken } from "~/plugins/firebase";

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
}
</script>

<style></style>
