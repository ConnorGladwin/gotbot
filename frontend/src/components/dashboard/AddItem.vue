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
          <span v-if="nameError == true" class="font-primary text-red-500 mt-2"
            >Please provide a name for the item</span
          >
        </div>
        <div class="w-1/2 flex flex-col items-center">
          <span class="font-secondary text-xl">Description</span>
          <input
            v-model="desc"
            placeholder="Item description"
            class="w-[90%] font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
            :class="descError ? 'border-red-500' : ''"
          />
          <span v-if="descError == true" class="font-primary text-red-500 mt-2"
            >Please provide a description for the item</span
          >
        </div>
      </div>
      <div class="w-[80%] flex">
        <div class="w-1/2 flex flex-col items-center">
          <span class="font-secondary text-xl">Price</span>
          <input
            v-model="price"
            placeholder="Item name"
            type="number"
            class="w-[90%] font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
            :class="priceError ? 'border-red-500' : ''"
          />
          <span v-if="priceError == true" class="font-primary text-red-500 mt-2"
            >Please provide a price for the item</span
          >
        </div>
        <div class="w-1/2 flex flex-col items-center">
          <span class="font-secondary text-xl">Stock</span>
          <input
            v-model="stock"
            placeholder="Item name"
            type="number"
            class="w-[90%] font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
            :class="stockError ? 'border-red-500' : ''"
          />
          <span v-if="stockError == true" class="font-primary text-red-500 mt-2"
            >Please provide a stock amout for the item</span
          >
        </div>
      </div>
      <span v-if="requestSuccess == true" class="font-primary text-green-500"
        >Item added successfully</span
      >
      <div @click="submit()" class="mt-4">
        <Button :loading="loading" text="add item" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import Button from "./Button.vue";
import { addItem } from "../../util/queries/food.js";

const name = ref("");
const desc = ref("");
const price = ref(0);
const stock = ref(0);
const loading = ref(false);
const nameError = ref(false);
const descError = ref(false);
const priceError = ref(false);
const stockError = ref(false);

const requestSuccess = ref(false);

async function submit() {
  loading.value = true;
  name.value == "" ? (nameError.value = true) : (nameError.value = false);
  desc.value == "" ? (descError.value = true) : (descError.value = false);
  price.value < 1 ? (priceError.value = true) : (priceError.value = false);
  stock.value < 1 ? (stockError.value = true) : (stockError.value = false);

  const errorArray = [
    nameError.value,
    descError.value,
    priceError.value,
    stockError.value,
  ];

  const item = {
    name: name.value,
    desc: desc.value,
    price: price.value,
    stock: stock.value,
  };

  if (!errorArray.includes(true)) {
    await addItem(item).then((res) => {
      if (res.status == true) {
        requestSuccess.value = true;
        loading.value = false;
        name.value = "";
        desc.value = "";
        price.value = 0;
        stock.value = 0;
        setTimeout(() => {
          requestSuccess.value = false;
        }, "1000");
      }
    });
  } else {
    loading.value = false;
  }
}
</script>
