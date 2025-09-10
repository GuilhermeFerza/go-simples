"use client";
import { useState } from "react";

interface Task {
  id: string; // Assuming the task has an ID
  task: string;
}

export default function Home() {
  const [inputTask, setInputTask] = useState("");
  const [tasksList, setTasksList] = useState<Task[]>([]);

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    handleCreate();
  }

  function handleCreate() {
    const newTask = { task: inputTask };
    fetch("http://localhost:8080/api/data", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newTask),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((createdTask) => {
        console.log("task criada com sucesso: ", createdTask);
        setTasksList((prev) => [...prev, createdTask as Task]); // Cast createdTask to Task
        setInputTask(""); // Limpa o input
      })
      .catch((error) => {
        console.error("erro ao enviar dados para o servidor: ", error);
      });
  }

  return (
    <div className="flex justify-center items-center h-screen">
      <form
        onSubmit={handleSubmit}
        className="border p-7 rounded bg-[#2e2e2e] text-[#dfdfdf] text-2xl flex flex-col items-center gap-10"
      >
        <h1>Tasks</h1>
        <input
          className="bg-[#dfdfdf] text-[#2e2e2e]"
          type="text"
          value={inputTask}
          onChange={(e) => setInputTask(e.target.value)}
        />
        <button type="submit">adicionar</button>
      </form>
      <div>
        <ul>
          {tasksList.map((task) => (
            <li key={task.id}>{task.task}</li>
          ))}
        </ul>
      </div>
    </div>
  );
}
