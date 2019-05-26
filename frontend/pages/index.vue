<template>
  <h1>Adventar</h1>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";

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
export default class extends Vue {}
</script>

<style></style>
