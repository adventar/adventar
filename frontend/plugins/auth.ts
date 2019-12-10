import { initFirebase, restoreUser, initAuth } from "~/lib/Auth";

initFirebase();

export default function({ app }) {
  app.store.commit("setUser", restoreUser());
  initAuth(app.store);
}
