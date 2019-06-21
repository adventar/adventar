import { initFirebase, restoreUser, handleAuthStateChanged } from "~/lib/Auth";

initFirebase();

export default function({ app }) {
  app.store.commit("setUser", restoreUser());
  handleAuthStateChanged(app.store);
}
