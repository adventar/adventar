<template>
  <div>
    <PageHeader>404 Not Found</PageHeader>
    <main>
      <div class="mainInner">
        <div>お探しのページは見つかりませんでした。</div>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import PageHeader from "~/components/PageHeader.vue";
@Component({
  components: { PageHeader }
})
export default class extends Vue {}
</script>
