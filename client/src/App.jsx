import React, { useState } from "react";

function App() {

  const [name, setName] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    const response = await fetch(import.meta.env.VITE_API + '/users', {
      method: 'POST',
      body: JSON.stringify({name})
    })
    const data = await response.json()
    console.log(data)
  };

  return (
    <div>
      <h1>Hola mundo</h1>
      <button
        onClick={async () => {
          const res = await fetch(import.meta.env.VITE_API + "/users");
          const data = await res.json();
          console.log(data);
        }}
      >
        Obtener datos !
      </button>
      {/* Formulario */}
      <form onSubmit={handleSubmit} >
        <input
          type="name"
          placeholder="name"
          onChange={(e) => setName(e.target.value)}
        />
        <button>Save</button>
      </form>
    </div>
  );
}

export default App;
