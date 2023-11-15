import { initializeApp } from "firebase/app";
import { detect } from "detect-browser";
import {
  getAuth,
  getRedirectResult,
  onAuthStateChanged,
  signInWithRedirect,
  signInWithPopup,
  GoogleAuthProvider,
  GithubAuthProvider,
  TwitterAuthProvider,
  FacebookAuthProvider
} from "firebase/auth";
import { signIn } from "~/lib/GrpcClient";
import { User } from "~/types/adventar";

const USER_STORAGE_KEY = "adventar.user";
const SIGNIN_STORAGE_KEY = "adventar.signin";
const SIGNIN_STORAGE_VALUE = "1";

export function restoreUser(): User | null {
  const user = localStorage.getItem(USER_STORAGE_KEY);
  if (user === null) return null;

  try {
    return JSON.parse(user);
  } catch (err) {
    console.error(err);
    return null;
  }
}

export function saveUser(user: User | null): void {
  if (user === null) {
    localStorage.removeItem(USER_STORAGE_KEY);
  } else {
    localStorage.setItem(USER_STORAGE_KEY, JSON.stringify(user));
  }
}

export function initFirebase(): void {
  initializeApp({
    apiKey: process.env.FIREBASE_API_KEY,
    authDomain: process.env.FIREBASE_AUTH_DOMAIN,
    projectId: process.env.FIREBASE_PROJECT_ID
  });
}

export function loginWithFirebase(provider: string): void {
  sessionStorage.setItem(SIGNIN_STORAGE_KEY, SIGNIN_STORAGE_VALUE);
  const auth = getAuth();
  const browser = detect();
  const usePopup = browser?.os === "iOS" || browser?.name === "safari" || browser?.name === "firefox";
  const signIn = usePopup ? signInWithPopup : signInWithRedirect;
  switch (provider) {
    case "google":
      signIn(auth, new GoogleAuthProvider());
      break;
    case "github":
      signIn(auth, new GithubAuthProvider());
      break;
    case "twitter":
      signIn(auth, new TwitterAuthProvider());
      break;
    case "facebook":
      signIn(auth, new FacebookAuthProvider());
      break;
    default:
      throw new Error("Invalid provider");
  }
}

export async function logoutWithFirebase() {
  await getAuth().signOut();
}

export function getToken(): Promise<string> {
  const user = getAuth().currentUser;
  if (user) {
    return user.getIdToken();
  } else {
    return Promise.reject(new Error("currentUser is null"));
  }
}

export function initAuth(store) {
  return Promise.all([getAuthRedirectResult(), handleAuthStateChanged(store)]);
}

function isProcessingSignin() {
  return sessionStorage.getItem(SIGNIN_STORAGE_KEY) === SIGNIN_STORAGE_VALUE;
}

function getAuthRedirectResult() {
  if (!isProcessingSignin()) return;

  return getRedirectResult(getAuth()).catch(err => {
    const COOKIE_ERROR_MSG =
      "third-party cookie の設定が無効になってる可能性があります。ブラウザの設定をご確認ください。";
    const msg = err.code === "auth/web-storage-unsupported" ? COOKIE_ERROR_MSG : err.message;
    alert(`ログインに失敗しました。\n${msg}`);
    console.error(err);
  });
}

let listenedAuthStateChanged = false;
function handleAuthStateChanged(store): Promise<void> {
  if (listenedAuthStateChanged === true) return Promise.resolve();
  listenedAuthStateChanged = true;

  if (isProcessingSignin()) {
    store.commit("setProcessingSignin");
  }

  return new Promise((resolve, reject) => {
    onAuthStateChanged(getAuth(), user => {
      sessionStorage.removeItem(SIGNIN_STORAGE_KEY);

      if (!user) {
        store.commit("setUser", null);
        return resolve();
      }

      const iconUrl = (user.providerData && user.providerData[0] && user.providerData[0]!.photoURL) || "";

      user
        .getIdToken()
        .then(token => signIn(token, iconUrl))
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
