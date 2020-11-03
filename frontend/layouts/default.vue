<template>
  <div>
    <GlobalHeader />
    <nuxt />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import GlobalHeader from "~/components/GlobalHeader.vue";
@Component({
  components: { GlobalHeader }
})
export default class extends Vue {}
</script>

<style lang="scss">
body {
  margin: 0;
  padding: 0;
  font-family: "Hiragino Kaku Gothic Pro", Meiryo, "MS PGothic", Verdana, Arial, sans-serif;
  background-color: #f1f1f1;
}

.mainInner {
  max-width: $content-max-width;
  padding: 30px 12px;
  margin: 0 auto;
}

a {
  color: #e45541;
}

a:hover {
  color: #ef8171;
}

button {
  background-color: transparent;
  border: none;
  padding: 0;
  margin: 0;
  text-align: inherit;
  color: inherit;
  font: inherit;
  appearance: none;

  &:focus {
    outline-width: 0;
  }

  &:focus-visible {
    outline-width: initial;
  }
}
</style>
