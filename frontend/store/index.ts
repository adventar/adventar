export function state() {
  return {
    user: false
  };
}

export const mutations = {
  setUser(state, user) {
    state.user = user;
  }
};
