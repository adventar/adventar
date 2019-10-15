<template>
  <div>
    <PageHeader>カレンダーを作る</PageHeader>

    <main>
      <div class="mainInner">
        <div v-if="calendarCreatable">
          <CalendarForm />
        </div>
        <div v-else>
          現在カレンダーの新規作成はできません。
        </div>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import PageHeader from "~/components/PageHeader.vue";
import CalendarForm from "~/components/CalendarForm.vue";
import { getCalendarCreatable } from "~/lib/Configuration";

@Component({
  components: { PageHeader, CalendarForm }
})
export default class extends Vue {
  calendarCreatable = getCalendarCreatable();
}
</script>
