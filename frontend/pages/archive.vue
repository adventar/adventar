<template>
  <div>
    <PageHeader>これまでの Adventar Calendar</PageHeader>
    <main>
      <div class="mainInner">
        <ul class="calendars">
          <li v-for="stat in stats" :key="stat.year" class="item">
            <nuxt-link class="title" :to="`/calendars?year=${stat.year}`">
              <font-awesome-icon :icon="['far', 'calendar']" />
              {{ stat.year }}年
            </nuxt-link>
            <span class="label">
              カレンダー<strong>{{ stat.calendarsCount }}</strong>
            </span>
            <span class="label">
              参加者<strong>{{ stat.entriesCount }}</strong>
            </span>
          </li>
        </ul>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { listCalendarStats } from "~/lib/GrpcClient";
import { CalendarStat } from "~/types/adventar";
import PageHeader from "~/components/PageHeader.vue";
@Component({
  components: { PageHeader }
})
export default class extends Vue {
  stats: CalendarStat[] = [];

  async mounted() {
    this.stats = await listCalendarStats();
  }
}
</script>

<style lang="scss" scoped>
.calendars {
  margin: 0;
  padding: 0;
  list-style: none;
}

.item {
  font-size: 16px;
  margin: 7px 0px 15px;
}

.title {
  margin-bottom: 5px;
}

.label {
  display: inline-block;
  background: #ddd;
  border-radius: 3px;
  font-size: 12px;
  padding: 3px 5px;
  margin: 0 0 0 3px;
}
</style>
