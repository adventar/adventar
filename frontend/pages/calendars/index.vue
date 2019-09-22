<template>
  <div>
    <PageHeader>{{ year }}年のAdvent Calnedar</PageHeader>

    <main>
      <div class="mainInner">
        <CalendarSearchForm :defaultQuery="query" @submit="handleSubmit" />
        <CalendarList :calendars="calendars" />
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { listCalendars } from "~/lib/GrpcClient";
import { Calendar } from "~/types/adventar";
import PageHeader from "~/components/PageHeader.vue";
import CalendarSearchForm from "~/components/CalendarSearchForm.vue";
import CalendarList from "~/components/CalendarList.vue";

@Component({
  components: { PageHeader, CalendarSearchForm, CalendarList }
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

  async handleSubmit(query) {
    this.query = query;
    this.calendars = [];
    this.calendars = await listCalendars({ year: this.year, query: this.query });
    this.$router.push(`/calendars?year=${this.year}&query=${query}`);
  }
}
</script>

<style lang="scss" scoped>
.mainInner {
  padding-top: 15px;
}

.CalendarSearchForm {
  margin-bottom: 15px;
}

@media (min-width: $mq-break-small) {
  .mainInner {
    padding-top: 30px;
  }

  .CalendarSearchForm {
    margin-bottom: 30px;
  }
}
</style>
