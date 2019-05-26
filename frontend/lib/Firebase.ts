import firebase from "firebase";
import "firebase/auth";

export function initFirebase(): void {
  firebase.initializeApp({
    apiKey: "AIzaSyCDGROd2Z-2rWjHl-cDepyGnuQ975cpqQo",
    authDomain: "api-project-837626752936.firebaseapp.com",
    projectId: "api-project-837626752936"
  });

  firebase.auth().onAuthStateChanged(user => {
    console.log(user);
    if (user) {
      const currentUser = firebase.auth().currentUser;
      if (currentUser) {
        currentUser
          .getIdToken(/* forceRefresh */ true)
          .then(idToken => {
            console.log(idToken);
          })
          .catch(error => {
            console.log(error);
          });
      }
    }
  });

  firebase
    .auth()
    .getRedirectResult()
    .then(result => {
      console.log(result.credential);
    })
    .catch(error => {
      console.log(error);
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
