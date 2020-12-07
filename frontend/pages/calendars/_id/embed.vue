<template>
  <div>
    <template v-if="calendar">
      <header class="header" :style="{ backgroundColor: calendarColor }">
        <div class="inner">
          <h2 class="title">
            <a :href="calendarUrl()" target="_blank" rel="noopener">
              {{ title }}
              <font-awesome-icon icon="external-link-alt" />
            </a>
          </h2>
        </div>
      </header>
      <main>
        <div class="mainInner">
          <div class="CalendarTable">
            <table class="table">
              <thead>
                <tr>
                  <th><div class="inner">SUN</div></th>
                  <th><div class="inner">MON</div></th>
                  <th><div class="inner">TUE</div></th>
                  <th><div class="inner">WED</div></th>
                  <th><div class="inner">THU</div></th>
                  <th><div class="inner">FRI</div></th>
                  <th><div class="inner">SAT</div></th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(cells, i) in rows" :key="i">
                  <td
                    v-for="(cell, idx) in cells"
                    :key="cell.day"
                    :ref="`cell-${cell.day}`"
                    class="cell"
                    :class="{
                      'is-posted': cell.entry && cell.entry.url && !isFutureEntry(cell.entry)
                    }"
                    :data-idx="idx"
                  >
                    <div v-if="cell.entryable" class="inner">
                      <span class="day">{{ cell.day }}</span>
                      <div v-if="cell.entry">
                        <span class="entryUser">
                          <UserIcon :user="cell.entry.owner" size="24" />
                          <div class="userName">{{ cell.entry.owner.name }}</div>
                        </span>
                      </div>
                      <div v-else class="entryAction">
                        <a :href="calendarUrl()" target="_blank" rel="noopener">登録</a>
                      </div>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </main>
    </template>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import * as JsonApiClient from "~/lib/JsonApiClient";
import { calendarColor } from "~/lib/utils/Colors";
import { Calendar, Entry } from "~/types/adventar";
import UserIcon from "~/components/UserIcon.vue";
import dayjs from "dayjs";
import { getToday } from "~/lib/Configuration";

@Component({
  components: { UserIcon }
})
export default class extends Vue {
  calendar: Calendar | null = null;

  layout() {
    return "embed";
  }

  async asyncData({ params, error }) {
    const meta = document.createElement("meta");
    meta.setAttribute("name", "viewport");
    meta.setAttribute(
      "content",
      "height=device-height, width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no, target-densitydpi=device-dpi"
    );
    document.head.appendChild(meta);
    let calendar: Calendar;
    try {
      calendar = await JsonApiClient.getCalendar(params.id);
    } catch (err) {
      if (err.response) {
        error({ statusCode: err.response.status });
      } else {
        console.error(err);
        error({ statusCode: 500 });
      }

      return;
    }
    return { calendar };
  }

  get calendarColor(): string {
    return calendarColor(this.calendar!.id);
  }

  get title(): string {
    if (this.calendar === null) {
      return "Adventar";
    }
    return `${this.calendar.title} Advent Calendar ${this.calendar.year}`;
  }

  get rows() {
    if (!this.calendar) return [];
    const year = this.calendar.year;
    const endDay = dayjs(new Date(year, 11, 25));
    let currentDay = dayjs(new Date(year, 11, 1)).startOf("week");

    const entryMapByDay: { [key: number]: Entry } = {};
    if (this.calendar !== null && this.calendar.entries !== undefined) {
      this.calendar.entries.forEach(entry => {
        if (entry.day) {
          entryMapByDay[entry.day] = entry;
        }
      });
    }

    const rows: any[] = [];
    while (currentDay <= endDay) {
      const cells: any[] = [];
      for (let i = 0; i < 7; i++) {
        const isDecember = currentDay.month() === 11;
        const day = currentDay.date();
        const entry = (isDecember && entryMapByDay[day]) || null;
        const entryable = isDecember && day <= 25;
        const cell = { day, entry, entryable };
        cells.push(cell);
        currentDay = currentDay.add(1, "day");
      }
      rows.push(cells);
    }

    return rows;
  }

  isFutureEntry(entry): boolean {
    // TODO: Fix to JST
    if (!this.calendar) return false;
    return new Date(this.calendar.year, 11, entry.day) > getToday();
  }

  calendarUrl(): string {
    if (!this.calendar) return "";
    return `https://adventar.org/calendars/${this.calendar.id}`;
  }
}
</script>

<style lang="scss" scoped>
.mainInner {
  padding: 0;
}

.header {
  width: 100%;
}

.header > .inner {
  max-width: $content-max-width;
  margin: 0 auto;
}

.header .title {
  margin: 0;
}

.header .title a {
  font-size: 14px;
  display: block;
  padding: 0 10px;
  height: 38px;
  line-height: 38px;
  white-space: nowrap;
  color: #fff;
  text-decoration: none;
}

.header .title svg {
  margin-left: 5px;
}

.CalendarTable {
  position: relative;
}

.table {
  width: 100%;
  table-layout: fixed;
  border-collapse: collapse;
}

.table thead th {
  border: 1px solid #f1f1f1;
  border-bottom: 3px solid #f1f1f1;
  background: #aaa;
  color: #fff;
  text-align: center;
  font-size: 12px;
  text-transform: uppercase;

  .inner {
    height: 18px;
    line-height: 18px;
  }
}

.table thead th:first-child {
  background: #e7998e;
}

.table thead th:last-child {
  background: #87a3d0;
}

.table .cell {
  border: 1px solid #e5e5e5;
  background-color: #fff;
  vertical-align: top;

  .inner {
    height: 72px;
    overflow: hidden;
  }

  &.is-editing {
    background-color: #ffe2e7;
  }

  &.is-posted {
    background-image: url("~assets/bg-posted.png");
    background-repeat: no-repeat;
    background-position: center top;
  }
}

.table .day {
  display: inline-block;
  font-size: 14px;
  padding: 0 3px;
  font-weight: bold;
  color: #aaa;

  td:first-child & {
    color: #e7998e;
  }

  td:last-child & {
    color: #87a3d0;
  }
}

.table .entryUser {
  text-align: center;
  color: #666;
  display: block;
  text-decoration: none;
}

.table .userName {
  font-size: 10px;
  line-height: 1.2;
  margin: 5px 0;
  overflow: hidden;
  white-space: nowrap;
}

.table .forceCancelBtn {
  position: absolute;
  top: 0px;
  right: 2px;
  color: #e5004f;
  cursor: pointer;

  &:hover {
    opacity: 0.8;
  }
}

.table .editBtn {
  position: absolute;
  top: 0px;
  right: 2px;
  color: #13b5b1;
  cursor: pointer;

  &:hover {
    opacity: 0.8;
  }
}

.table .entryAction {
  text-align: center;
  margin: 5px 0 10px;
}

.table .entryAction a {
  padding: 5px 10px;
  font-size: 11px;
  display: inline-block;
  border-radius: 5px;
  margin: 0;
  font-family: inherit;
  cursor: pointer;
  text-decoration: none;
  border: none;
  outline: none;
  font-weight: bold;
  background-color: #e0e0e0;
  background-image: linear-gradient(to bottom, #e0e0e0, #e0e0e0 50%, #d3d3d3 50%, #d3d3d3);
  color: #999;
}

.table .entryAction a:hover {
  color: #fff;
  background-color: #ef7266;
  background-image: linear-gradient(to bottom, #ef7266, #ef7266 50%, #e45541 50%, #e45541);
}
</style>
