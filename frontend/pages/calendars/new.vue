<template>
  <div>
    <GlobalHeader />

    <PageHeader>カレンダーを作る</PageHeader>

    <main>
      <section>
        <SectionHeader>タイトル</SectionHeader>
        <input v-model="title" type="text" /> Advent Calendar {{ currentYear }}
      </section>

      <section>
        <SectionHeader>概要</SectionHeader>
        <textarea v-model="description"></textarea>
      </section>

      <button @click="submit()">作成</button>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { createCalendar } from "~/lib/GrpcClient";
import { getCurrentYear } from "~/lib/Configuration";
import { getToken } from "~/plugins/auth";
import GlobalHeader from "~/components/GlobalHeader.vue";
import PageHeader from "~/components/PageHeader.vue";
import SectionHeader from "~/components/SectionHeader.vue";

@Component({
  components: { GlobalHeader, PageHeader, SectionHeader }
})
export default class extends Vue {
  title: string = "";
  description: string = "";
  currentYear: number = getCurrentYear();

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
