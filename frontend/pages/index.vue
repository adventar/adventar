<template>
  <div>
    <advHeader />

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
import advHeader from "~/components/header.vue";

@Component({
  components: { advHeader }
})
export default class extends Vue {
  async asyncData() {
    const calendars = await listCalendar();
    return { calendars };
  }
}
</script>
