<template>
  <header class="GlobalHeader">
    <div class="inner">
      <h1 class="logo">
        <nuxt-link to="/"><img src="~/assets/logo.png" alt="Adventar" width="220" height="28"/></nuxt-link>
      </h1>
      <no-ssr>
        <div class="right">
          <div v-if="$store.state.user">
            <span role="button" @click.stop="showDropdown = true" class="menuBtn">
              <UserIcon class="userIcon" :user="$store.state.user" size="28" />
              <font-awesome-icon icon="sort-down" />
            </span>
            <div class="dropdown is-login" v-if="showDropdown" @click.stop>
              <ul>
                <li v-if="calendarCreatable">
                  <nuxt-link to="/calendars/new">カレンダーを作る</nuxt-link>
                </li>
                <li>
                  <nuxt-link :to="`/users/${$store.state.user.id}`">自分のカレンダー</nuxt-link>
                </li>
                <li>
                  <nuxt-link to="/setting">ユーザー設定</nuxt-link>
                </li>
                <li>
                  <span role="button" @click="logout()">ログアウト</span>
                </li>
              </ul>
            </div>
          </div>
          <div v-else>
            <span role="button" @click.stop="showDropdown = true" class="menuBtn is-signin">
              <font-awesome-icon icon="sign-in-alt" />
            </span>
            <div class="dropdown" v-if="showDropdown" @click.stop>
              <ul>
                <li>
                  <span role="button" @click="login('google')">
                    <font-awesome-icon :icon="['fab', 'google']" /> Google でログイン
                  </span>
                </li>
                <li>
                  <span role="button" @click="login('github')">
                    <font-awesome-icon :icon="['fab', 'github']" /> GitHub でログイン
                  </span>
                </li>
                <li>
                  <span role="button" @click="login('twitter')">
                    <font-awesome-icon :icon="['fab', 'twitter']" /> Twitter でログイン
                  </span>
                </li>
                <li>
                  <span role="button" @click="login('facebook')">
                    <font-awesome-icon :icon="['fab', 'facebook']" /> Facebook でログイン
                  </span>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </no-ssr>
    </div>
  </header>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { loginWithFirebase, logoutWithFirebase } from "~/lib/Auth";
import { getCalendarCreatable } from "~/lib/Configuration";
import UserIcon from "~/components/UserIcon.vue";

@Component({
  components: { UserIcon }
})
export default class extends Vue {
  showDropdown = false;
  calendarCreatable = getCalendarCreatable();

  mounted() {
    document.addEventListener("click", this.handleClickDocument);
  }

  destroyed() {
    document.removeEventListener("click", this.handleClickDocument);
  }

  handleClickDocument() {
    this.showDropdown = false;
  }

  login(provider) {
    loginWithFirebase(provider);
  }

  logout() {
    this.$router.push("/");
    logoutWithFirebase();
  }
}
</script>

<style lang="scss" scoped>
.GlobalHeader {
  background-color: #fff;
}
.inner {
  max-width: $content-max-width;
  margin: 0 auto;
  position: relative;
  padding: 5px 12px;
}

.right {
  position: absolute;
  right: 15px;
  top: 18px;
}

.logo {
  margin: 0;
  padding: 10px 0;
  font-size: 24px;
  font-weight: bold;
}
.logo a {
  color: #e4523d;
  text-transform: uppercase;
  text-decoration: none;
}

.logo img {
  width: 165px;
  height: 21px;
}

.menuBtn {
  color: #666;
  cursor: pointer;
  display: block;
  padding-bottom: 10px;
}

.menuBtn.is-signin {
  font-size: 20px;
}

.menuBtn:hover {
  color: #000;
}

.menuBtn .userIcon {
  margin-right: 5px;
}

.dropdown {
  position: absolute;
  width: 100%;
  z-index: 1;
}

.dropdown ul {
  /* display: none; */
  border: 1px solid #dadada;
  border-radius: 3px;
  background: #fff;
  width: 180px;
  margin: 0;
  padding: 8px 0;
  font-size: 14px;
  float: right;
}

.dropdown.is-login ul {
  width: 140px;
}

.dropdown li {
  margin: 0;
  padding: 0;
  list-style: none;
}

.dropdown li svg {
  margin-right: 5px;
}

.dropdown li [role="button"],
.dropdown li a {
  display: block;
  color: #666;
  font-size: 13px;
  padding: 5px 10px;
  text-decoration: none;
  cursor: pointer;

  &:hover {
    color: #fff;
    background: #e45541;
  }
}

@media (min-width: $mq-break-small) {
  .inner {
    padding: 20px 12px;
  }

  .right {
    top: 32px;
    right: 50px;
  }

  .logo img {
    width: 220px;
    height: 28px;
  }
}
</style>
