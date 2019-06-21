import { User } from "~/types/adventar";
import { saveUser } from "~/lib/Auth";

export function state() {
  return { user: null };
}

export const mutations = {
  setUser(state, user: User | null) {
    saveUser(user);
    state.user = user;
  }
};
