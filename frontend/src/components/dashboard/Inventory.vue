<template>
  <div class="w-full">
    <div class="w-full flex justify-center mt-4">
      <span
        class="w-1/2 flex justify-center font-secondary text-3xl border-b border-black pb-2"
        >Inventory</span
      >
    </div>
    <div class="flex flex-col items-center gap-4 mt-8">
      <div v-for="item in list" :key="item" class="w-[80%]">
        <ItemCard :item="item" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { inventory } from "../../util/queries/food.js";
import ItemCard from "./ItemCard.vue";

const list = ref();

onMounted(async () => {
  list.value = await inventory();

  console.log(list.value);
});
</script>
