<template>
  <div>
    <GlobalHeader />
    <h2>Setting</h2>
    <ul>
      <li>id: {{ $store.state.user.id }}</li>
      <li><img :src="$store.state.user.iconUrl" width="50" height="50" /></li>
      <li v-if="!editmode">{{ $store.state.user.name }} <button @click="editmode = true">Edit</button></li>
      <li v-if="editmode"><input type="text" :value="$store.state.user.name" @change="onChangeName" /></li>
    </ul>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { updateUser } from "~/lib/GrpcClient";
import { getToken } from "~/plugins/auth";
import GlobalHeader from "~/components/GlobalHeader.vue";

@Component({
  components: { GlobalHeader },
  middleware: "requireUser"
})
export default class extends Vue {
  editmode = false;

  async onChangeName(e) {
    const token = await getToken();
    if (token === null) {
      throw new Error("Token is null");
    }
    const user = await updateUser(e.target.value, token);
    this.$store.commit("setUser", user);
    this.editmode = false;
  }
}
</script>
