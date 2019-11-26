<template>
  <div>
    <div class="hero">
      <div class="inner">
        <button v-if="calendarCreatable" @click="handleClickCreateBtn">
          {{ currentYear }}年のAdvent Calendarを作る
        </button>
      </div>
    </div>

    <main>
      <div class="mainInner">
        <CalendarSearchForm @submit="handleSubmit" />
        <CalendarList :calendars="calendars" />

        <nuxt-link v-if="hasMore" :to="`/calendars?year=${currentYear}`">もっと見る</nuxt-link>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { listCalendars } from "~/lib/GrpcClient";
import { getCurrentYear, getCalendarCreatable } from "~/lib/Configuration";
import { Calendar } from "~/types/adventar";
import PageHeader from "~/components/PageHeader.vue";
import CalendarSearchForm from "~/components/CalendarSearchForm.vue";
import CalendarList from "~/components/CalendarList.vue";

@Component({
  components: { PageHeader, CalendarSearchForm, CalendarList }
})
export default class extends Vue {
  currentYear = getCurrentYear();
  calendarCreatable = getCalendarCreatable();
  calendars: Calendar[] | null = null;

  pageSize = 30;

  async mounted() {
    const calendars = await listCalendars({ year: this.currentYear, pageSize: this.pageSize });
    this.calendars = calendars;
  }

  handleSubmit(query) {
    this.$router.push(`/calendars?year=${this.currentYear}&query=${query}`);
  }

  get hasMore() {
    return this.calendars && this.calendars.length === this.pageSize;
  }

  handleClickCreateBtn() {
    if (this.$store.state.user) {
      this.$router.push("/new");
    } else {
      alert("ログインしてください。");
    }
  }
}
</script>

<style lang="scss" scoped>
.hero {
  background-image: url("~assets/hero.png");
  background-position: right -120px top;
  background-size: auto 160px;
  height: 160px;
  position: relative;
  padding: 0 12px;
}

.hero .inner {
  max-width: $content-max-width;
  margin: 0 auto;
  position: relative;
}

.hero button {
  padding: 10px;
  font-size: 12px;
  background: rgba(255, 255, 255, 0.85);
  color: #333;
  border-radius: 5px;
  position: absolute;
  top: 60px;
  left: 0;
  border: none;
  outline: none;
  cursor: pointer;

  &:hover {
    background: rgba(255, 255, 255, 0.8);
  }
}

.mainInner {
  padding-top: 15px;
}

.CalendarSearchForm {
  margin-bottom: 15px;
}

@media (min-width: $mq-break-small) {
  .hero {
    background-size: auto 280px;
    background-position: right -210px top;
    height: 280px;
  }

  .hero button {
    top: 120px;
    font-size: 18px;
    padding: 12px 20px;
  }

  .mainInner {
    padding-top: 30px;
  }

  .CalendarSearchForm {
    margin-bottom: 30px;
  }
}

@media (min-width: $mq-break-medium) {
  .hero {
    background-position: center top;
  }

  .hero button {
    left: 0;
  }
}
</style>
