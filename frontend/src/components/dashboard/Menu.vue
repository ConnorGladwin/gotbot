<template>
  <div class="w-[400px] h-full flex flex-col items-center">
    <div
      class="h-1/5 w-2/3 flex justify-center items-center border-b border-black"
    >
      <span class="font-secondary text-3xl">Menu</span>
    </div>
    <div
      :key="menuStore.option"
      class="h-4/5 w-full flex flex-col items-center justify-evenly"
    >
      <div
        v-for="item in menuList"
        :key="item"
        class="w-full flex flex-col items-center cursor-pointer group"
      >
        <div
          @click="
            menuStore.setOption(item.name);
            current = menuStore.getOption;
          "
          class="w-1/2 flex items-center justify-start mb-3"
        >
          <div class="w-8 h-8">
            <img :src="item.icon" alt="" />
          </div>
          <div
            class="font-secondary text-xl capitalize py-1 px-2 ml-2 group-hover:bg-black group-hover:text-white transition duration-100"
            :class="
              current == item.name
                ? 'bg-black text-white'
                : 'group-hover:bg-black group-hover:text-white'
            "
          >
            {{ item.name }}
          </div>
        </div>
        <div class="w-1/2 border-b border-black"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import dashboard from "../../assets/dashboard.svg";
import {
  DashboardIcon,
  AddIcon,
  EditIcon,
  UserIcon,
  LogoutIcon,
} from "../../assets/index.js";
import { useMenuStore } from "../../store/menuStore.js";

const menuStore = useMenuStore();

const menuList = [
  {
    name: "dashboard",
    icon: DashboardIcon,
  },
  {
    name: "add item",
    icon: AddIcon,
  },
  {
    name: "edit item",
    icon: EditIcon,
  },
  {
    name: "user",
    icon: UserIcon,
  },
  {
    name: "logout",
    icon: LogoutIcon,
  },
];

const current = ref(menuStore.option);
</script>
