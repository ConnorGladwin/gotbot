<template>
  <div>
    <Header class="pt-[5%] pb-[5%]" />
    <div class="h-full flex flex-col justify-center items-center">
      <div class="h-[500px] w-[90%] flex gap-3 pr-[100px]">
        <div class="w-[30%]">
          <Menu />
        </div>
        <div class="w-[80%]">
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
