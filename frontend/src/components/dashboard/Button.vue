<template>
  <div
    class="w-full flex justify-center items-center py-2 border border-black bg-black text-white font-secondary text-xl cursor-pointer hover:bg-white hover:text-black transition duration-200"
  >
    <span v-if="loading == false" class="px-2 capitalize">{{ text }}</span>
    <div v-else class="flex items-center justify-center ml-10">
      <div
        class="inline-block h-8 w-8 animate-spin rounded-full border-4 border-solid border-current border-r-transparent align-[-0.125em] motion-reduce:animate-[spin_1.5s_linear_infinite]"
        role="status"
      >
        <span
          class="!absolute !-m-px !h-px !w-px !overflow-hidden !whitespace-nowrap !border-0 !p-0 ![clip:rect(0,0,0,0)]"
          >Loading...</span
        >
      </div>
      <div
        class="inline-block h-8 w-8 animate-[spinner-grow_0.75s_linear_infinite] rounded-full bg-current align-[-0.125em] opacity-0 motion-reduce:animate-[spinner-grow_1.5s_linear_infinite]"
        role="status"
      >
        <span
          class="!absolute !-m-px !h-px !w-px !overflow-hidden !whitespace-nowrap !border-0 !p-0 ![clip:rect(0,0,0,0)]"
          >Loading...</span
        >
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps(["loading", "text"]);
</script>
