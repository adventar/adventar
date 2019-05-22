<template>
  <section class="container">
    <div>
      <logo />
      <h1 class="title">adventar</h1>
      <h2 class="subtitle">Adventar frontend</h2>
      <div class="links">
        <a href="https://nuxtjs.org/" target="_blank" class="button--green">Documentation</a>
        <a href="https://github.com/nuxt/nuxt.js" target="_blank" class="button--grey">GitHub</a>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import Logo from "~/components/Logo.vue";

import { GetCalendarRequest } from "~/lib/grpc/adventar/v1/adventar_pb";
import { AdventarClient } from "~/lib/grpc/adventar/v1/adventar_grpc_web_pb";
const client = new AdventarClient("http://localhost:8000", null, null);
const request = new GetCalendarRequest();
request.setCalendarId(3);
client.getCalendar(request, {}, (err, res) => {
  console.log(err);
  console.log(res.getCalendar().getTitle());
});

@Component({
  components: {
    Logo
  }
})
export default class extends Vue {}
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue",
    Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
