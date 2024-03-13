<template>
  <div>
    <Header class="pt-[10%] pb-[5%]" />
    <div class="h-full flex flex-col justify-center items-center">
      <div class="h-[500px] w-3/4 flex gap-6">
        <div class="w-1/4">
          <Menu />
        </div>
        <div class="w-3/4 border">
          <Content />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import Header from "../components/auth/Header.vue";
import Menu from "../components/dashboard/Menu.vue";
import Content from "../components/dashboard/Content.vue";
</script>
