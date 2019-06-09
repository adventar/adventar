import { User } from "~/types/adventar";
const STORAGE_USER_KEY = "adventar.user";

export function state() {
  const u = localStorage.getItem(STORAGE_USER_KEY);
  return {
    user: u ? JSON.parse(u) : null
  };
}

export const mutations = {
  setUser(state, user: User | null) {
    localStorage.setItem(STORAGE_USER_KEY, JSON.stringify(user));
    state.user = user;
  }
};
