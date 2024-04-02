import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import Button from './components/form/button';
import AlertPage from './alertPage';

const handleSubmit = () => {
  ReactDOM.render(
    <React.StrictMode>
      <header className="header">
        <Button onClick={(handleSubmit)}> Voltar </Button>
      </header>
      <AlertPage />
    </React.StrictMode>,
    document.getElementById('root')  );
};

ReactDOM.render(
  <React.StrictMode>
    <header className="header">
      <Button onClick={(handleSubmit)}> Alertas </Button>
    </header>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);