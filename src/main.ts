import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";

import { createWebHistory, createRouter } from "vue-router";

import Habits from "./pages/Habits.vue";
import Add from "./pages/Add.vue";

const routes = [
  { path: "/", component: Habits },
  { path: "/add", component: Add },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

createApp(App).use(router).mount("#app");
