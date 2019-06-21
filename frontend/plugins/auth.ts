import { firebase, restoreUser, handleAuthStateChanged } from "~/lib/Auth";

firebase.initializeApp({
  apiKey: "AIzaSyCDGROd2Z-2rWjHl-cDepyGnuQ975cpqQo",
  authDomain: "api-project-837626752936.firebaseapp.com",
  projectId: "api-project-837626752936"
});

export default function({ app }) {
  app.store.commit("setUser", restoreUser());
  handleAuthStateChanged(app.store);
}
