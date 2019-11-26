import { User } from "~/types/adventar";
import { saveUser } from "~/lib/Auth";

export function state() {
  return { user: null, isProcessingSignin: false };
}

export const mutations = {
  setUser(state, user: User | null) {
    saveUser(user);
    state.user = user;
    state.isProcessingSignin = false;
  },

  setProcessingSignin(state) {
    state.isProcessingSignin = true;
  }
};
