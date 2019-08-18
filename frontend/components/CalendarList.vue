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
        <div class="info">
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

<style lang="scss" scoped>
ul {
  margin: 0;
  padding: 0;
}

li {
  font-size: 14px;
  margin: 0 0 10px 0;
  padding: 0;
  list-style: none;
  display: flex;
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.07);
  justify-content: center;
  align-items: center;
  background-color: #fff;
  border-radius: 3px;
}

li .title {
  border-radius: 3px 0 0 3px;
  font-size: 14px;
  padding: 15px;
  color: #fff;
  width: 75%;
  box-sizing: border-box;
  font-weight: bold;
  text-decoration: none;
}

li .info {
  font-size: 12px;
  width: 25%;
  padding: 0px 10px;
  background-color: #fff;
  text-align: right;
  font-weight: bold;
  box-sizing: border-box;
}

li .counter {
  color: #ef7266;
}

li .indicator {
  display: none;
}

@media (min-width: $mq-break-small) {
  .CalendarList::after {
    content: "";
    display: block;
    clear: both;
  }

  ul {
    display: flex;
    flex-wrap: wrap;
    justify-content: flex-start;
  }

  li {
    width: 47%;
    display: block;
    margin: 0 3% 40px 0;
  }

  li .title {
    display: block;
    display: flex;
    align-items: center;
    font-size: 20px;
    position: relative;
    height: 150px;
    border-radius: 3px 3px 0 0;
    width: 100%;
    padding: 0 15px;
    color: #fff;
    font-weight: bold;
    text-decoration: none;
  }

  li .info {
    width: auto;
    padding: 10px;
    border-radius: 0 0 3px 3px;
    position: relative;
    height: 40px;
  }

  li .counter {
    float: right;
  }

  li .indicator {
    display: inline-block;
    vertical-align: middle;
    background-color: #f1f1f1;
    height: 21px;
    width: 175px;
    border-radius: 3px;
    box-shadow: 0px 1px 1px rgba(0, 0, 0, 0.2) inset;
    position: absolute;
    left: 10px;
  }

  li [data-value] {
    position: absolute;
    left: 0;
    top: 0;
    height: 21px;
    background-color: #d4d4d4;
    border-radius: 3px 0 0 3px;
  }

  @for $i from 1 through 25 {
    li [data-value="#{$i}"] {
      width: $i * 7px;
    }
  }

  li [data-value="25"] {
    border-radius: 3px;
    background-color: #5fca6b;
  }
}

@media (min-width: $mq-break-middle) {
  li {
    width: 30%;
  }
}
</style>
