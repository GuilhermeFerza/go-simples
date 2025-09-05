"use client";
import { useEffect, useState } from "react";

export default function Home() {
  interface Person {
    id: number;
    nome: string;
    email: string;
  }

  const [data, setData] = useState<Person[]>([]);
  const [nome, setNome] = useState("");
  const [email, setEmail] = useState("");

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

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    const newPerson = { nome, email };

    fetch("http://localhost:8080/api/data", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newPerson),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network responsa was not ok");
        }
        return response.json();
      })
      .then((createdPerson) => {
        console.log("pessoa criada com sucesso: ", createdPerson);
        setData((prev) => [...prev, createdPerson]);
        setNome("");
        setEmail("");
      })
      .catch((error) => {
        console.error("erro ao enviar dados para o servidor: ", error);
      });
  }

  function handleDelete(id: number) {
    fetch(`http://localhost:8080/api/data/${id}`, {
      method: "DELETE",
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((deletedPerson) => {
        console.log(`pessoa com ID ${id} com sucesso:`, deletedPerson);
        setData((prev) => prev.filter((person) => person.id !== id));
      })
      .catch((error) => {
        console.error("erro ao deletar pessoa", error);
      });
  }

  return (
    <div className="flex flex-col justify-center items-center">
      <h1 className="mt-10">CRUD Simples</h1>
      <form
        className="flex flex-col mt-10 gap-4 border p-5"
        onSubmit={handleSubmit}
      >
        <h1 className="font-bold">Forms</h1>
        <input
          name="nome"
          id="nome"
          type="text"
          placeholder="name"
          value={nome}
          onChange={(e) => setNome(e.target.value)}
          className="border px-2 rounded"
        />
        <input
          name="email"
          id="email"
          type="email"
          placeholder="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className="border px-2 rounded"
        />
        <button
          className="border rounded-2xl p-1 cursor-pointer active:bg-gray-100 active:text-black"
          type="submit"
        >
          sei la teste
        </button>
      </form>

      <div>
        <h2 className="mt-5 font-bold">Dados do Backend:</h2>
        {data && data.length > 0 ? (
          <ul>
            {data.map((person) => (
              <li key={person.id}>
                <strong>ID:</strong> {person.id}, <strong>Nome:</strong>{" "}
                {person.nome}, <strong>Email:</strong> {person.email}
                <button
                  onClick={() => handleDelete(person.id)}
                  className="border rounded-2xl p-1 px-4 mt-3 cursor-pointer active:bg-red-100 active:text-black"
                >
                  excluir
                </button>
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
