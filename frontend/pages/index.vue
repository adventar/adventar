<template>
  <div>
    <advHeader />

    <h2>Index!</h2>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import advHeader from "~/components/header.vue";

@Component({
  components: { advHeader }
})
export default class extends Vue {}
</script>
