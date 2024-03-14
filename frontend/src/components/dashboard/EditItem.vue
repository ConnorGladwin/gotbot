<template>
  <div class="w-full">
    <div class="w-full flex justify-center mt-4">
      <span
        class="w-1/2 flex justify-center font-secondary text-3xl border-b border-black pb-2"
        >Edit Item</span
      >
    </div>
    <div
      class="h-[500px] flex flex-col items-center gap-4 mt-8 overflow-y-scroll"
    >
      <div v-for="item in list" :key="item" class="w-[80%]">
        <EditItemCard :item="item" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { inventory } from "../../util/queries/food.js";
import EditItemCard from "./EditItemCard.vue";

const list = ref();

onMounted(async () => {
  list.value = await inventory();
});
</script>
