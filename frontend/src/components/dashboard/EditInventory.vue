<template>
  <div class="flex flex-col items-center">
    <span class="capitalize font-primary">{{ type }} inventory</span>
    <input v-model="num" type="number" class="w-1/3 text-center font-primary" />
    <span @click="submit()" class="font-primary cursor-pointer hover:underline"
      >Submit</span
    >
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { updateInventory } from "../../util/queries/food.js";

const router = useRouter();

const props = defineProps(["type", "value", "id"]);
const num = ref(0);

async function submit() {
  let updateValue;
  props.type == "add"
    ? (updateValue = props.value + num.value)
    : (updateValue = props.value - num.value);

  await updateInventory(props.id, updateValue).then((res) => {
    if (res.status == "success") {
      router.go();
    }
  });
}
</script>
