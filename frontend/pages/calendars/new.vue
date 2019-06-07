<template>
  <div>
    <advHeader />

    <h2>カレンダーを作る</h2>

    <h3>タイトル</h3>
    <input v-model="title" type="text" /> Advent Calendar 20xx

    <h3>概要</h3>
    <textarea v-model="description"></textarea>

    <button @click="submit()">作成</button>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { createCalendar } from "~/lib/grpc/Client";
import { getToken } from "~/plugins/firebase";
import advHeader from "~/components/header.vue";

@Component({
  components: { advHeader }
})
export default class extends Vue {
  title = "";
  description = "";

  async submit() {
    const token = await getToken();
    if (!token) {
      throw new Error("Token is null");
    }
    const calendarId = await createCalendar({
      title: this.title,
      description: this.description,
      token: token
    });
    this.$router.push(`/calendars/${calendarId}`);
  }
}
</script>
