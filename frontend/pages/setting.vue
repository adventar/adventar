<template>
  <div>
    <PageHeader>ユーザー設定</PageHeader>

    <main>
      <div class="mainInner">
        <section>
          <SectionHeader>表示名</SectionHeader>
          <div v-if="!editmode">
            <button class="name" @click="handleClickEdit">
              {{ $store.state.user.name }}
              <font-awesome-icon icon="edit" />
            </button>
          </div>
          <div v-if="editmode">
            <form class="inputForm" @submit="handleSubmit">
              <input type="text" ref="inputName" :value="$store.state.user.name" />
              <button type="submit">Submit</button>
              <button class="cancel" @click="editmode = false">Cancel</button>
            </form>
          </div>
        </section>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { updateUser } from "~/lib/GrpcClient";
import { getToken } from "~/lib/Auth";
import PageHeader from "~/components/PageHeader.vue";
import SectionHeader from "~/components/SectionHeader.vue";

@Component({
  components: { PageHeader, SectionHeader },
  middleware: "requireUser"
})
export default class extends Vue {
  editmode = false;

  async handleSubmit(e) {
    e.preventDefault();
    const token = await getToken();
    const user = await updateUser((this.$refs.inputName as any).value, token);
    this.$store.commit("setUser", user);
    this.editmode = false;
  }

  handleClickEdit() {
    this.editmode = true;
    setTimeout(() => {
      (this.$refs.inputName as any).focus();
    });
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
