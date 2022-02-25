import React, { useState, useEffect }  from 'react';

const Home = () => {
const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [historial, setHistorial] = useState([]);
    useEffect(() => {
        fetch("http://localhost:3000/historial")
            .then(res => res.json())
            .then(
                (data) => {
                    setIsLoaded(true);
                    setHistorial(data);
                },
                (error) => {
                    setIsLoaded(true);
                    setError(error);
                }
            )
      }, [])
if (error) {
        return <div>Error: {error.message}</div>;
    } else if (!isLoaded) {
        return <div>Loading...</div>;
    } else {
        return (
            <div>
            <div>Historial Operaciones</div>
            <ul>
                {historial.map(operacion => (
                <li key={operacion.id}>
                    {operacion.valor_izq} {operacion.operador} {operacion.valor_der} {"="} {operacion.resultado} {"Hora: "} {operacion.CreatedAt}
                </li>
                ))}
            </ul>
            </div>
           
        );
    }
}
export default Home;