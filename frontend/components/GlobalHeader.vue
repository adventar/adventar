<template>
  <header class="GlobalHeader">
    <div class="inner">
      <h1 class="logo">
        <nuxt-link to="/"><img src="~/assets/logo.png" alt="Adventar" width="220" height="28"/></nuxt-link>
      </h1>
      <div class="right">
        <no-ssr>
          <button class="menuBtn" :aria-expanded="isShownDropdown ? 'true' : 'false'" @click.stop="showDropdown()">
            <UserIcon v-if="$store.state.user" class="userIcon" :user="$store.state.user" size="28" />
            <font-awesome-icon v-else-if="$store.state.isProcessingSignin" icon="circle-notch" spin />
            <font-awesome-icon v-else icon="bars"></font-awesome-icon>
          </button>
          <div v-if="isShownDropdown" class="dropdown" @click.stop>
            <ul v-if="$store.state.user" class="loginMenu">
              <li class="user">
                <UserIcon class="userIcon" :user="$store.state.user" size="28" />
                {{ $store.state.user.name }}
              </li>
              <li>
                <nuxt-link to="/new" @click.native="hideDropdown()">
                  <font-awesome-icon icon="calendar-plus" />
                  カレンダーを作る
                </nuxt-link>
              </li>
              <li>
                <nuxt-link :to="`/users/${$store.state.user.id}`" @click.native="hideDropdown()">
                  <font-awesome-icon icon="user" /> マイページ
                </nuxt-link>
              </li>
              <li>
                <nuxt-link to="/setting" @click.native="hideDropdown()">
                  <font-awesome-icon icon="cog" /> 設定
                </nuxt-link>
              </li>
              <li>
                <button @click.native="hideDropdown()" @click="logout()">
                  <font-awesome-icon icon="sign-out-alt" /> ログアウト
                </button>
              </li>
            </ul>
            <ul v-if="!$store.state.user" class="loginMenu">
              <li>
                <button @click.native="hideDropdown()" @click="login('google')">
                  <font-awesome-icon :icon="['fab', 'google']" /> Google でログイン
                </button>
              </li>
              <li>
                <button @click.native="hideDropdown()" @click="login('github')">
                  <font-awesome-icon :icon="['fab', 'github']" /> GitHub でログイン
                </button>
              </li>
              <li>
                <button @click.native="hideDropdown()" @click="login('twitter')">
                  <font-awesome-icon :icon="['fab', 'twitter']" /> Twitter でログイン
                </button>
              </li>
              <li>
                <button @click.native="hideDropdown()" @click="login('facebook')">
                  <font-awesome-icon :icon="['fab', 'facebook']" /> Facebook でログイン
                </button>
              </li>
            </ul>
            <ul class="generalMenu">
              <li>
                <nuxt-link to="/archive" @click.native="hideDropdown()">
                  <font-awesome-icon icon="calendar-minus" /> 過去のカレンダー
                </nuxt-link>
              </li>
              <li>
                <nuxt-link to="/help" @click.native="hideDropdown()">
                  <font-awesome-icon icon="question-circle" /> ヘルプ
                </nuxt-link>
              </li>
            </ul>
          </div>
        </no-ssr>
      </div>
    </div>
  </header>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { loginWithFirebase, logoutWithFirebase } from "~/lib/Auth";
import UserIcon from "~/components/UserIcon.vue";

@Component({
  components: { UserIcon }
})
export default class extends Vue {
  isShownDropdown = false;

  mounted() {
    document.addEventListener("click", this.handleClickDocument);
    document.addEventListener("keyup", this.handleKeyupDocument);
  }

  destroyed() {
    document.removeEventListener("click", this.handleClickDocument);
    document.removeEventListener("keyup", this.handleKeyupDocument);
  }

  handleClickDocument() {
    this.hideDropdown();
  }

  handleKeyupDocument(e) {
    if (e.key === "Escape") {
      this.hideDropdown();
    }
  }

  showDropdown() {
    this.isShownDropdown = true;
  }

  hideDropdown() {
    this.isShownDropdown = false;
  }

  login(provider) {
    loginWithFirebase(provider);
  }

  async logout() {
    this.$router.push("/");
    this.isShownDropdown = false;
    await logoutWithFirebase();
    window.alert("ログアウトしました。");
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
  top: 8px;
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
  background: none;
  border: none;
  color: #333;
  cursor: pointer;
  display: block;
  padding: 10px;
  font-size: 20px;
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
  border: 1px solid #dadada;
  border-radius: 3px;
  background: #fff;
  width: 200px;
  margin: 0;
  padding: 0;
  font-size: 14px;
  float: right;

  &.loginMenu {
    border-radius: 3px 3px 0 0;
  }

  &.generalMenu {
    border-radius: 0 0 3px 3px;
    border-top: none;
  }
}

.dropdown li {
  margin: 0;
  padding: 0;
  list-style: none;
}

.dropdown li.user {
  padding: 5px 10px;
  margin-bottom: 5px;
  background-color: #eaeaea;
}

.dropdown li svg {
  margin-right: 5px;
}

.dropdown li button,
.dropdown li a {
  display: block;
  color: #666;
  font-size: 13px;
  padding: 10px 10px;
  text-decoration: none;
  width: 100%;
  box-sizing: border-box;

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
    top: 27px;
    right: 12px;
  }

  .logo img {
    width: 220px;
    height: 28px;
  }
}
</style>
