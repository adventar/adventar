<template>
  <div>
    <GlobalHeader />

    <PageHeader>ユーザー名</PageHeader>

    <main>
      <SectionHeader>作成したカレンダー</SectionHeader>
      <SectionHeader>登録したカレンダー</SectionHeader>
      <SectionHeader>iCal</SectionHeader>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { User } from "~/types/adventar";
import GlobalHeader from "~/components/GlobalHeader.vue";
import PageHeader from "~/components/PageHeader.vue";
import SectionHeader from "~/components/SectionHeader.vue";
import CalendarList from "~/components/CalendarList.vue";

@Component({
  components: { GlobalHeader, PageHeader, SectionHeader, CalendarList }
})
export default class extends Vue {
  user: User;
  async asyncData({ params }) {
    const userId = params.id;
    console.log(userId);
  }
}
</script>
