import { getRedirectResult } from "~/plugins/firebase";

export default function({ store }) {
  return getRedirectResult(store);
}
