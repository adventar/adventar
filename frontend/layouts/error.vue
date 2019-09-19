<template>
  <div>
    <PageHeader>{{ notFound ? "404 Not Found" : "Error" }}</PageHeader>
    <main>
      <div>
        <div v-if="notFound">お探しのページは見つかりませんでした。</div>
        <div v-else>予期しないエラーが発生しました。</div>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "nuxt-property-decorator";
import PageHeader from "~/components/PageHeader.vue";
@Component({
  components: { PageHeader }
})
export default class extends Vue {
  @Prop() readonly error;

  get notFound(): boolean {
    return this.error.statusCode === 404;
  }
}
</script>
