<template>
  <header class="GlobalHeader">
    <div class="GlobalHeader-inner">
      <h1 class="GlobalHeader-logo"><nuxt-link to="/">Adventar</nuxt-link></h1>
      <no-ssr>
        <div class="GlobalHeader-right">
          <div v-if="$store.state.user">
            <span role="button" @click.stop="showDropdown = true" class="GlobalHeader-menuBtn">
              <img :src="$store.state.user.iconUrl" class="GlobalHeader-userIcon" width="25" height="25" />
              {{ $store.state.user.name }}
              <font-awesome-icon icon="sort-down" />
            </span>
            <div class="GlobalHeader-dropdown is-login" v-if="showDropdown" @click.stop>
              <ul>
                <li v-if="calendarCreatable">
                  <nuxt-link to="/calendars/new">カレンダーを作る</nuxt-link>
                </li>
                <li>
                  <nuxt-link :to="`/users/${$store.state.user.id}`">マイページ</nuxt-link>
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
            <span role="button" @click.stop="showDropdown = true" class="GlobalHeader-menuBtn">
              Log In <font-awesome-icon icon="sign-in-alt" />
            </span>
            <div class="GlobalHeader-dropdown" v-if="showDropdown" @click.stop>
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

@Component
export default class extends Vue {
  showDropdown = false;
  calendarCreatable = getCalendarCreatable();

  // created() {
  //   if (process.server) return;
  //   // ????
  //   // this.$store.commit("setUser", JSON.parse(localStorage.getItem("adventar.user") || ""));
  //   console.log("crewate");
  //   console.log(this.$store.state.user);
  // }

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

<style scoped>
.GlobalHeader {
}

.GlobalHeader-inner {
  max-width: 1000px;
  padding: 6px 12px;
  margin: 0 auto;
  position: relative;
}

.GlobalHeader-right {
  position: absolute;
  right: 10px;
  top: 18px;
}

.GlobalHeader-logo {
  margin: 0;
  padding: 10px 0;
  font-size: 24px;
  font-weight: bold;
}
.GlobalHeader-logo a {
  color: #e4523d;
  text-transform: uppercase;
  text-decoration: none;
}

.GlobalHeader-menuBtn {
  color: #666;
  cursor: pointer;
  display: block;
  padding-bottom: 10px;
}

.GlobalHeader-menuBtn:hover {
  color: #000;
}

.GlobalHeader-menuBtn .GlobalHeader-userIcon {
  border-radius: 50%;
  vertical-align: middle;
  margin-right: 5px;
}

.GlobalHeader-dropdown {
  position: absolute;
  width: 100%;
  z-index: 1;
}

.GlobalHeader-dropdown ul {
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

.GlobalHeader-dropdown.is-login ul {
  width: 140px;
}

.GlobalHeader-dropdown li {
  margin: 0;
  padding: 0;
  list-style: none;
}

.GlobalHeader-dropdown li [role="button"],
.GlobalHeader-dropdown li a {
  display: block;
  color: #666;
  font-size: 13px;
  padding: 5px 10px;
  text-decoration: none;
  cursor: pointer;
}

.GlobalHeader-dropdown li [role="button"]:hover,
.GlobalHeader-dropdown li a:hover {
  color: #fff;
  background: #e45541;
}

.GlobalHeader-dropdown li svg {
  margin-right: 5px;
}
</style>
