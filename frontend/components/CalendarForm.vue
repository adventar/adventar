<template>
  <form @submit.prevent="handleSubmit()">
    <section>
      <SectionHeader>タイトル</SectionHeader>
      <input v-model="title" type="text" /> Advent Calendar {{ year }}
    </section>

    <section>
      <SectionHeader>概要</SectionHeader>
      <textarea v-model="description"></textarea>
      <p class="note">Markdown記法が使えます。</p>
    </section>

    <div v-if="calendar">
      <button type="submit">更新</button>
      <br />
      <button type="button" class="deleteBtn" @click="handleDelete()">削除</button>
    </div>
    <div v-else>
      <button type="submit">作成</button>
    </div>
  </form>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "nuxt-property-decorator";
import { Calendar } from "~/types/adventar";
import { createCalendar, updateCalendar, deleteCalendar } from "~/lib/GrpcClient";
import { getCurrentYear } from "~/lib/Configuration";
import { getToken } from "~/lib/Auth";
import SectionHeader from "~/components/SectionHeader.vue";

@Component({
  components: { SectionHeader }
})
export default class extends Vue {
  @Prop() readonly calendar: Calendar;

  title: string = "";
  description: string = "";
  year: number;

  created() {
    if (this.calendar) {
      this.title = this.calendar.title;
      this.description = this.calendar.description;
      this.year = this.calendar.year;
    } else {
      this.year = getCurrentYear();
    }
  }

  handleSubmit() {
    if (this.calendar) {
      this.update();
    } else {
      this.create();
    }
  }

  async create() {
    const token = await getToken();
    const calendarId = await createCalendar({
      title: this.title,
      description: this.description,
      token: token
    });
    this.$router.push(`/calendars/${calendarId}`);
  }

  async update() {
    const token = await getToken();
    console.log(this.calendar.id, this.title, this.description);
    await updateCalendar({
      id: this.calendar.id,
      title: this.title,
      description: this.description,
      token: token
    });
    this.$router.push(`/calendars/${this.calendar.id}`);
  }

  async handleDelete() {
    if (!window.confirm("削除しますか？")) return;
    const token = await getToken();
    await deleteCalendar({ id: this.calendar.id, token });
    alert("削除しました");
    this.$router.push("/");
  }
}
</script>

<style scoped>
.note {
  margin-top: 10px;
  font-size: 13px;
  color: #666;
}

input[type="text"] {
  font-size: 16px;
  width: 300px;
  padding: 5px;
  border: 1px solid #ccc;
}

textarea {
  width: 500px;
  height: 200px;
  font-size: 14px;
  padding: 5px;
  border: 1px solid #ccc;
}

button[type="submit"] {
  font-size: 16px;
  color: #fff;
  border: 2px solid #28a745;
  background-color: #28a745;
  padding: 8px 30px;
  border-radius: 10px;
  cursor: pointer;
}

button.deleteBtn {
  font-size: 16px;
  color: #fff;
  border: 2px solid #dc3545;
  background-color: #dc3545;
  padding: 8px 30px;
  border-radius: 10px;
  cursor: pointer;
  margin-top: 20px;
}
</style>
