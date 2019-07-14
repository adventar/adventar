<template>
  <div class="CalendarList">
    <ul>
      <li v-for="calendar in calendars" :key="calendar.id">
        <nuxt-link
          :to="`/calendars/${calendar.id}`"
          class="title"
          :style="`background-color: ${getCalendarColor(calendar.id)}`"
          >{{ calendar.title }}</nuxt-link
        >
        <div class="bottom">
          <span class="indicator"><span :data-value="calendar.entryCount"></span></span>
          <span class="counter">{{ calendar.entryCount }}/25人</span>
        </div>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "nuxt-property-decorator";
import { Calendar } from "~/types/adventar";
import { calendarColor } from "~/lib/utils/Colors";

@Component
export default class extends Vue {
  @Prop() readonly calendars!: Calendar[];

  getCalendarColor(id) {
    return calendarColor(id);
  }
}
</script>

<style scoped>
.CalendarList::after {
  content: "";
  display: block;
  clear: both;
}

ul {
  margin: 0;
  padding: 0;
}

li {
  margin: 0 20px 0 0;
  padding: 0;
  list-style: none;
  float: left;
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.07);
}

li .bottom {
  padding: 10px;
}

li .title {
  font-size: 20px;
  display: table-cell;
  position: relative;
  vertical-align: middle;
  height: 150px;
  width: 280px;
  padding: 0 15px;
  color: #fff;
  font-weight: bold;
  text-decoration: none;
  -webkit-font-smoothing: antialiased;
}

li .counter {
  color: #ef7266;
  float: right;
}

li .indicator {
  display: inline-block;
  vertical-align: middle;
  background-color: #f1f1f1;
  height: 21px;
  width: 225px;
  border-radius: 3px;
  box-shadow: 0px 1px 1px rgba(0, 0, 0, 0.2) inset;
  position: relative;
}

li [data-value] {
  position: absolute;
  left: 0;
  top: 0;
  height: 21px;
  background-color: #d4d4d4;
  border-radius: 3px 0 0 3px;
}

li [data-value="1"] {
  width: 9px;
}
li [data-value="2"] {
  width: 18px;
}
li [data-value="3"] {
  width: 27px;
}
li [data-value="4"] {
  width: 36px;
}
li [data-value="5"] {
  width: 45px;
}
li [data-value="6"] {
  width: 54px;
}
li [data-value="7"] {
  width: 63px;
}
li [data-value="8"] {
  width: 72px;
}
li [data-value="9"] {
  width: 81px;
}
li [data-value="10"] {
  width: 90px;
}
li [data-value="11"] {
  width: 99px;
}
li [data-value="12"] {
  width: 108px;
}
li [data-value="13"] {
  width: 117px;
}
li [data-value="14"] {
  width: 126px;
}
li [data-value="15"] {
  width: 135px;
}
li [data-value="16"] {
  width: 144px;
}
li [data-value="17"] {
  width: 153px;
}
li [data-value="18"] {
  width: 162px;
}
li [data-value="19"] {
  width: 171px;
}
li [data-value="20"] {
  width: 180px;
}
li [data-value="21"] {
  width: 189px;
}
li [data-value="22"] {
  width: 198px;
}
li [data-value="23"] {
  width: 207px;
}
li [data-value="24"] {
  width: 216px;
}
li [data-value="25"] {
  width: 225px;
  border-radius: 3px;
  background-color: #5fca6b;
}
</style>
