<template>
  <div class="w-full flex justify-center mb-8">
    <div class="w-2/3 flex flex-col items-center">
      <span class="text-4xl font-secondary mb-6">Login</span>
      <div class="w-2/3 flex flex-col items-center">
        <div class="w-full flex flex-col items-center mb-2">
          <label class="font-primary text-xl">Username</label>
          <input
            v-model="id"
            placeholder="Username/Email"
            class="w-3/5 font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
            :class="idError ? 'border-red-500' : ''"
          />
        </div>
        <div class="w-full flex flex-col items-center">
          <label class="font-primary text-xl">Password</label>
          <input
            v-model="password"
            type="password"
            placeholder="Password"
            class="w-3/5 font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
            :class="passwordError ? 'border-red-500' : ''"
          />
        </div>
      </div>
      <div class="my-4">
        <span v-if="idError == true" class="font-primary text-red-500">
          Please provide a username.
        </span>
        <span v-if="passwordError == true" class="font-primary text-red-500">
          Please provide your password.
        </span>
        <span
          v-if="requestError == true"
          class="font-primary text-red-500 capitalize"
          >{{ requestErrorMessage }}</span
        >
      </div>
      <div @click="submit()" class="w-1/5">
        <Button :loading="loading" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import Button from "./Button.vue";
import { login } from "../../util/queries/user.js";

const router = useRouter();

const loading = ref(false);
const id = ref("");
const password = ref("");

const idError = ref(false);
const passwordError = ref(false);
const requestError = ref(false);
const requestErrorMessage = ref("");
const error = ref(false);

function validateId() {
  if (id.value == "") {
    idError.value = true;
    error.value = true;
    loading.value = false;
  }
}

function validatePassword() {
  if (password.value == "") {
    passwordError.value = true;
    error.value = true;
    loading.value = false;
  }
}

async function submit() {
  idError.value = passwordError.value = false;
  loading.value = true;
  validateId();
  validatePassword();

  if (error.value == false) {
    const userObject = {
      id: id.value.trim(),
      password: password.value.trim(),
    };
    const res = await login(userObject);
    if (res.status == "error") {
      requestError.value = true;
      requestErrorMessage.value = res.message;
      loading.value = false;
    } else if (res.status == "success") {
      localStorage.setItem("token", res.token);
      localStorage.setItem("id", res.id);
      router.push("/");
    }
  }
}
</script>
