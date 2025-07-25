<template>
  <div class="p-6 max-w-xl mx-auto">
    <router-link to="/Add" class="text-blue-500 underline"
      >Go to Add Habit</router-link
    >

    <h1 class="text-2xl font-bold mb-4">🧠 Habit Tracker</h1>

    <ul class="space-y-2">
      <li
        v-for="habit in habits"
        :key="habit.id"
        class="p-4 rounded bg-gray-100 shadow"
      >
        {{ habit.name }}
      </li>
    </ul>

    <Button>Click me </Button>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { Button } from "@/components/ui/button";

const habits = ref([]);

onMounted(async () => {
  const res = await fetch("http://localhost:8080/habits");
  const data = await res.json();
  habits.value = data;
});
</script>
