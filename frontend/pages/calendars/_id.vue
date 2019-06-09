<template>
  <div>
    <advHeader />
    <div v-if="calendar">
      <h2>{{ calendar.title }} Advent Calendar {{ calendar.year }}</h2>
      <p>{{ calendar.description }}</p>
      <hr />
      <ul>
        <li v-for="item in listDays()" :key="item.day">
          <span>{{ calendar.year }}/12/{{ item.day }}</span>
          <span v-if="!item.entry"><button @click="onClickCreateEntry(item.day)">登録</button></span>
          <span v-if="item.entry">
            <span>{{ item.entry.owner.name }}</span>
            <button v-if="cancelable(item.entry)" @click="onClickDeleteEntry(item.entry.id)">キャンセル</button>
          </span>
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { getCalendar, createEntry, deleteEntry } from "~/lib/GrpcClient";
import { Calendar, Entry } from "~/types/adventar";
import { getToken } from "~/plugins/auth";
import advHeader from "~/components/header.vue";

@Component({
  components: { advHeader }
})
export default class extends Vue {
  calendar: Calendar;

  async asyncData({ params }) {
    const calendar = await getCalendar(Number(params.id));
    const days = Array.from({ length: 25 }, (_, k) => k + 1);
    return { calendar, days };
  }

  async onClickCreateEntry(day): Promise<void> {
    const calendarId = this.calendar.id;
    const token = await getToken();
    if (!token) {
      throw new Error("Token is null");
    }
    await createEntry({ calendarId, day, token });
    this.calendar = await getCalendar(this.calendar.id);
  }

  async onClickDeleteEntry(entryId): Promise<void> {
    const token = await getToken();
    if (!token) {
      throw new Error("Token is null");
    }
    await deleteEntry({ entryId, token });
    this.calendar = await getCalendar(this.calendar.id);
  }

  listDays(): { day: number; entry: Entry }[] {
    const entryMapByDay: { [key: number]: Entry } = {};
    if (this.calendar !== null && this.calendar.entries !== undefined) {
      this.calendar.entries.forEach(entry => {
        if (entry.day) {
          entryMapByDay[entry.day] = entry;
        }
      });
    }

    const items: { day: number; entry: Entry }[] = [];
    for (let day = 1; day <= 25; day++) {
      const entry = entryMapByDay[day];
      items.push({ day, entry });
    }
    return items;
  }

  cancelable(entry: Entry): boolean {
    // TODO: Calendar owner can cancel entry
    return entry.owner && this.$store.state.user && entry.owner.id === this.$store.state.user.id;
  }
}
</script>
