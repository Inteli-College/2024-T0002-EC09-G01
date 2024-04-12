import React, { useState } from 'react';
import Input from './form/input';
import Button from './form/button';
import ReactDOM from 'react-dom';
import App from '../sensorPage';
import AlertForm from './alertForm';

interface FormData {
    username: string;
    password: string;
}

const LoginForm: React.FC = () => {
  const [formData, setFormData] = useState<FormData>({ username: '',  password: ''});
  
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };
  
  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
        const url = 'http://localhost:8000'
        const response = await fetch(url+'/login', {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json' // Assuming you're sending JSON data
            },
            body: JSON.stringify(formData) // Convert formData to JSON string
        });
        if (response.ok) {
            const handleAlertas = () => {            
                ReactDOM.render(
                <React.StrictMode>
                    <header className="header"><Button type="submit" onClick={handleAlertas /* Função para voltar para a pagina de alertas*/ }>Voltar</Button></header>
                    <AlertForm />
                </React.StrictMode>,
                document.getElementById('root')
            );}

            ReactDOM.render(
                <React.StrictMode>
                    <header className="header"><Button type="submit" onClick={handleAlertas}>Alertas</Button></header>
                    <App />
                </React.StrictMode>,
                document.getElementById('root')
            );
            console.log('Data read successfully');
        } else {
            console.error('Failed to read data:', response.statusText);
        }
        } catch (error) {
            console.error('Error reading data:', error);
        }
        console.log('Form confirmed data:', formData);
    };
  
    return (
        <form onSubmit={handleSubmit}>
            <Input
                type="text"
                name="username"
                placeholder="username"
                value={formData.username}
                onChange={handleChange}
            />
            <Input
                type="text"
                name="password"
                placeholder="password"
                value={formData.password}
                onChange={handleChange}
            />
            <Button type="submit">Login</Button>
        </form>
    );
};

export default LoginForm;