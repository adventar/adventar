<template>
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
            class="cell"
            :class="{
              'is-editing': displayedDialogEntry && cell.day === displayedDialogEntry.day,
              'is-posted': cell.entry && cell.entry.url && !isFutureEntry(cell.entry)
            }"
            v-for="(cell, idx) in cells"
            :data-idx="idx"
            :key="cell.day"
            :ref="`cell-${cell.day}`"
          >
            <div v-if="cell.entryable" class="inner">
              <span class="day">{{ cell.day }}</span>
              <div v-if="cell.entry">
                <span class="entryUser">
                  <UserIcon :user="cell.entry.owner" size="24" />
                  <div class="userName">{{ cell.entry.owner.name }}</div>
                </span>
                <span
                  class="editBtn"
                  role="button"
                  v-if="isOwnEntry(cell.entry)"
                  @click.stop="handleClickEditEntry(cell.entry)"
                >
                  <font-awesome-icon icon="edit" />
                </span>
                <span
                  class="forceCancelBtn"
                  role="button"
                  v-if="forceCancelable(cell.entry)"
                  @click="handleClickForceCancel(cell.entry)"
                >
                  <font-awesome-icon icon="times" />
                </span>
              </div>
              <div v-else class="entryAction">
                <button @click="handleClickCreateEntry(cell.day)">登録</button>
              </div>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="dialog" :style="dialogStyle" v-if="displayedDialogEntry" @click.stop>
      <div class="day">12/{{ displayedDialogEntry.day }}</div>
      <span role="button" class="closeBtn" @click="hideDialog()">
        <font-awesome-icon icon="times" />
      </span>
      <form @submit.prevent="handleSubmitEntryForm()">
        <div class="formRow">
          <font-awesome-icon icon="comment" />
          <input type="text" placeholder="記事の内容の予定などを入力してください" v-model="inputComment" />
        </div>
        <div class="formRow">
          <font-awesome-icon icon="link" />
          <input type="text" placeholder="登録した日になったらURLを入力してください" v-model="inputUrl" />
        </div>
        <div class="buttons">
          <button class="submit" type="submit">保存</button>
          <span role="button" class="cancel" @click="handleClickDeleteEntry()">
            登録をキャンセル
          </span>
        </div>
      </form>
      <span class="arrow" :style="dialogArrowStyle"></span>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "nuxt-property-decorator";
import UserIcon from "~/components/UserIcon.vue";
import { Calendar, Entry, User } from "~/types/adventar";
import { getToday } from "~/lib/Configuration";
import dayjs from "dayjs";

@Component({
  components: { UserIcon }
})
export default class extends Vue {
  @Prop() readonly calendar: Calendar;
  @Prop() readonly currentUser: User;
  @Prop() readonly onCreateEntry: (day: number) => Promise<Entry>;
  @Prop() readonly onUpdateEntry: (
    entryId: number,
    { comment, url }: { comment: string; url: string }
  ) => Promise<void>;
  @Prop() readonly onDeleteEntry: (entryId: number) => Promise<void>;

  displayedDialogEntry: Entry | null;
  inputComment: string;
  inputUrl: string;
  dialogStyle: object;
  dialogArrowStyle: object;

