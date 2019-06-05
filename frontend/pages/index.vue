<template>
  <div>
    <h1><nuxt-link to="/">Adventar</nuxt-link></h1>
    <div v-if="user !== null">
      <div>{{ user.name }}</div>
    </div>
    <nuxt-link to="/login">Login</nuxt-link>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { initializeFirebaseApp, auth } from "~/lib/Firebase";
import { User } from "~/lib/grpc/Client";

initializeFirebaseApp();

@Component
export default class extends Vue {
  user: User | null = null;

  async mounted() {
    this.user = await auth();
  }
}
</script>

<style></style>
