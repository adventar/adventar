<template>
  <div>
    <PageHeader>サポーター</PageHeader>
    <main>
      <div class="mainInner">
        <section class="sectionItem">
          <SectionHeader>サポーター</SectionHeader>
          <ul class="supporters">
            <li>
              <a href="https://tech.cybozu.io/" rel="noopener noreferrer"
                ><img src="~/assets/support/cybozu.png" alt="Cybozu" width="301" height="273" class="logoImg"
              /></a>
            </li>
          </ul>
        </section>
      </div>
      <div class="mainInner">
        <section class="sectionItem">
          <SectionHeader>サポーターについて</SectionHeader>
          <p>
            Adventarは個人が趣味で運営しているサービスなのでこれまでもこれからも課金や広告などを導入することは考えていませんが、サーバーの運用費などをサポートしてくる企業や個人の方がおりましたら歓迎いたします。サイトに掲載するぐらいの見返りしかありませんが、興味がある方は以下までご連絡ください。
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
