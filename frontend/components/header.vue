<template>
  <header class="Header">
    <div class="Header-inner">
      <h1 class="Header-logo"><nuxt-link to="/">Adventar</nuxt-link></h1>
      <div class="Header-right">
        <div v-if="$store.state.user">
          <span role="button" @click.stop="showDropdown = true" class="Header-menuBtn">
            <img :src="$store.state.user.iconUrl" class="Header-userIcon" width="25" height="25" />
            {{ $store.state.user.name }}
            <font-awesome-icon icon="sort-down" />
          </span>
          <div class="Header-dropdown is-login" v-if="showDropdown" @click.stop>
            <ul>
              <li>
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
          <span role="button" @click.stop="showDropdown = true" class="Header-menuBtn">
            Log In <font-awesome-icon icon="sign-in-alt" />
          </span>
          <div class="Header-dropdown" v-if="showDropdown" @click.stop>
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
    </div>
  </header>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { loginWithFirebase, logoutWithFirebase } from "~/plugins/auth";

@Component
export default class extends Vue {
  showDropdown = false;

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
.Header {
  border-bottom: 1px solid #ccc;
}

.Header-inner {
  max-width: 1000px;
  padding: 0 12px;
  margin: 0 auto;
  position: relative;
}

.Header-right {
  position: absolute;
  right: 10px;
  top: 15px;
}

.Header-logo {
  margin: 0;
  padding: 10px;
  font-size: 24px;
  font-weight: bold;
}
.Header-logo a {
  color: #e4523d;
  text-transform: uppercase;
  text-decoration: none;
}

.Header-menuBtn {
  color: #666;
  cursor: pointer;
  display: block;
  padding-bottom: 10px;
}

.Header-menuBtn:hover {
  color: #000;
}

.Header-menuBtn .Header-userIcon {
  border-radius: 50%;
  vertical-align: middle;
  margin-right: 5px;
}

.Header-dropdown {
  position: absolute;
  width: 100%;
  z-index: 1;
}

.Header-dropdown ul {
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

.Header-dropdown.is-login ul {
  width: 140px;
}

.Header-dropdown li {
  margin: 0;
  padding: 0;
  list-style: none;
}

.Header-dropdown li [role="button"],
.Header-dropdown li a {
  display: block;
  color: #666;
  font-size: 13px;
  padding: 5px 10px;
  text-decoration: none;
  cursor: pointer;
}

.Header-dropdown li [role="button"]:hover,
.Header-dropdown li a:hover {
  color: #fff;
  background: #e45541;
}

.Header-dropdown li svg {
  margin-right: 5px;
}
</style>
