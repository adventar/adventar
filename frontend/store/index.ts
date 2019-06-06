export function state() {
  return {
    user: null
  };
}

export const mutations = {
  setUser(state, user) {
    state.user = user;
  }
};
