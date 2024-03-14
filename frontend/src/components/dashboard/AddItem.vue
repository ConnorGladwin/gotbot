<template>
  <div class="w-full">
    <div class="w-full flex justify-center mt-4">
      <span
        class="w-1/2 flex justify-center font-secondary text-3xl border-b border-black pb-2"
        >Add Item</span
      >
    </div>
    <div class="flex flex-col items-center gap-4 mt-8">
      <div class="w-[80%] flex">
        <div class="w-1/2 flex flex-col items-center">
          <span class="font-secondary text-xl">Name</span>
          <input
            v-model="name"
            placeholder="Item name"
            class="w-[90%] font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
            :class="nameError ? 'border-red-500' : ''"
          />
        </div>
        <div class="w-1/2 flex flex-col items-center">
          <span class="font-secondary text-xl">Description</span>
          <input
            v-model="desc"
            placeholder="Item name"
            class="w-[90%] font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
            :class="descError ? 'border-red-500' : ''"
          />
        </div>
      </div>
      <div class="w-[80%] flex">
        <div class="w-1/2 flex flex-col items-center">
          <span class="font-secondary text-xl">Price</span>
          <input
            v-model="price"
            placeholder="Item name"
            class="w-[90%] font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
            :class="priceError ? 'border-red-500' : ''"
          />
        </div>
        <div class="w-1/2 flex flex-col items-center">
          <span class="font-secondary text-xl">Stock</span>
          <input
            v-model="stock"
            placeholder="Item name"
            class="w-[90%] font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
            :class="stockError ? 'border-red-500' : ''"
          />
        </div>
      </div>
      <button @click="submit()">submit</button>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const name = ref("");
const desc = ref("");
const price = ref(0);
const stock = ref(0);
const nameError = ref(false);
const descError = ref(false);
const priceError = ref(false);
const stockError = ref(false);

async function submit() {
  name.value == "" ? (nameError.value = true) : (nameError.value = false);
  desc.value == "" ? (descError.value = true) : (descError.value = false);
  price.value < 1 ? (priceError.value = true) : (priceError.value = false);
  stock.value < 1 ? (stockError.value = true) : (stockError.value = false);
}
</script>
