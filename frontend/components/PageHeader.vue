<template>
  <div class="PageHeader">
    <h1><slot></slot></h1>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";

@Component
export default class extends Vue {}
</script>

<style lang="scss" scoped>
.PageHeader {
  background: #80b2d0;
  padding: 15px 0;
  color: #fff;
}

.PageHeader h1 {
  max-width: 1000px;
  margin: 0 auto;
  padding: 0 12px;
  position: relative;
  font-size: 20px;
  font-weight: normal;
  color: #fff;
}

@media (min-width: $mq-break-small) {
  .PageHeader h1 {
    font-size: 24px;
  }
}
</style>
