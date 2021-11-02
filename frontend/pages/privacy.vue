<template>
  <div>
    <PageHeader>プライバシーポリシー</PageHeader>
    <main>
      <div class="mainInner">
        <p>Adventar（以下本サービス）ではプライバシーポリシーについて以下のように定めます。</p>

        <section class="sectionItem">
          <SectionHeader>個人情報の定義</SectionHeader>
          <p>
            本プライバシーポリシーにおいて、個人情報とは、個人情報保護法第2条第1項により定義された個人情報、すなわち、生存する個人に関する情報であって、当該情報に含まれる氏名、生年月日その他の記述等により特定の個人を識別することができるものを意味することとします。
          </p>
        </section>

        <section class="sectionItem">
          <SectionHeader>個人情報の利用目的</SectionHeader>
          <p>個人情報の利用にあたっては、本サービスが提供する目的の範囲を超えた個人情報の取扱いを行いません。</p>
        </section>

        <section class="sectionItem">
          <SectionHeader>個人情報の第三者への開示・提供の禁止</SectionHeader>
          <p>
            お預かりした個人情報を適切に管理し、次のいずれかに該当する場合を除き、個人情報を第三者に開示いたしません。
          </p>

          <ul>
            <li>ユーザーご自身の同意がある場合</li>
            <li>法令に基づき開示することが必要である場合</li>
          </ul>
        </section>

        <section class="sectionItem">
          <SectionHeader>個人情報の管理</SectionHeader>
          <p>
            本サービスでは、個人情報を正確かつ最新の状態に保ち、個人情報への不正アクセス・紛失・破損・改ざん・漏洩などを防止するため、セキュリティシステムの維持・管理体制の整備・安全対策を実施し個人情報の厳重な管理を行ないます。
          </p>
        </section>

        <section class="sectionItem">
          <SectionHeader>法令、規範の遵守と見直し</SectionHeader>
          <p>
            保有する個人情報に関して適用される日本の法令、その他規範を遵守するとともに、本ポリシーの内容を適宜見直し、その改善に努めます。
          </p>
        </section>
        <section class="sectionItem">
          <SectionHeader>個人情報の訂正および削除</SectionHeader>
          <p>
            ユーザーは本サービスの保有する自己の個人情報が誤った情報である場合には、本サービスに対して個人情報の訂正、追加または削除を請求することができます。本サービスは、ユーザーから請求を受けてその請求に応じる必要があると判断した場合には、遅滞なく、当該個人情報の訂正・削除等を行うものとします。
          </p>
        </section>
        <section class="sectionItem">
          <SectionHeader>問い合わせ先</SectionHeader>
          <p>
            個人情報の訂正や削除、その他のプライバシーに関するお問い合わせは以下までお願いします。<br />
            Email: k.hokamura@gmail.com
          </p>
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
</style>
