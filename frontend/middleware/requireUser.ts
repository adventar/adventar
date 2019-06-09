export default function({ store, redirect }) {
  if (!store.state.user) {
    // TODO: display error page
    redirect("/");
  }
}
