<template>
  <div class="p-6 max-w-xl mx-auto">
    <nav class="flex space-x-4 p-4 bg-gray-100 rounded shadow">
      <router-link to="/" class="font-bold text-lg hover:text-green-600"
        >All Habit</router-link
      >
      <router-link to="/add" class="font-bold text-lg hover:text-green-600"
        >Add Habit</router-link
      >
    </nav>
    <h1 class="text-2xl font-bold my-4">🧠 Habit Tracker</h1>

    <div v-if="habits.length === 0" class="text-gray-500 mb-4">
      No habits found. Add a new habit to get started!
    </div>
    <div v-else class="mb-6">
      <ul class="space-y-2">
        <li
          v-for="habit in habits"
          :key="habit.id"
          class="p-3 rounded bg-gray-100 shadow flex justify-between items-center"
        >
          <p>{{ habit.name }}</p>
          <div class="flex space-x-2">
            <Button
              @click="toggleDone(habit)"
              class="text-xl cursor-pointer"
              variant="ghost"
              size="icon"
            >
              {{ habit.done ? "✅" : "🔲" }}
            </Button>
            <Button
              @click="handleDelete(habit.id)"
              type="button"
              variant="default"
              class="px-4 cursor-pointer"
            >
              Delete
            </Button>
          </div>
        </li>
      </ul>
    </div>

    <form @submit.prevent="addHabit" class="flex space-x-2 items-center">
      <input
        v-model="newHabit"
        type="text"
        placeholder="New habit"
        required
        class="border rounded p-2 flex-grow"
      />
      <Button type="submit"> Add </Button>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { Button } from "@/components/ui/button";

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

async function handleDelete(id) {
  await fetch(`http://localhost:8080/habits/${id}`, {
    method: "DELETE",
  });
  habits.value = habits.value.filter((habit) => habit.id !== id);
}

async function toggleDone(habit) {
  const res = await fetch(`http://localhost:8080/habits/${habit.id}/done`, {
    method: "PATCH",
  });
  if (res.ok) {
    const updated = await res.json();
    habit.done = updated.done; // update the local state
  }
}
</script>
