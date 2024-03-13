<template>
  <div class="h-full flex items-center">
    <div class="w-1/2 flex justify-center items-center">
      <Header />
    </div>
    <div class="h-[500px] border border-black"></div>
    <div class="w-2/3 flex justify-center items-center">
      <div
        v-if="authState == 'signup'"
        class="w-full flex flex-col items-center"
      >
        <SignUp />
        <div>
          <span>Already have an account? </span>
          <span @click="authState = 'login'" class="underline cursor-pointer"
            >Login</span
          >
        </div>
      </div>
      <div
        v-if="authState == 'login'"
        class="w-full flex flex-col items-center"
      >
        <Login />
        <div>
          <span>Don't have an account yet? </span>
          <span @click="authState = 'signup'" class="underline cursor-pointer"
            >Sign up</span
          >
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import Header from "../components/auth/Header.vue";
import SignUp from "../components/auth/SignUp.vue";
import Login from "../components/auth/Login.vue";

const authState = ref("login");
</script>
