import React from 'react';
import AlertForm from './components/alertForm';
import './App.css';

const AlertPage: React.FC = () => {
  return (
    <div className="App">
      <div className="form-container">
        <h1>Add Alerta</h1>
        <div className="form-content">
          <AlertForm />
        </div>
      </div>
    </div>
  );
};

export default AlertPage;