import React from 'react';
import Form from './components/form/form';
import './App.css';

const App: React.FC = () => {
  return (
    <div className="App">
      <div className="form-container">
        <h1>Add Sensor</h1>
        <div className="form-content">
          <Form />
        </div>
      </div>
    </div>
  );
};

export default App;
