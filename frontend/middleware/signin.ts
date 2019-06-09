import { getRedirectResult } from "~/plugins/auth";

export default function({ store }) {
  return getRedirectResult(store);
}
