<template>
  <div>
    <GlobalHeader />

    <PageHeader>カレンダー編集</PageHeader>

    <main>
      <div v-if="calendar">
        <CalendarForm :calendar="calendar" />
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { getCalendar } from "~/lib/GrpcClient";
import { Calendar } from "~/types/adventar";
import GlobalHeader from "~/components/GlobalHeader.vue";
import PageHeader from "~/components/PageHeader.vue";
import CalendarForm from "~/components/CalendarForm.vue";

@Component({
  components: { GlobalHeader, PageHeader, CalendarForm }
})
export default class extends Vue {
  calendar: Calendar | null = null;

  async created() {
    // TODO: 404 if not found
    this.calendar = await getCalendar(Number(this.$route.params.id));
  }
}
</script>
