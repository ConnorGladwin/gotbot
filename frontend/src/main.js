import { createApp } from "vue";
import "./style.css";
import { createPinia } from "pinia";
import router from "./router/Routes.js";
import App from "./App.vue";

const pinia = createPinia();

createApp(App).use(router).use(pinia).mount("#app");
