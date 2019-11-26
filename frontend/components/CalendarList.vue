<template>
  <div class="CalendarList">
    <div v-if="calendars === null" class="loading">
      <font-awesome-icon icon="circle-notch" spin />
    </div>
    <ul v-else class="list">
      <li v-for="calendar in calendars" :key="calendar.id" class="item">
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
  @Prop() readonly calendars: Calendar[] | null;

  getCalendarColor(id) {
    return calendarColor(id);
  }
}
</script>

<style lang="scss" scoped>
.loading {
  text-align: center;
  font-size: 42px;
  padding: 30px 0;
  color: #e35541;
}

.list {
  margin: 0;
  padding: 0;
}

.item {
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

.title {
  border-radius: 3px 0 0 3px;
  font-size: 14px;
  padding: 15px;
  color: #fff;
  width: 75%;
  box-sizing: border-box;
  font-weight: bold;
  text-decoration: none;
}

.info {
  font-size: 12px;
  width: 25%;
  padding: 0px 10px;
  background-color: #fff;
  text-align: right;
  font-weight: bold;
  box-sizing: border-box;
}

.counter {
  color: #ef7266;
}

.indicator {
  display: none;
}

@media (min-width: $mq-break-small) {
  .CalendarList::after {
    content: "";
    display: block;
    clear: both;
  }

  .list {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
  }

  .item {
    width: 48%;
    display: block;
    margin: 0 0 40px 0;
  }

  .title {
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

  .info {
    width: auto;
    padding: 10px;
    border-radius: 0 0 3px 3px;
    position: relative;
    height: 40px;
  }

  .counter {
    float: right;
  }

  .indicator {
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

  .item [data-value] {
    position: absolute;
    left: 0;
    top: 0;
    height: 21px;
    background-color: #d4d4d4;
    border-radius: 3px 0 0 3px;
  }

  @for $i from 1 through 25 {
    .item [data-value="#{$i}"] {
      width: $i * 7px;
    }
  }

  .item [data-value="25"] {
    border-radius: 3px;
    background-color: #5fca6b;
  }
}

@media (min-width: $mq-break-medium) {
  .item {
    width: 31%;
  }

  .list::after {
    content: "";
    width: 31%;
  }
}
</style>
