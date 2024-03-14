<template>
  <div class="w-full">
    <div class="w-full flex justify-center mt-4">
      <span
        class="w-1/2 flex justify-center font-secondary text-3xl border-b border-black pb-2"
        >Edit Item</span
      >
    </div>
    <div
      v-if="editStore.getPageState == 'list'"
      class="h-[500px] flex flex-col items-center gap-4 mt-8 overflow-y-scroll"
    >
      <div v-for="item in list" :key="item" class="w-[80%] pb-2">
        <EditItemCard :item="item" />
      </div>
    </div>
    <div v-else class="h-[500px] flex flex-col items-center gap-4 mt-8">
      <div class="w-full flex justify-center">
        <div class="w-1/2 flex flex-col gap-4">
          <div class="flex flex-col items-center mb-2">
            <label class="font-secondary text-xl">Name</label>
            <input
              v-model="name"
              placeholder="Item name"
              class="w-[90%] font-primary text-xl border border-gray-400 p-2 focus:outline-none focus:border-black"
              :class="nameError ? 'border-red-500' : ''"
            />
          </div>
          <div class="flex flex-col items-center mb-2">
            <label class="font-secondary text-xl">Description</label>
            <input
              v-model="desc"
              placeholder="Item name"
              class="w-[90%] font-primary text-xl border border-gray-400 p-2 focus:outline-none focus:border-black"
              :class="descError ? 'border-red-500' : ''"
            />
          </div>
          <div class="flex flex-col items-center mb-2">
            <label class="font-secondary text-xl">Price</label>
            <input
              v-model="price"
              placeholder="Item price"
              type="number"
              class="w-[90%] font-primary text-xl border border-gray-400 p-2 focus:outline-none focus:border-black"
              :class="priceError ? 'border-red-500' : ''"
            />
          </div>
          <div @click="submit()" class="flex justify-center">
            <Button text="Submit" :loading="loading" class="w-[80%]" />
          </div>
          <div @click="editStore.stateReset()" class="flex justify-center">
            <Button text="Cancel" :loading="loading" class="w-[80%]" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref, watch } from "vue";
import { inventory } from "../../util/queries/food.js";
import { useEditStore } from "../../store/editStore";
import { getById, updateItem } from "../../util/queries/food.js";
import EditItemCard from "./EditItemCard.vue";
import Button from "./Button.vue";

const editStore = useEditStore();

const loading = ref(false);
const list = ref();
const id = ref("");
const name = ref("");
const desc = ref("");
const price = ref("");
const nameError = ref(false);
const descError = ref(false);
const priceError = ref(false);

watch(editStore, async () => {
  if (editStore.pageState == "item") {
    await getById(editStore.getItemId).then((res) => {
      id.value = res.Id;
      name.value = res.Name;
      desc.value = res.Desc;
      price.value = res.Price;
    });
  } else if (editStore.pageState == "") {
    editStore.stateReset();
    list.value = await inventory();
  } else {
    list.value = await inventory();
  }
});

async function submit() {
  name.value == "" ? (nameError.value = true) : (nameError.value = false);
  desc.value == "" ? (descError.value = true) : (descError.value = false);
  price.value < 1 ? (priceError.value = true) : (priceError.value = false);

  const errorArray = [nameError.value, descError.value, priceError.value];

  if (!errorArray.includes(true)) {
    const item = {
      id: id.value,
      name: name.value,
      desc: desc.value,
      price: price.value,
    };

    await updateItem(item).then((res) => {
      if (res.status == true) {
        editStore.stateReset();
      }
    });
  }
}

onMounted(async () => {
  list.value = await inventory();
});
</script>
