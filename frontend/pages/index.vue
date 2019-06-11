<template>
  <div>
    <GlobalHeader />

    <PageHeader>{{ currentYear }}年のAdvent Calnedar</PageHeader>

    <ul>
      <li v-for="calendar in calendars" :key="calendar.id">
        <nuxt-link :to="`/calendars/${calendar.id}`">{{ calendar.title }}</nuxt-link>
        <span>{{ calendar.entryCount }}/25</span>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { listCalendar } from "~/lib/GrpcClient";
import { getCurrentYear } from "~/lib/Configuration";
import GlobalHeader from "~/components/GlobalHeader.vue";
import PageHeader from "~/components/PageHeader.vue";

@Component({
  components: { GlobalHeader, PageHeader }
})
export default class extends Vue {
  currentYear = getCurrentYear();

  async asyncData() {
    const calendars = await listCalendar();
    return { calendars };
  }
}
</script>