  get rows() {
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

  mounted() {
    document.addEventListener("click", this.handleClickDocument);
  }

  destroyed() {
    document.removeEventListener("click", this.handleClickDocument);
  }

  handleClickDocument() {
    this.hideDialog();
  }

  data() {
    return {
      displayedDialogEntry: null,
      inputComment: "",
      inputUrl: ""
    };
  }

  forceCancelable(entry: Entry): boolean {
    return (
      this.calendar !== null &&
      !this.isOwnEntry(entry) &&
      this.isOwnCalendar(this.calendar) &&
      !entry.url &&
      dayjs(new Date(this.calendar.year, 11, entry.day))
        .add(1, "day")
        .endOf("day") < dayjs(getToday())
    );
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

  showDialog(entry): void {
    const boxWidth = 338;
    const boxHeight = 180;
    const cell = this.$refs[`cell-${entry.day}`][0];
    const rect = cell.getBoundingClientRect();
    const cellCenterLeft = cell.offsetLeft + cell.clientWidth / 2;
    const windowLeftPadding = rect.left - cell.offsetLeft;
    let dialogLeft = cellCenterLeft - boxWidth / 2;
    if (dialogLeft < -windowLeftPadding) {
      dialogLeft = -windowLeftPadding;
    }
    if (dialogLeft + boxWidth > window.innerWidth - windowLeftPadding) {
      dialogLeft = window.innerWidth - windowLeftPadding - boxWidth;
    }
    this.dialogStyle = {
      top: `${cell.offsetTop - boxHeight + 5}px`,
      left: `${dialogLeft}px`
    };
    this.dialogArrowStyle = {
      left: `${cellCenterLeft - dialogLeft - 12}px`
    };
    this.inputComment = entry.comment;
    this.inputUrl = entry.url;
    this.displayedDialogEntry = entry;
  }

  hideDialog(): void {
    this.inputComment = "";
    this.inputUrl = "";
    this.dialogStyle = {};
    this.dialogArrowStyle = {};
    this.displayedDialogEntry = null;
  }

  isFutureEntry(entry): boolean {
    // TODO: Fix to JST
    return new Date(this.calendar.year, 11, entry.day) > getToday();
  }

  async handleClickCreateEntry(day: number): Promise<void> {
    if (!this.currentUser) {
      window.alert("登録にはログインが必要です。");
      return;
    }
    const entry = await this.onCreateEntry(day);
    this.showDialog(entry);
  }

  async handleClickEditEntry(entry: Entry): Promise<void> {
    this.showDialog(entry);
  }

  async handleClickDeleteEntry(): Promise<void> {
    if (this.displayedDialogEntry === null) {
      throw new Error("Entry is not selected");
    }
    if (!window.confirm("登録をキャンセルします")) return;
    this.onDeleteEntry(this.displayedDialogEntry.id);
    this.hideDialog();
  }

  async handleSubmitEntryForm(): Promise<void> {
    if (this.displayedDialogEntry === null) {
      throw new Error("Entry is not selected");
    }
    const comment = this.inputComment;
    const url = this.inputUrl;
    await this.onUpdateEntry(this.displayedDialogEntry.id, { comment, url });
    this.hideDialog();
  }

  handleClickForceCancel(entry: Entry) {
    const message =
      "カレンダーのオーナーに限り、登録日を1日以上過ぎても投稿のないエントリーを強制的にキャンセルできます。キャンセルしますか？";
    if (window.confirm(message)) {
      this.onDeleteEntry(entry.id);
    }
  }
}
</script>

<style lang="scss" scoped>
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
  height: 50px;
  color: #fff;
  text-align: center;
  font-size: 11px;
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
  font-size: 16px;
  padding: 3px;
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
  padding-top: 5px;
  display: block;
  text-decoration: none;
}

.table .userName {
  font-size: 10px;
  line-height: 1.2;
  margin: 5px 0;
  overflow: hidden;
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

.table .entryAction button {
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

.table .entryAction button:hover {
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
  font-size: 14px;
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

@media (min-width: $mq-break-small) {
  .table thead th {
    font-size: 15px;
  }

  .table .day {
    font-size: 24px;
    padding: 5px 10px;
  }

  .table .UserIcon {
    width: 50px;
    height: 50px;
  }

  .table .userName {
    font-size: 14px;
    line-height: 1.4;
    margin: 10px 0;
  }

  .table .forceCancelBtn {
    top: 10px;
    right: 10px;
  }

  .table .editBtn {
    top: 10px;
    right: 10px;
  }

  .table .entryAction {
    padding: 15px 0 35px;
  }

  .table .entryAction button {
    font-size: 15px;
    padding: 5px 15px;
  }
}
</style>
