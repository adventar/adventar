import { getRedirectResult } from "~/lib/Auth";

export default function({ store }) {
  if (process.server) return;
  return getRedirectResult(store);
}
