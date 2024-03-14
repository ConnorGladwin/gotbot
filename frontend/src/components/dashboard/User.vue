<template>
  <div class="w-full">
    <div class="w-full flex justify-center mt-4">
      <span
        class="w-1/2 flex justify-center font-secondary text-3xl border-b border-black pb-2"
        >Edit Item</span
      >
    </div>
  </div>
  <div class="h-[500px] flex flex-col items-center gap-4 mt-8">
    <div class="w-full flex justify-center">
      <div class="w-1/2 flex flex-col gap-4">
        <div class="flex flex-col items-center mb-2">
          <label class="font-secondary text-xl">Username</label>
          <input
            v-model="username"
            placeholder="Username"
            class="w-[90%] font-primary text-xl border border-gray-400 p-2 focus:outline-none focus:border-black"
            :class="usernameError ? 'border-red-500' : ''"
          />
        </div>
        <div class="flex flex-col items-center mb-2">
          <label class="font-secondary text-xl">First Name</label>
          <input
            v-model="firstName"
            placeholder="First Name"
            class="w-[90%] font-primary text-xl border border-gray-400 p-2 focus:outline-none focus:border-black"
            :class="firstNameError ? 'border-red-500' : ''"
          />
        </div>
        <div class="flex flex-col items-center mb-2">
          <label class="font-secondary text-xl">Last Name</label>
          <input
            v-model="lastName"
            placeholder="Last Name"
            class="w-[90%] font-primary text-xl border border-gray-400 p-2 focus:outline-none focus:border-black"
            :class="lastNameError ? 'border-red-500' : ''"
          />
        </div>
        <div
          v-if="requestSuccess == true"
          class="w-full text-center font-primary text-green-500"
        >
          User successfully updated
        </div>
        <div @click="submit()" class="flex justify-center">
          <Button text="Submit" :loading="loading" class="w-[80%]" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import Button from "./Button.vue";
import { getUser, updateUser } from "../../util/queries/user.js";

const loading = ref(false);
const id = ref("");
const firstName = ref("");
const lastName = ref("");
const username = ref("");
const firstNameError = ref(false);
const lastNameError = ref(false);
const usernameError = ref(false);
const requestSuccess = ref(false);

async function submit() {
  firstName.value == ""
    ? (firstNameError.value = true)
    : (firstNameError.value = false);
  lastName.value == ""
    ? (lastNameError.value = true)
    : (lastNameError.value = false);
  username.value == ""
    ? (usernameError.value = true)
    : (usernameError.value = false);

  const errorArray = [
    firstNameError.value,
    lastNameError.value,
    usernameError.value,
  ];

  if (!errorArray.includes(true)) {
    const user = {
      firstName: firstName.value,
      lastName: lastName.value,
      username: username.value,
    };

    await updateUser(user).then(async (res) => {
      if (res.status == true) {
        await getUser().then((res) => {
          id.value = res.Id;
          firstName.value = res.FirstName;
          lastName.value = res.LastName;
          username.value = res.Email;
          requestSuccess.value = true;
          setTimeout(() => {
            requestSuccess.value = false;
          }, "1000");
        });
      }
    });
  }
}

onMounted(async () => {
  await getUser().then((res) => {
    id.value = res.Id;
    firstName.value = res.FirstName;
    lastName.value = res.LastName;
    username.value = res.Email;
  });
});
</script>
