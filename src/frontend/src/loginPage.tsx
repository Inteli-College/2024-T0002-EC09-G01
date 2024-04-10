import React from 'react';
import LoginForm from './components/loginForm';
import './App.css';

const LoginPage: React.FC = () => {
  return (
    <div className="App">
        <div className="form-container">
            <h1>Login</h1>
            <div className="form-content">
                <LoginForm />
            </div>
        </div>
    </div>
  );
};

export default LoginPage;
