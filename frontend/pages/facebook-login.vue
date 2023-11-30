<template>
  <div>
    <PageHeader>Facebookログインについて</PageHeader>
    <main>
      <div class="mainInner">
        <p>
          Facebookログインの機能を提供するのに、<a
            href="https://developers.facebook.com/blog/post/2023/02/01/developer-platform-requiring-business-verification-for-advanced-access/?locale=ja_JP"
            >ビジネスアカウントの認証が必要になりました</a
          >。Adventarは個人で運営しているサービスのため、ビジネスアカウントの認証を通すのが難しく、Facebookログインの機能の提供を終了することにしました。
        </p>
        <p>
          すでにFacebookログインでAdventarに登録しており、データの変更が必要な場合は、<a
            href="https://twitter.com/hokaccha"
            >@hokaccha</a
          >宛にDMをお願いします。
        </p>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import PageHeader from "~/components/PageHeader.vue";
import SectionHeader from "~/components/SectionHeader.vue";

@Component({
  components: { PageHeader }
})
export default class extends Vue {}
</script>
