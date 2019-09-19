<template>
  <div>
    <div class="hero"></div>

    <main>
      <div>
        <CalendarSearchForm @submit="handleSubmit" style="margin-bottom: 30px" />
        <CalendarList :calendars="calendars" />

        <nuxt-link :to="`/calendars?year=${currentYear}`">もっと見る</nuxt-link>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { listCalendars } from "~/lib/GrpcClient";
import { getCurrentYear } from "~/lib/Configuration";
import { Calendar } from "~/types/adventar";
import PageHeader from "~/components/PageHeader.vue";
import CalendarSearchForm from "~/components/CalendarSearchForm.vue";
import CalendarList from "~/components/CalendarList.vue";

@Component({
  components: { PageHeader, CalendarSearchForm, CalendarList }
})
export default class extends Vue {
  currentYear = getCurrentYear();
  calendars: Calendar[] = [];

  async mounted() {
    const pageSize = 30;
    const calendars = await listCalendars({ year: this.currentYear, pageSize });
    this.calendars = calendars;
  }

  async handleSubmit(query) {
    this.$router.push(`/calendars?year=${this.currentYear}&query=${query}`);
  }
}
</script>

<style lang="scss" scoped>
.hero {
  background-image: url("~assets/hero.png");
  background-position: center top;
  background-size: auto 160px;
  height: 160px;
  position: relative;
}

.hero h1 {
  font-size: 20px;
  font-weight: bold;
  margin: 0;
  padding: 20px 0;
}

.hero a {
  display: inline-block;
  padding: 10px;
  background-color: #efefef;
  color: #000;
  border-radius: 5px;
  text-decoration: none;
  position: absolute;
  left: 10px;
}

.hero a:first-child {
  bottom: 80px;
}

.hero a:last-child {
  bottom: 20px;
}

@media (min-width: $mq-break-small) {
  .hero {
    background-size: auto 280px;
    height: 280px;
  }
  .hero h1 {
    margin: 0;
    padding: 40px 0 15px 0;
  }
}
</style>
