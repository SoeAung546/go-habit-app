<template>
  <div class="p-6 max-w-xl mx-auto">
    <h1 class="text-2xl font-bold mb-4">🧠 Habit Tracker</h1>

    <ul class="space-y-2 mb-6">
      <li
        v-for="habit in habits"
        :key="habit.id"
        class="p-4 rounded bg-gray-100 shadow"
      >
        {{ habit.name }}
      </li>
    </ul>

    <form @submit.prevent="addHabit" class="flex space-x-2">
      <input
        v-model="newHabit"
        type="text"
        placeholder="New habit"
        required
        class="border rounded px-3 py-2 flex-grow"
      />
      <button type="submit" class="bg-blue-600 text-white px-4 rounded">
        Add
      </button>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";

const habits = ref([]);
const newHabit = ref("");

async function fetchHabits() {
  const res = await fetch("http://localhost:8080/habits");
  habits.value = await res.json();
}

onMounted(() => {
  fetchHabits();
});

async function addHabit() {
  if (!newHabit.value.trim()) return;

  const res = await fetch("http://localhost:8080/habits", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name: newHabit.value.trim() }),
  });

  if (res.ok) {
    const habit = await res.json();
    habits.value.push(habit);
    newHabit.value = "";
  } else {
    alert("Failed to add habit");
  }
}
</script>
