<template>
  <div class="w-full flex justify-center mb-8">
    <div class="w-2/3 flex flex-col items-center">
      <span class="text-4xl font-secondary mb-6">Sign Up</span>
      <div class="w-full flex flex-col items-center">
        <div class="w-full flex">
          <div class="w-1/2 flex flex-col items-center mb-2">
            <label class="font-primary text-xl">First Name</label>
            <input
              v-model="firstName"
              type="text"
              placeholder="John"
              class="w-4/5 font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
              :class="nameError ? 'border-red-500' : ''"
            />
          </div>
          <div class="w-1/2 flex flex-col items-center mb-2">
            <label class="font-primary text-xl">Last Name</label>
            <input
              v-model="lastName"
              type="text"
              placeholder="Doe"
              class="w-4/5 font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
              :class="nameError ? 'border-red-500' : ''"
            />
          </div>
        </div>
        <div class="w-full flex">
          <div class="w-1/2 flex flex-col items-center mb-2">
            <label class="font-primary text-xl">Email</label>
            <input
              v-model="email"
              type="text"
              placeholder="name@company.com"
              class="w-4/5 font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
              :class="emailError ? 'border-red-500' : ''"
            />
          </div>
          <div class="w-1/2 flex flex-col items-center mb-2">
            <label class="font-primary text-xl">Confirm Email</label>
            <input
              v-model="emailConfirm"
              type="text"
              placeholder="name@company.com"
              class="w-4/5 font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
              :class="emailError ? 'border-red-500' : ''"
            />
          </div>
        </div>
        <div class="w-full flex">
          <div class="w-1/2 flex flex-col items-center mb-2">
            <label class="font-primary text-xl">Password</label>
            <input
              v-model="password"
              :type="passwordType"
              placeholder="Password"
              class="w-4/5 font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
              :class="passwordError ? 'border-red-500' : ''"
            />
          </div>
          <div class="w-1/2 flex flex-col items-center">
            <label class="font-primary text-xl">Confirm Password</label>
            <input
              v-model="passwordConfirm"
              :type="confirmPasswordType"
              placeholder="Password"
              class="w-4/5 font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
              :class="passwordError ? 'border-red-500' : ''"
            />
          </div>
        </div>
        <div class="w-1/2 flex flex-col items-center">
          <label class="font-primary text-xl">Username</label>
          <input
            v-model="username"
            type="text"
            placeholder="password"
            class="w-3/4 font-primary text-xl border border-gray-400 p-1 focus:outline-none focus:border-black"
            :class="nameError ? 'border-red-500' : ''"
          />
        </div>
      </div>
      <div class="my-4 text-center">
        <div v-if="nameError == true" class="font-primary text-red-500">
          Please make sure you've added your first and last name, as well as a
          username.
        </div>
        <div v-if="emailError == true" class="font-primary text-red-500">
          Email and email confirmation don't match.
        </div>
        <div v-if="passwordError == true" class="font-primary text-red-500">
          Password and password confirmation don't match.
        </div>
        <div
          v-if="requestError == true"
          class="font-primary text-red-500 capitalize"
        >
          {{ requestErrorMessage }}
        </div>
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
import { signUp } from "../../util/queries/user.js";
import Button from "./Button.vue";

const router = useRouter();

const firstName = ref("");
const lastName = ref("");
const email = ref("");
const emailConfirm = ref("");
const password = ref("");
const passwordConfirm = ref("");
const username = ref("");
const loading = ref(false);

const passwordType = ref("password");
const confirmPasswordType = ref("password");

const error = ref(false);
const nameError = ref(false);
const emailError = ref(false);
const passwordError = ref(false);
const requestError = ref(false);
const requestErrorMessage = ref("");

function validateEmail() {
  if (email.value == "" || emailConfirm.value == "") {
    emailError.value = true;
    error.value = true;
    loading.value = false;
  }
  if (email.value != emailConfirm.value) {
    emailError.value = true;
    error.value = true;
    loading.value = false;
  }
}

function validatePassword() {
  if (password.value == "" || passwordConfirm.value == "") {
    passwordError.value = true;
    error.value = true;
    loading.value = false;
  }
  if (password.value != passwordConfirm.value) {
    passwordError.value = true;
    error.value = true;
    loading.value = false;
  }
}

function validateNames() {
  if (firstName.value == "" || lastName.value == "" || username.value == "") {
    nameError.value = true;
    error.value = true;
    loading.value = false;
  }
}

async function submit() {
  emailError.value = false;
  passwordError.value = false;
  nameError.value = false;
  error.value = false;
  loading.value = true;
  validateNames();
  validateEmail();
  validatePassword();

  if (error.value == false) {
    const userObject = {
      firstName: firstName.value.trim(),
      lastName: lastName.value.trim(),
      email: email.value.trim(),
      password: password.value.trim(),
      username: username.value.trim(),
    };

    const res = await signUp(userObject);

    if (res.id == "failed") {
      requestError.value = true;
      requestErrorMessage.value = res.message;
      loading.value = false;
    } else if (res.message == "success") {
      console.log(res);
      localStorage.setItem("token", res.token);
      localStorage.setItem("id", res.id);
      router.push("/");
    }
  }
}
</script>
