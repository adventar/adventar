<template>
  <img
    class="UserIcon"
    :src="src"
    :alt="`${user.name}'s icon`"
    :width="size"
    :height="size"
    @error="handleErrorImage()"
  />
</template>

<script lang="ts">
import { Component, Vue, Prop } from "nuxt-property-decorator";
import { User } from "~/types/adventar";
import defaultImage from "~/assets/no_image_user.png";

@Component
export default class extends Vue {
  @Prop() readonly user: User;
  @Prop() readonly size: number;

  loadError: boolean = false;

  get src() {
    if (this.user.iconUrl === "" || this.loadError) {
      return defaultImage;
    } else {
      return this.user.iconUrl;
    }
  }

  handleErrorImage() {
    this.loadError = true;
  }
}
</script>

<style scoped>
img {
  border-radius: 50%;
  vertical-align: middle;
}
</style>
