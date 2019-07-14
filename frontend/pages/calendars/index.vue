<template>
  <div>
    <GlobalHeader />

    <PageHeader>{{ year }}年のAdvent Calnedar</PageHeader>

    <main>
      <CalendarSearchForm :year="year" :defaultQuery="query" style="margin-bottom: 30px" />
      <CalendarList :calendars="calendars" />
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { listCalendars } from "~/lib/GrpcClient";
import { Calendar } from "~/types/adventar";
import GlobalHeader from "~/components/GlobalHeader.vue";
import PageHeader from "~/components/PageHeader.vue";
import CalendarSearchForm from "~/components/CalendarSearchForm.vue";
import CalendarList from "~/components/CalendarList.vue";

@Component({
  components: { GlobalHeader, PageHeader, CalendarSearchForm, CalendarList }
})
export default class extends Vue {
  year: number;
  query: string;
  calendars: Calendar[] = [];

  async created() {
    this.year = Number(this.$route.query.year);
    this.query = String(this.$route.query.query || "");
    this.calendars = await listCalendars({ year: this.year, query: this.query });
  }
}
</script>
