import firebase from "firebase/app";
import "firebase/auth";
import { signIn } from "~/lib/grpc/Client";

firebase.initializeApp({
  apiKey: "AIzaSyCDGROd2Z-2rWjHl-cDepyGnuQ975cpqQo",
  authDomain: "api-project-837626752936.firebaseapp.com",
  projectId: "api-project-837626752936"
});

export function loginWithFirebase(provider: string): void {
  switch (provider) {
    case "google":
      firebase.auth().signInWithRedirect(new firebase.auth.GoogleAuthProvider());
      break;
    case "github":
      firebase.auth().signInWithRedirect(new firebase.auth.GithubAuthProvider());
      break;
    case "twitter":
      firebase.auth().signInWithRedirect(new firebase.auth.TwitterAuthProvider());
      break;
    case "facebook":
      firebase.auth().signInWithRedirect(new firebase.auth.FacebookAuthProvider());
      break;
    default:
      throw new Error("Invalid provider");
  }
}

export function getToken(): Promise<string | null> {
  const user = firebase.auth().currentUser;
  if (user) {
    return user.getIdToken(true);
  } else {
    return Promise.resolve(null);
  }
}

export default function({ app }) {
  firebase.auth().onAuthStateChanged(user => {
    if (!user) {
      app.store.commit("setUser", null);
      return;
    }

    user
      .getIdToken(true)
      .then(token => signIn(token))
      .then(user => app.store.commit("setUser", user))
      .catch(err => {
        app.store.commit("setUser", null);
        console.log(err);
      });
  });
}
