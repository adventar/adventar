<template>
  <div>
    <GlobalHeader />

    <div v-if="calendar">
      <header :style="{ backgroundColor: calendarColor() }">
        <div>
          <h2>{{ calendar.title }} Advent Calendar {{ calendar.year }}</h2>
          <div>{{ calendar.entries.length }}/25人</div>
          <div>
            作成者: <UserIcon :user="calendar.owner" size="22" />
            {{ calendar.owner.name }}
          </div>
          <nuxt-link :to="`/calendars/${calendar.id}/edit`" v-if="isOwnCalendar(calendar)">編集</nuxt-link>
        </div>
      </header>
      <main>
        <div>
          <VueMarkdown>{{ calendar.description }}</VueMarkdown>
          <table class="calendarTable">
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
                <td v-for="cell in cells" :key="cell.day">
                  <div class="popup" v-if="cell.entryable && displayedPopupCellDay === cell.day" @click.stop>
                    <span role="button" class="popupCloseBtn" @click="displayedPopupCellDay = null">
                      <font-awesome-icon icon="times" />
                    </span>
                    <form @submit.prevent="handleSubmitPopupForm(cell.entry)">
                      <div class="popupRow">
                        <font-awesome-icon icon="comment" />
                        <input
                          type="text"
                          placeholder="記事の内容の予定などを入力してください"
                          :value="cell.entry.comment"
                          :ref="`inputComment${cell.entry.day}`"
                        />
                      </div>
                      <div class="popupRow">
                        <font-awesome-icon icon="link" />
                        <input
                          type="text"
                          placeholder="登録した日になったらURLを入力してください"
                          :value="cell.entry.url"
                          :ref="`inputUrl${cell.entry.day}`"
                        />
                      </div>
                      <div class="popupAction">
                        <button class="submit" type="submit">保存</button>
                        <span role="button" class="cancel" @click="handleClickDeleteEntry(cell.entry)">
                          登録をキャンセル
                        </span>
                      </div>
                    </form>
                  </div>
                  <div v-if="cell.entryable" class="inner">
                    <span class="day">{{ cell.day }}</span>
                    <div v-if="cell.entry">
                      <div class="calendarTable-entryUser">
                        <UserIcon :user="cell.entry.owner" size="50" />
                        <div>{{ cell.entry.owner.name }}</div>
                      </div>
                      <span
                        class="editBtn"
                        role="button"
                        v-if="isOwnEntry(cell.entry)"
                        @click.stop="handleClickEditEntry(cell.entry)"
                      >
                        <font-awesome-icon icon="edit" />
                      </span>
                      <span
                        class="cancelBtn"
                        role="button"
                        v-if="forceCancelable(cell.entry)"
                        @click="handleClickDeleteEntry(cell.entry)"
                      >
                        <font-awesome-icon icon="times" />
                      </span>
                    </div>
                    <div v-else class="calendarTable-entryForm">
                      <button @click="handleClickCreateEntry(cell.day)">登録</button>
                    </div>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
          <table class="entryList">
            <tr v-for="entry in calendar.entries" :key="entry.day">
              <th class="date">12/{{ entry.day }}</th>
              <td class="user">
                <UserIcon :user="entry.owner" size="24" />
                {{ entry.owner.name }}
              </td>
              <td class="body">
                <div v-if="entry.comment"><font-awesome-icon icon="comment" /> {{ entry.comment }}</div>
                <div v-if="entry.title && !isFutureEntry(entry)">
                  <font-awesome-icon icon="file" /> {{ entry.title }}
                </div>
                <div v-if="entry.url && !isFutureEntry(entry)">
                  <font-awesome-icon icon="link" />
                  <a :href="entry.url">{{ entry.url }}</a>
                </div>
              </td>
              <td class="image">
                <img :src="entry.imageUrl" v-if="entry.imageUrl && !isFutureEntry(entry)" width="100" height="100" />
              </td>
            </tr>
          </table>
        </div>
      </main>
    </div>
  </div>
</template>

<script lang="ts">
import dayjs from "dayjs";
import { Component, Vue } from "nuxt-property-decorator";
import VueMarkdown from "vue-markdown";
import { getCalendar, createEntry, updateEntry, deleteEntry } from "~/lib/GrpcClient";
import * as RestClient from "~/lib/RestClient";
import { calendarColor } from "~/lib/utils/Colors";
import { Calendar, Entry } from "~/types/adventar";
import { getToken } from "~/lib/Auth";
import GlobalHeader from "~/components/GlobalHeader.vue";
import UserIcon from "~/components/UserIcon.vue";

