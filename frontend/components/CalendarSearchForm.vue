<template>
  <form @submit.prevent="handleSubmit">
    <input type="text" v-model="query" />
    <button><font-awesome-icon icon="search" /></button>
  </form>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "nuxt-property-decorator";

@Component
export default class extends Vue {
  @Prop(Number) readonly year!: number;
  @Prop(String) readonly defaultQuery: string;

  query: string;

  created() {
    this.query = this.defaultQuery || "";
  }

  handleSubmit() {
    this.$router.push(`/calendars?year=${this.year}&query=${this.query}`);
  }
}
</script>

<style scoped>
form {
  position: relative;
  display: inline-block;
}

input[type="text"] {
  border-radius: 15px;
  width: 200px;
  padding: 6px 40px 6px 15px;
  border: 1px solid #d0d0d0;
  font-size: 12px;
  color: #666;
  outline: none;
}

button {
  position: absolute;
  right: 12px;
  top: 5px;
  color: #666;
  background: none;
  border: none;
  border-radius: 6px;
  font-size: 16px;
  cursor: pointer;
}

button:hover {
  color: #2f7bff;
  background-color: #d6e5ff;
}
</style>
