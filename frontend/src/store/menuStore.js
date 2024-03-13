import { defineStore } from "pinia";

export const useMenuStore = defineStore("menu", {
  state: () => {
    return {
      option: "dashboard",
    };
  },
  getters: {
    getOption: (state) => state.option,
  },
  actions: {
    setOption(option) {
      this.option = option;
    },
  },
});
