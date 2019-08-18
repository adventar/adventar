<template>
  <div>
    <GlobalHeader />

    <PageHeader>カレンダーを作る</PageHeader>

    <main>
      <div>
        <form @submit.prevent="submit()">
          <section>
            <SectionHeader>タイトル</SectionHeader>
            <input v-model="title" type="text" /> Advent Calendar {{ currentYear }}
          </section>

          <section>
            <SectionHeader>概要</SectionHeader>
            <textarea v-model="description"></textarea>
            <p class="note">Markdown記法が使えます。</p>
          </section>

          <button type="submit">作成</button>
        </form>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { createCalendar } from "~/lib/GrpcClient";
import { getCurrentYear } from "~/lib/Configuration";
import { getToken } from "~/lib/Auth";
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
    const calendarId = await createCalendar({
      title: this.title,
      description: this.description,
      token: token
    });
    this.$router.push(`/calendars/${calendarId}`);
  }
}
</script>

<style scoped>
.note {
  margin-top: 10px;
  font-size: 13px;
  color: #666;
}

input[type="text"] {
  font-size: 16px;
  width: 300px;
  padding: 5px;
  border: 1px solid #ccc;
}

textarea {
  width: 500px;
  height: 200px;
  font-size: 14px;
  padding: 5px;
  border: 1px solid #ccc;
}

button[type="submit"] {
  font-size: 16px;
  color: #fff;
  border: 2px solid #28a745;
  background-color: #28a745;
  padding: 8px 30px;
  border-radius: 10px;
  cursor: pointer;
}
</style>
