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
  const [editingPerson, setEditingPerson] = useState<Person | null>(null);

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
    if (editingPerson) {
      handleUpdate(editingPerson.id);
    } else {
      handleCreate();
    }
  }

  function handleCreate() {
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
          throw new Error("Network response was not ok");
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
      .then(() => {
        console.log(`pessoa com ID ${id} deletada com sucesso.`);
        setData((prev) => prev.filter((person) => person.id !== id));
      })
      .catch((error) => {
        console.error("erro ao deletar pessoa", error);
      });
  }

  function handleEditClick(person: Person) {
    setEditingPerson(person);
    setNome(person.nome);
    setEmail(person.email);
  }

  function handleUpdate(id: number) {
    const updatedData = {
      id,
      nome,
      email,
    };
    fetch(`http://localhost:8080/api/data/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(updatedData),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to update person");
        }
        return response.json();
      })
      .then((updatedPerson) => {
        console.log("Pessoa atualizada:", updatedPerson);
        setData((prev) =>
          prev.map((person) => (person.id === id ? updatedPerson : person))
        );
        // Reset form state
        setEditingPerson(null);
        setNome("");
        setEmail("");
      })
      .catch((error) => {
        console.error("Erro ao atualizar a pessoa:", error);
      });
  }

  return (
    <div className="flex flex-col justify-center items-center">
      <h1 className="mt-10">CRUD Simples</h1>
      <form
        className="flex flex-col mt-10 gap-4 border p-5"
        onSubmit={handleSubmit}
      >
        <h1 className="font-bold">
          {editingPerson ? "Editar Pessoa" : "Adicionar Nova Pessoa"}
        </h1>
        <input
          name="nome"
          id="nome"
          type="text"
          placeholder="nome"
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
          {editingPerson ? "Salvar Alterações" : "Adicionar"}
        </button>
      </form>

      <div>
        <h2 className="mt-5 font-bold">Dados do Backend:</h2>
        {data && data.length > 0 ? (
          <ul>
            {data.map((person) => (
              <li key={person.id} className="flex items-center gap-2">
                <strong>ID:</strong> {person.id}, <strong>Nome:</strong>{" "}
                {person.nome}, <strong>Email:</strong> {person.email}
                <button
                  onClick={() => handleEditClick(person)}
                  className="border rounded-2xl p-1 px-2 text-xs cursor-pointer bg-blue-500 text-white hover:bg-blue-600"
                >
                  Editar
                </button>
                <button
                  onClick={() => handleDelete(person.id)}
                  className="border rounded-2xl p-1 px-2 text-xs cursor-pointer bg-red-500 text-white hover:bg-red-600"
                >
                  Excluir
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
