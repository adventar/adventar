<template>
  <div>
    <h1>Adventar</h1>

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
import { initFirebase, loginWithFirebase } from "~/lib/Firebase";

initFirebase();

// import { GetCalendarRequest } from "~/lib/grpc/adventar/v1/adventar_pb";
// import { AdventarClient } from "~/lib/grpc/adventar/v1/adventar_grpc_web_pb";
// const client = new AdventarClient("http://localhost:8000", null, null);
// const request = new GetCalendarRequest();
// request.setCalendarId(2);
// client.getCalendar(request, {}, (err, res) => {
//   console.log(err);
//   console.log(res.getCalendar().getTitle());
// });

@Component
export default class extends Vue {
  login(provider) {
    loginWithFirebase(provider);
  }
}
</script>

<style></style>
