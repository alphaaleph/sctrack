import React, {useState} from 'react';
import CarriersForm from './components/carrier';
import './App.css';

function App() {

    const [showCarriersForm, setShowCarriersForm] = useState(false);

  return (
    <div className="App">
        <h1 style={{ color: 'blue' }}>Tasks Demo</h1>
        <button onClick={() => setShowCarriersForm(!showCarriersForm)}>Carrier List</button>
        {showCarriersForm && <CarriersForm />}
    </div>
  );
}

export default App;