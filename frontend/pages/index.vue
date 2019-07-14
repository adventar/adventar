<template>
  <div>
    <GlobalHeader />

    <PageHeader>{{ currentYear }}年のAdvent Calnedar</PageHeader>

    <main>
      <CalendarSearchForm @submit="handleSubmit" style="margin-bottom: 30px" />
      <CalendarList :calendars="calendars" />
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { listCalendars } from "~/lib/GrpcClient";
import { getCurrentYear } from "~/lib/Configuration";
import { Calendar } from "~/types/adventar";
import GlobalHeader from "~/components/GlobalHeader.vue";
import PageHeader from "~/components/PageHeader.vue";
import CalendarSearchForm from "~/components/CalendarSearchForm.vue";
import CalendarList from "~/components/CalendarList.vue";

@Component({
  components: { GlobalHeader, PageHeader, CalendarSearchForm, CalendarList }
})
export default class extends Vue {
  currentYear = getCurrentYear();
  calendars: Calendar[] = [];

  async mounted() {
    const pageSize = 20;
    const calendars = await listCalendars({ year: this.currentYear, pageSize });
    this.calendars = calendars;
  }

  async handleSubmit(query) {
    this.$router.push(`/calendars?year=${this.currentYear}&query=${query}`);
  }
}
</script>
