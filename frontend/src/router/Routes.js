import { createRouter, createWebHashHistory } from "vue-router";
import { validateRoute } from "../util/queries/user";

const routes = [
  {
    path: "/",
    name: "dashboard",
    component: () => import("../views/Dashboard.vue"),
  },
  {
    path: "/auth",
    name: "auth",
    component: () => import("../views/Auth.vue"),
  },
];

const router = createRouter({
  routes,
  history: createWebHashHistory(),
});

router.beforeEach(async (to) => {
  const token = localStorage.getItem("token");
  if (token == null && to.name != "auth") {
    return "auth";
  } else {
    const serverResponse = await validateRoute(token);
    if (serverResponse != 200 && to.name != "auth") {
      return "auth";
    }
  }
});

export default router;