function getRows(calendar: Calendar): any[] {
  const year = calendar.year;
  const endDay = dayjs(new Date(year, 11, 25));
  let currentDay = dayjs(new Date(year, 11, 1)).startOf("week");

  const entryMapByDay: { [key: number]: Entry } = {};
  if (calendar !== null && calendar.entries !== undefined) {
    calendar.entries.forEach(entry => {
      if (entry.day) {
        entryMapByDay[entry.day] = entry;
      }
    });
  }

  const rows: any[] = [];
  while (currentDay <= endDay) {
    const cells: any[] = [];
    for (let i = 0; i < 7; i++) {
      const day = currentDay.date();
      const entry = entryMapByDay[day] || null;
      const entryable = currentDay.month() === 11 && day <= 25;
      const cell = { day, entry, entryable };
      cells.push(cell);
      currentDay = currentDay.add(1, "day");
    }
    rows.push(cells);
  }

  return rows;
}

@Component({
  components: { GlobalHeader, UserIcon, VueMarkdown }
})
export default class extends Vue {
  calendar: Calendar | null = null;
  displayedPopupCellDay: number | null;
  rows: any[];

  async asyncData({ params }) {
    if (process.server) {
      const calendar = await RestClient.getCalendar(params.id);
      const rows = getRows(calendar);
      return { calendar, rows };
    }
  }

  async mounted() {
    this.calendar = await getCalendar(Number(this.$route.params.id));
    this.rows = getRows(this.calendar);
    document.addEventListener("click", this.handleClickDocument);
  }

  destroyed() {
    document.removeEventListener("click", this.handleClickDocument);
  }

  handleClickDocument() {
    this.displayedPopupCellDay = null;
  }

  data() {
    return {
      displayedPopupCellDay: null
    };
  }

  async refetchCalendar() {
    this.calendar = await getCalendar(this.calendar!.id);
    this.rows = getRows(this.calendar);
  }

  async handleClickCreateEntry(day): Promise<void> {
    const calendarId = this.calendar!.id;
    const token = await getToken();
    await createEntry({ calendarId, day, token });
    await this.refetchCalendar();
    this.displayedPopupCellDay = day;
  }

  async handleClickEditEntry(entry: Entry): Promise<void> {
    this.displayedPopupCellDay = entry.day || null;
  }

  async handleClickDeleteEntry(entry: Entry): Promise<void> {
    if (!window.confirm("登録をキャンセルします")) return;
    const token = await getToken();
    await deleteEntry({ entryId: entry.id, token });
    this.calendar = await getCalendar(this.calendar!.id);
    this.displayedPopupCellDay = null;
    await this.refetchCalendar();
  }

  async handleSubmitPopupForm(entry: Entry): Promise<void> {
    const entryId = entry.id;
    const comment = this.$refs[`inputComment${entry.day}`][0].value;
    const url = this.$refs[`inputUrl${entry.day}`][0].value;
    const token = await getToken();
    await updateEntry({ entryId, comment, url, token });
    await this.refetchCalendar();
    this.displayedPopupCellDay = null;
  }

  forceCancelable(entry: Entry): boolean {
    return (
      this.calendar !== null &&
      !this.isOwnEntry(entry) &&
      this.isOwnCalendar(this.calendar) &&
      !entry.url &&
      // TODO: Fix to JST
      dayjs(new Date(this.calendar.year, 12, entry.day)).add(1, "day") < dayjs(new Date())
    );
  }

  calendarColor(): string {
    return calendarColor(this.calendar!.id);
  }

  isOwnCalendar(calendar: Calendar): boolean {
    if (!this.$store.state.user) return false;
    if (!calendar.owner) return false;
    return calendar.owner.id === this.$store.state.user.id;
  }

  isOwnEntry(entry: Entry): boolean {
    if (!this.$store.state.user) return false;
    if (!entry.owner) return false;
    return entry.owner.id === this.$store.state.user.id;
  }

  isFutureEntry(_): boolean {
    // TODO
    return false;
  }
}
</script>

<style scoped>
header {
  color: #fff;
}

header > div {
  width: 1000px;
  padding: 30px 12px;
  margin: 0 auto;
}

header .userIcon {
  width: 22px;
  height: 22px;
  border-radius: 50%;
  vertical-align: middle;
}

