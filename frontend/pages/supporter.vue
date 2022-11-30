<template>
  <div>
    <PageHeader>サポーター</PageHeader>
    <main>
      <div class="mainInner">
        <section class="sectionItem">
          <SectionHeader>サポーターについて</SectionHeader>
          <p>
            Adventarは個人が趣味で運営しているサービスなのでこれまでもこれからも課金や広告などを導入することは考えていませんが、サーバーの運用費や開発費などをサポートしてくる企業や個人の方がおりましたら歓迎いたします。
          </p>
          <p>
            サポートは<a href="https://github.com/sponsors/hokaccha">GitHub Sponsor</a>から受け付けています。 GitHub
            Sponsor以外で支払いがしたい、サポートしたのでこのページにロゴを掲載してほしいなど、サポーターに関するお問い合わせは以下にご連絡ください。
          </p>
          <ul>
            <li>Email: k.hokamura@gmail.com</li>
            <li>Twitter: <a href="https://twitter.com/hokaccha">@hokaccha</a></li>
          </ul>
        </section>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import PageHeader from "~/components/PageHeader.vue";
import SectionHeader from "~/components/SectionHeader.vue";

@Component({
  components: { PageHeader, SectionHeader }
})
export default class extends Vue {}
</script>

<style scoped>
.sectionItem {
  margin: 30px 0;
}

.supporters {
  padding: 0;
  margin: 30px 0 0 0;
  list-style: none;
}

.logoImg {
  width: 200px;
  height: auto;
}
</style>
