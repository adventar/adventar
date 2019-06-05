import firebase from "firebase/app";
import "firebase/auth";
import { signIn, User } from "~/lib/grpc/Client";

let initialized = false;

export function initializeFirebaseApp(): void {
  if (initialized) return;

  firebase.initializeApp({
    apiKey: "AIzaSyCDGROd2Z-2rWjHl-cDepyGnuQ975cpqQo",
    authDomain: "api-project-837626752936.firebaseapp.com",
    projectId: "api-project-837626752936"
  });

  initialized = true;
}

export function auth(): Promise<User | null> {
  return new Promise((resolve, reject) => {
    firebase.auth().onAuthStateChanged(user => {
      if (user) {
        user
          .getIdToken(/* forceRefresh */ true)
          .then(token => signIn(token))
          .then(user => resolve(user))
          .catch(err => reject(err));
      } else {
        resolve(null);
      }
    });
  });
}

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
