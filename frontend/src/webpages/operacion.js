import { useState } from "react";

const Operacion = () => {
  const [valor_izq, setValor_Izq] = useState('');
  const [valor_der, setValor_Der] = useState('');
  const [operador, setOperador] = useState('');
  const [resultado, setResultado] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    var valorTemp = 0.0;
    console.log("valor operador "+operador);
    if(null == operador || operador == ""){
        setOperador("+");
    }
    if(operador == "+" ){
        valorTemp = Number(valor_izq) + Number(valor_der);
    }else if(operador == "-" ){
        valorTemp = Number(valor_izq) - Number(valor_der);
    }else if(operador == "/" ){
        valorTemp = Number(valor_izq) / Number(valor_der);
    }else if(operador == "*" ){
        valorTemp = Number(valor_izq) * Number(valor_der);
    }
    console.log('valorTemp: '+valorTemp);
    setResultado(String(valorTemp));
    const operacion = { valor_izq, valor_der, operador,"Resultado":String(valorTemp) };
    console.log(JSON.stringify(operacion))
    fetch('http://localhost:3000/operacion', {
      method: 'POST',
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(operacion)
    }).then(() => {
      console.log('nueva operacion agregada');
    })
  }

  return (
    <div className="operacion">
      <h2>Agregar una operacion</h2>
      <form onSubmit={handleSubmit}>
        <label>Valor 1:</label>
        <input 
          type="text" 
          required 
          value={valor_izq}
          onChange={(e) => setValor_Izq(e.target.value)}
        />
        <label>Operador:</label>
        <select
          value={operador}
          onChange={(e) => setOperador(e.target.value)}
        >
          <option value="+">+</option>
          <option value="-">-</option>
          <option value="*">*</option>
          <option value="/">/</option>
        </select>
        <label>Valor 2:</label>
        <input 
          type="text" 
          required 
          value={valor_der}
          onChange={(e) => setValor_Der(e.target.value)}
        />
        <button>Guardar Operacion</button>
      </form>
    </div>
  );
}
 
export default Operacion;