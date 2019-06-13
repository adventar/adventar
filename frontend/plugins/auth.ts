import firebase from "firebase/app";
import "firebase/auth";
import { signIn } from "~/lib/GrpcClient";

firebase.initializeApp({
  apiKey: "AIzaSyCDGROd2Z-2rWjHl-cDepyGnuQ975cpqQo",
  authDomain: "api-project-837626752936.firebaseapp.com",
  projectId: "api-project-837626752936"
});

const SIGNIN_STORAGE_KEY = "adventar.signin";
const SIGNIN_STORAGE_VALUE = "1";

export function loginWithFirebase(provider: string): void {
  sessionStorage.setItem(SIGNIN_STORAGE_KEY, SIGNIN_STORAGE_VALUE);
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

export function logoutWithFirebase(): Promise<void> {
  return firebase.auth().signOut();
}

export function getToken(): Promise<string> {
  const user = firebase.auth().currentUser;
  if (user) {
    return user.getIdToken();
  } else {
    return Promise.reject(new Error("currentUser is null"));
  }
}

let listenedAuthStateChanged = false;
function handleAuthStateChanged(store): Promise<void> {
  if (listenedAuthStateChanged === true) return Promise.resolve();
  listenedAuthStateChanged = true;

  return new Promise((resolve, reject) => {
    firebase.auth().onAuthStateChanged(user => {
      sessionStorage.removeItem(SIGNIN_STORAGE_KEY);

      if (!user) {
        store.commit("setUser", null);
        return resolve();
      }

      user
        .getIdToken()
        .then(token => signIn(token))
        .then(user => {
          store.commit("setUser", user);
          resolve();
        })
        .catch(err => {
          store.commit("setUser", null);
          console.error(err);
          reject(err);
        });
    });
  });
}

export function getRedirectResult(store) {
  if (sessionStorage.getItem(SIGNIN_STORAGE_KEY) !== SIGNIN_STORAGE_VALUE) {
    return;
  }

  const p1 = handleAuthStateChanged(store);
  const p2 = firebase
    .auth()
    .getRedirectResult()
    .catch(err => {
      alert("Login Failed 😫");
      console.error(err);
    });

  return Promise.all([p1, p2]);
}

export default function({ app }) {
  handleAuthStateChanged(app.store);
}