.calendarTable {
  width: 100%;
  table-layout: fixed;
  border-collapse: collapse;
}
.calendarTable thead th {
  border: 1px solid #f1f1f1;
  border-bottom: 3px solid #f1f1f1;
  background: #aaa;
  height: 50px;
  color: #fff;
  text-align: left;
  padding-left: 15px;
  font-size: 15px;
  text-transform: uppercase;
}

.calendarTable thead th:first-child {
  background: #e7998e;
}

.calendarTable thead th:last-child {
  background: #87a3d0;
}

.calendarTable td {
  border: 1px solid #e5e5e5;
  background-color: #fff;
  vertical-align: top;
  position: relative;
}

.calendarTable .inner {
  position: relative;
  min-height: 140px;
}

.calendarTable .day {
  position: absolute;
  top: 12px;
  left: 15px;
  font-size: 24px;
  font-weight: bold;
  color: #aaa;
}

.calendarTable td:first-child .day {
  color: #e7998e;
}

.calendarTable td:last-child .day {
  color: #87a3d0;
}

.calendarTable-entryForm {
  padding-top: 60px;
  text-align: center;
}

.calendarTable-entryForm button {
  padding: 0 15px;
  height: 32px;
  line-height: 32px;
  font-size: 15px;
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

.calendarTable-entryForm button:hover {
  color: #fff;
  background-color: #ef7266;
  background-image: linear-gradient(to bottom, #ef7266, #ef7266 50%, #e45541 50%, #e45541);
}

.calendarTable-entryUser {
  text-align: center;
  position: relative;
  top: 50px;
  color: #666;
}

.calendarTable-entryUser .userIcon {
  width: 50px;
  height: 50px;
  border-radius: 50px;
}

.calendarTable .cancelBtn {
  position: absolute;
  top: 15px;
  right: 10px;
  color: #e5004f;
  cursor: pointer;
}

.calendarTable .editBtn {
  position: absolute;
  top: 15px;
  right: 10px;
  color: #13b5b1;
  cursor: pointer;
}

.entryList {
  border-collapse: collapse;
  border-spacing: 0;
  font-size: 15px;
  color: #666;
  width: 100%;
  margin-top: 50px;
}

.entryList td,
.entryList th {
  vertical-align: top;
  padding: 20px 10px;
  border-top: 1px solid #e3e4e3;
  background-color: #fcfcfc;
}

.entryList .date {
  font-size: 16px;
  font-weight: bold;
  margin: 0;
  width: 70px;
}

.entryList .userIcon {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  vertical-align: middle;
  margin-right: 5px;
}

.entryList .body {
  width: 570px;
}

.entryList .body > div {
  padding-bottom: 5px;
}

.entryList .body svg {
  margin-right: 5px;
  width: 15px;
}

.entryList .image {
  width: 100px;
  text-align: center;
  vertical-align: top;
}

.popup {
  width: 350px;
  height: 130px;
  border: 1px solid #ccc;
  padding: 35px 20px 15px 20px;
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  color: #666;
  top: -152px;
  left: -126px;
  position: absolute;
  background-color: #fff;
  z-index: 1;
}

.popup:before,
.popup:after {
  top: 100%;
  left: 50%;
  border: solid transparent;
  content: " ";
  height: 0;
  width: 0;
  position: absolute;
  pointer-events: none;
}
.popup:before {
  border-color: rgba(204, 204, 204, 0);
  border-top-color: #cccccc;
  border-width: 12px;
  margin-left: -12px;
}
.popup:after {
  border-color: rgba(255, 255, 255, 0);
  border-top-color: #ffffff;
  border-width: 10px;
  margin-left: -10px;
}

.popup .popupCloseBtn {
  position: absolute;
  top: 10px;
  right: 30px;
  cursor: pointer;
}

.popupRow {
  padding: 5px 0;
}

.popup input[type="text"] {
  width: 310px;
  padding: 5px;
  border: 1px solid #e5e5e5;
  outline: none;
  font-size: 13px;
  border-radius: 3px;
  font-family: inherit;
}

.popup .popupAction {
  margin: 10px 0 0 20px;
}

.popup .submit {
  font-size: 14px;
  color: #fff;
  border: 2px solid #28a745;
  background-color: #28a745;
  padding: 6px 20px;
  border-radius: 10px;
  cursor: pointer;
  margin-right: 100px;
}

.popup .cancel {
  font-size: 14px;
  color: #dc3545;
  padding: 6px 20px;
  cursor: pointer;
}
</style>
