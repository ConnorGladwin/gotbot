<template>
  <div class="w-full h-full flex flex-col items-center">
    <div
      class="h-1/5 w-2/3 flex justify-center items-center border-b border-black"
    >
      <span class="font-secondary text-3xl">Menu</span>
    </div>
    <div class="h-4/5 w-full flex flex-col items-center justify-evenly">
      <div
        v-for="item in menuList"
        :key="item"
        class="w-full flex flex-col items-center cursor-pointer group"
      >
        <div class="flex items-center mb-3">
          <div class="w-5 h-5">
            <img :src="item.icon" alt="" />
          </div>
          <div
            class="font-secondary text-xl capitalize py-1 px-2 ml-2 group-hover:bg-black group-hover:text-white transition duration-100"
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

const menuList = [
  {
    name: "dashboard",
    icon: dashboard,
  },
  {
    name: "add item",
    icon: "",
  },
  {
    name: "edit item",
    icon: "",
  },
  {
    name: "user",
    icon: "",
  },
  {
    name: "logout",
    icon: "",
  },
];
</script>
