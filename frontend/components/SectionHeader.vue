<template>
  <h2><slot></slot></h2>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";

@Component
export default class extends Vue {}
</script>

<style scoped>
h2 {
  position: relative;
  font-size: 18px;
  font-weight: bold;
  color: #666;
  margin: 20px 0;
  white-space: nowrap;
  overflow: hidden;
}

h2:after {
  content: "";
  display: inline-block;
  height: 1px;
  width: 100%;
  background: #fcfcfc;
  vertical-align: middle;
  margin-left: 20px;
  box-shadow: 0 -1px 0 #d8d8d8;
}
</style>
