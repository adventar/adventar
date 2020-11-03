<template>
  <div>
    <PageHeader>ヘルプ</PageHeader>
    <main>
      <div class="mainInner">
        <details id="help-whats-adventar" class="section">
          <summary class="summary">Adventarとは？</summary>
          <div class="sectionBody">
            <p>
              Advent
              Calendarは本来、12月1日から24日までクリスマスを待つまでに1日に1つ、穴が空けられるようになっているカレンダーです。WebでのAdvent
              Calendarは、その風習に習い、12月1日から25日まで1日に1つ、みんなで記事を投稿していくというイベントです。
            </p>
            <p>AdventarではAdvent Calendarの作成や登録などを行うことができます。</p>
          </div>
        </details>

        <details id="help-calendar-creation" class="section">
          <summary class="summary">カレンダーを作成できる期間</summary>
          <div class="sectionBody">
            <p>カレンダーはその年の11月1日から12月25日までの期間に作成できます。</p>
          </div>
        </details>

        <details id="help-entry-reservation" class="section">
          <summary class="summary">予約投稿</summary>
          <div class="sectionBody">
            <p>
              投稿予定日より前にURLを投稿することができます。その場合、登録日になるまで記事のURLやタイトルなどは自分以外のユーザーには表示されません。
            </p>
          </div>
        </details>

        <details id="help-force-cancel" class="section">
          <summary class="summary">登録の強制解除</summary>
          <div class="sectionBody">
            <p>
              登録日を過ぎても投稿がない場合、カレンダーを作成したユーザーに限り、その登録を削除する権限があります。
            </p>
            <p>
              削除できるのは、投稿日の翌々日以降です。例えば12/3の登録に対して投稿がない場合、12/5の0時以降は削除可能になります。
            </p>
          </div>
        </details>

        <details id="help-creators" class="section">
          <summary class="summary">製作者について</summary>
          <div class="sectionBody">
            <p>
              Adventar は <a href="https://twitter.com/hokaccha">@hokaccha</a>,
              <a href="https://twitter.com/june29">@june29</a>,
              <a href="https://twitter.com/ayumiko">@ayumiko</a>
              が作成しています。お気軽にフィードバックなどお願いします。
            </p>
            <p>
              また、Adventar のソースコードは
              <a href="https://github.com/adventar/adventar">GitHub</a> でオープンソースとして公開されています。
            </p>
          </div>
        </details>

        <details id="help-privacy" class="section">
          <summary class="summary">プライバシーポリシー</summary>
          <div class="sectionBody"><nuxt-link to="/privacy">プライバシーポリシー</nuxt-link>を参照してください。</div>
        </details>

        <details id="help-inquiry" class="section">
          <summary class="summary">その他のお問い合わせ</summary>
          <div class="sectionBody">
            使い方で不明な点や機能要望などがありましたら
            <a href="https://twitter.com/hokaccha">@hokaccha</a> までお気軽にどうぞ。
          </div>
        </details>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import PageHeader from "~/components/PageHeader.vue";

@Component({
  components: { PageHeader }
})
export default class extends Vue {
  mounted() {
    const el = document.getElementById(location.hash.slice(1));
    if (el) {
      el.setAttribute("open", "open");
    }
  }
}
</script>

<style lang="scss" scoped>
.mainInner {
  padding-top: 5px;
}

.section {
  padding: 0 5px;
}

.summary {
  position: relative;
  font-size: 18px;
  font-weight: bold;
  color: #666;
  margin: 20px 0;
  white-space: nowrap;
  overflow: hidden;
  cursor: pointer;

  &:focus {
    outline-width: 0;
  }

  &:focus-visible {
    outline-width: initial;
  }
}

.summary:after {
  content: "";
  display: inline-block;
  height: 1px;
  width: 100%;
  background: #fcfcfc;
  vertical-align: middle;
  margin-left: 20px;
  box-shadow: 0 -1px 0 #d8d8d8;
}

.sectionBody {
  font-size: 14px;
}
</style>
