<template></template>

<script setup>
import { onBeforeMount } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

onBeforeMount(() => {
  localStorage.removeItem("id");
  localStorage.removeItem("token");
  router.go();
});
</script>
