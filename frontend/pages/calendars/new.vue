<template>
  <div>
    <GlobalHeader />

    <PageHeader>カレンダーを作る</PageHeader>

    <main>
      <div>
        <CalendarForm />
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import GlobalHeader from "~/components/GlobalHeader.vue";
import PageHeader from "~/components/PageHeader.vue";
import CalendarForm from "~/components/CalendarForm.vue";

@Component({
  components: { GlobalHeader, PageHeader, CalendarForm }
})
export default class extends Vue {}
</script>
