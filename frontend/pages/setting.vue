<template>
  <div>
    <GlobalHeader />

    <PageHeader>ユーザー設定</PageHeader>

    <main>
      <div>
        <section>
          <SectionHeader>表示名</SectionHeader>
          <div v-if="!editmode">
            <span role="button" class="name" @click="editmode = true">
              {{ $store.state.user.name }}
              <font-awesome-icon icon="edit" />
            </span>
          </div>
          <div v-if="editmode">
            <form @submit.prevent class="inputForm">
              <input type="text" :value="$store.state.user.name" @change="onChangeName" />
              <button>Submit</button>
              <span role="button" class="cancel" @click="editmode = false">Cancel</span>
            </form>
          </div>
        </section>

        <section style="margin-top: 50px">
          <SectionHeader>画像</SectionHeader>
          <img :src="$store.state.user.iconUrl" width="80" height="80" style="border-radius: 80px" />
        </section>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { updateUser } from "~/lib/GrpcClient";
import { getToken } from "~/lib/Auth";
import GlobalHeader from "~/components/GlobalHeader.vue";
import PageHeader from "~/components/PageHeader.vue";
import SectionHeader from "~/components/SectionHeader.vue";

@Component({
  components: { GlobalHeader, PageHeader, SectionHeader },
  middleware: "requireUser"
})
export default class extends Vue {
  editmode = false;

  async onChangeName(e) {
    const token = await getToken();
    const user = await updateUser(e.target.value, token);
    this.$store.commit("setUser", user);
    this.editmode = false;
  }
}
</script>

<style scoped>
.name {
  color: #333;
  cursor: pointer;
}

.inputForm input[type="text"] {
  outline: none;
  font-size: 13px;
  padding: 3px;
  border: 1px solid #ccc;
  border-radius: 3px;
  width: 150px;
}

.inputForm button {
  font-size: 13px;
  border: 1px solid #ccc;
  border-radius: 3px;
  background-color: #efefef;
  margin-left: 10px;
}

.inputForm .cancel {
  display: inline-block;
  font-size: 12px;
  color: #79a0ff;
  cursor: pointer;
  margin-left: 10px;
}
</style>
