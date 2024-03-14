import { defineStore } from "pinia";

export const useEditStore = defineStore("edit", {
  state: () => {
    return {
      pageState: "list",
      itemId: "",
    };
  },
  getters: {
    getPageState: (state) => {
      return state.pageState;
    },
    getItemId: (state) => {
      return state.itemId;
    },
  },
  actions: {
    setPageState(state) {
      this.pageState = state;
    },
    setItemId(id) {
      this.itemId = id;
    },
    stateReset() {
      this.pageState = "list";
      this.itemId = "";
    },
  },
});
