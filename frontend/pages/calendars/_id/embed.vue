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
                  <th>SUN</th>
                  <th>MON</th>
                  <th>TUE</th>
                  <th>WED</th>
                  <th>THU</th>
                  <th>FRI</th>
                  <th>SAT</th>
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

.header h2 {
  margin: 0;
}

.header h2 a {
  font-size: 14px;
  white-space: nowrap;
  color: #fff;
  text-decoration: none;
}

.header h2 svg {
  margin-left: 5px;
}

.header > .inner {
  max-width: $content-max-width;
  padding: 3px 10px 5px 10px;
  margin: 0 auto;
  position: relative;
}

.title {
  font-size: 20px;
  font-weight: bold;
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
  position: relative;

  &.is-editing {
    background-color: #ffe2e7;
  }

  &.is-posted {
    background-image: url("~assets/bg-posted.png");
    background-repeat: no-repeat;
    background-position: center top;
  }
}

.table .inner {
  position: relative;
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

.dialog {
  width: 310px;
  border: 1px solid #ccc;
  padding: 35px 10px 15px 15px;
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  color: #666;
  position: absolute;
  background-color: #fff;
  z-index: 1;
}

.dialog .arrow {
  display: inline-block;
  position: absolute;
  bottom: -15px;
  width: 0;
  height: 0;
  border: 8px solid black;
  border-color: transparent transparent #fff #fff;
  transform-origin: 0 0;
  transform: rotate(-45deg);
  box-shadow: -2px 2px 2px 0 rgba(0, 0, 0, 0.2);
}

.dialog .day {
  position: absolute;
  top: 12px;
  left: 12px;
  font-weight: bold;
  font-size: 12px;
}
.dialog .closeBtn {
  position: absolute;
  top: 10px;
  right: 20px;
  cursor: pointer;

  &:hover {
    opacity: 0.8;
  }
}

.dialog .formRow {
  padding: 5px 0;
}

.dialog input[type="text"] {
  width: 270px;
  padding: 5px;
  border: 1px solid #e5e5e5;
  outline: none;
  font-size: 12px;
  border-radius: 3px;
  font-family: inherit;
  -webkit-appearance: none;
}

.dialog .buttons {
  position: relative;
  margin: 10px 0 0 20px;
}

.dialog .submit {
  font-size: 14px;
  color: #fff;
  border: 2px solid #28a745;
  background-color: #28a745;
  padding: 6px 20px;
  border-radius: 10px;
  cursor: pointer;

  &:hover {
    opacity: 0.8;
  }
}

.dialog .cancel {
  font-size: 12px;
  color: #dc3545;
  cursor: pointer;
  position: absolute;
  top: 8px;
  right: 10px;

  &:hover {
    opacity: 0.8;
  }
}
</style>
