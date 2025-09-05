"use client";
import { useEffect, useState } from "react";

export default function Home() {
  interface Person {
    id: number;
    nome: string;
    email: string;
  }

  const [data, setData] = useState<Person[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/api/data")
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((json) => setData(json))
      .catch((error) => console.error("Erro ao buscar dados: ", error));
  }, []);

  function handleSubmit() {}

  return (
    <div className="flex flex-col justify-center items-center">
      <h1 className="mt-10">CRUD Simples</h1>
      <form className="flex flex-col mt-10 gap-4 border p-5">
        <h1 className="font-bold">Forms</h1>
        <input
          name="nome"
          id="nome"
          type="text"
          placeholder="name"
          className="border px-2 rounded"
        />
        <input
          name="email"
          id="email"
          type="email"
          placeholder="email"
          className="border px-2 rounded"
        />
        <button
          className="border rounded-2xl p-1 cursor-pointer active:bg-gray-100 active:text-black"
          onClick={handleSubmit}
        >
          sei la teste
        </button>
      </form>
      <div>
        {/* Render the data as a list */}
        <h2 className="mt-5 font-bold">Dados do Backend:</h2>
        {data && data.length > 0 ? (
          <ul>
            {data.map((person) => (
              <li key={person.id}>
                <strong>ID:</strong> {person.id}, <strong>Nome:</strong>{" "}
                {person.nome}, <strong>Email:</strong> {person.email}
              </li>
            ))}
          </ul>
        ) : (
          <p>Carregando dados...</p>
        )}
      </div>
    </div>
  );
}
