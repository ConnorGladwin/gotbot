<template>
  <div class="h-full flex border border-black py-2 px-3">
    <div class="w-1/3 flex flex-col">
      <span class="font-primary text-xl">Name</span>
      <span class="font-secondary uppercase text-2xl">{{ item.Name }}</span>
    </div>
    <div class="border-l border-black pr-4"></div>
    <div class="w-1/3 flex flex-col justify-center px-2">
      <span class="font-primary text-xl">Description</span>
      <span class="font-secondary uppercase text-xl pr-2">{{ item.Desc }}</span>
    </div>
    <div class="border-l border-black pr-4"></div>
    <div class="w-1/4 flex flex-col justify-center">
      <span class="font-primary text-xl">Price</span>
      <span class="font-secondary uppercase text-2xl">R{{ item.Price }}</span>
    </div>
    <div class="border-l border-black mx-6"></div>
    <div class="w-1/4 flex flex-col items-center justify-center font-secondary">
      <span
        @click="editItem()"
        class="flex justify-center text-xl py-1 px-1 my-1 cursor-pointer hover:bg-black hover:text-white"
      >
        Edit
      </span>
      <span
        @click="deleteId()"
        class="flex justify-center text-xl py-1 px-1 my-1 cursor-pointer hover:bg-black hover:text-white"
      >
        Delete
      </span>
    </div>
  </div>
</template>

<script setup>
import { useEditStore } from "../../store/editStore.js";
import { deleteItem } from "../../util/queries/food.js";

const editStore = useEditStore();
const props = defineProps(["item"]);

function editItem() {
  editStore.setItemId(props.item.Id);
  editStore.setPageState("item");
}

async function deleteId() {
  await deleteItem(props.item.Id).then((res) => {
    editStore.setPageState("");
  });
}
</script>
