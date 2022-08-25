import React from 'react'

function App() {
  return (
    <div>
      <h1>Hola mundo</h1>
      <button onClick={async()=> {
        const res = await fetch("http://localhost:4000/users")
        const data = await res.json()
        console.log(data)
      }}>
        Obtener datos  !
      </button>
    </div>
  )
}

export default App