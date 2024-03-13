<template>
  <div class="w-full">
    <div class="w-full flex justify-center mt-4">
      <span
        class="w-1/2 flex justify-center font-secondary text-3xl border-b border-black pb-2"
        >Add Item</span
      >
    </div>
    <div class="flex flex-col items-center gap-4 mt-8">
      <div class="w-1/2 flex">
        <div class="w-1/2 flex flex-col items-center">
          <span class="font-secondary text-xl">Name</span>
          <input
            v-model="name"
            placeholder="Item name"
            class="w-3/5 font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
            :class="nameError ? 'border-red-500' : ''"
          />
        </div>
        <div class="w-1/2 flex flex-col items-center">
          <span class="font-secondary text-xl">Description</span>
        </div>
      </div>
      <div class="w-1/2 flex">
        <div class="w-1/2 flex flex-col items-center">
          <span class="font-secondary text-xl">Price</span>
        </div>
        <div class="w-1/2 flex flex-col items-center">
          <span class="font-secondary text-xl">Stock</span>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup></script>
