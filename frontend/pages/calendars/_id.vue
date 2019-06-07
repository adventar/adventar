<template>
  <div>
    <advHeader />
    <div v-if="calendar">
      <h2>{{ calendar.title }}</h2>
      <h2>{{ calendar.description }}</h2>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { getCalendar } from "~/lib/grpc/Client";
import advHeader from "~/components/header.vue";

@Component({
  components: { advHeader }
})
export default class extends Vue {
  async asyncData({ params }) {
    const calendar = await getCalendar(Number(params.id));
    return {
      calendar
    };
  }
}
</script>
