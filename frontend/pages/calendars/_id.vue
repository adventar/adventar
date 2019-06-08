<template>
  <div>
    <advHeader />
    <div v-if="calendar">
      <h2>{{ calendar.title }} Advent Calendar {{ calendar.year }}</h2>
      <p>{{ calendar.description }}</p>
      <hr />
      <ul>
        <li v-for="day in days" :key="day">{{ calendar.year }}/12/{{ day }}</li>
      </ul>
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
    const days = Array.from({ length: 25 }, (_, k) => k + 1);
    return { calendar, days };
  }
}
</script>
