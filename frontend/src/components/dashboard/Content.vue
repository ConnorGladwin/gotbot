<template>
  <div class="w-[1000px]">
    <div v-if="menuStore.getOption == 'dashboard'" class="w-full">
      <Inventory />
    </div>
    <div v-if="menuStore.getOption == 'add item'">
      <AddItem />
    </div>
    <div v-if="menuStore.getOption == 'edit item'" class="w-full">
      <EditItem />
    </div>
    <div v-if="menuStore.getOption == 'user'">
      <User />
    </div>
    <div v-if="menuStore.getOption == 'logout'">
      <Logout />
    </div>
  </div>
</template>

<script setup>
import { useMenuStore } from "../../store/menuStore.js";
import Inventory from "./Inventory.vue";
import AddItem from "./AddItem.vue";
import EditItem from "./EditItem.vue";
import User from "./User.vue";
import Logout from "./Logout.vue";

const menuStore = useMenuStore();
</script>
