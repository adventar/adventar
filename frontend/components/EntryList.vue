<template>
  <ul class="EntryList">
    <li v-for="entry in calendar.entries" :key="entry.day" class="item">
      <div class="head">
        <div class="date">12/{{ entry.day }}</div>
        <div class="user">
          <UserIcon :user="entry.owner" size="24" />
          <nuxt-link :to="`/users/${entry.owner.id}`">{{ entry.owner.name }}</nuxt-link>
        </div>
      </div>
      <div v-if="entry.comment" class="comment"><font-awesome-icon icon="comment" /> {{ entry.comment }}</div>
      <div v-if="entry.url && !isFutureEntry(entry)" class="article">
        <div class="left">
          <div class="link">
            <a :href="entry.url">{{ entry.url }}</a>
          </div>
          <div v-if="entry.title && !isFutureEntry(entry)">{{ entry.title }}</div>
        </div>
        <div class="image" v-if="entry.imageUrl && !isFutureEntry(entry)">
          <img :src="entry.imageUrl" />
        </div>
      </div>
    </li>
  </ul>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "nuxt-property-decorator";
import UserIcon from "~/components/UserIcon.vue";
import { Calendar } from "~/types/adventar";
import { getToday } from "~/lib/Configuration";

@Component({
  components: { UserIcon }
})
export default class extends Vue {
  @Prop() readonly calendar: Calendar;

  isFutureEntry(entry): boolean {
    // TODO: Fix to JST
    return new Date(this.calendar.year, 11, entry.day) > getToday();
  }
}
</script>

<style lang="scss" scoped>
.EntryList {
  list-style: none;
  margin: 10px 0 0 0;
  padding: 0;
  background-color: #fff;
  color: #666;
}

.item {
  padding: 10px;
  border-top: 1px solid #e3e4e3;
  position: relative;
  font-size: 14px;
  clear: both;
}

.date {
  font-size: 18px;
  font-weight: bold;
  display: inline-block;
  vertical-align: middle;
}

.user {
  margin-left: 10px;
  display: inline-block;
}

.comment {
  margin-top: 10px;
}

.article {
  display: table;
  border-left: 5px solid #e3e4e3;
  margin-top: 10px;
  padding: 5px 0 5px 10px;
  width: 100%;
  box-sizing: border-box;
}

.article .left {
  display: table-cell;
  vertical-align: top;
  word-break: break-all;
}

.link {
  margin-bottom: 5px;
}

.image {
  display: table-cell;
  padding-left: 10px;
  vertical-align: top;
  width: 80px;

  img {
    width: 80px;
  }
}

@media (min-width: $mq-break-small) {
  .item {
    font-size: 16px;
  }

  .date {
    width: 55px;
  }

  .comment {
    margin-top: 15px;
    margin-left: 73px;
  }

  .comment .fa-comment {
    font-size: 18px;
  }

  .article {
    margin-left: 77px;
    margin-top: 15px;
    border-left-width: 7px;
    width: calc(100% - 77px);
  }

  .image {
    width: 120px;

    img {
      width: 120px;
    }
  }
}
</style>
