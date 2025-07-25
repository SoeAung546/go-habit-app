import { createRouter, createWebHistory } from "vue-router";
import Home from "../pages/Home.vue";
import Add from "../pages/Add.vue";

const routes = [
  { path: "/", name: "Home", component: Home },
  { path: "/add", name: "Add", component: Add },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
