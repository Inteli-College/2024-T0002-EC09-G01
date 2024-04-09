import React, { useState } from 'react';
import Input from './form/input';
import Button from './form/button';

interface FormData {
    alerta: string;
    tipo: string;
    longitude: string;
    latitude: string;
}

const AlertForm: React.FC = () => {
  const [formData, setFormData] = useState<FormData>({ alerta: '',  tipo: '', longitude: '', latitude: '' });
  
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
        const response = await fetch(url+'/alerts', {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json' // Assuming you're sending JSON data
            },
            body: JSON.stringify(formData) // Convert formData to JSON string
        });
        if (response.ok) {
            console.log('Data sent successfully');
        } else {
            console.error('Failed to send data:', response.statusText);
        }
        } catch (error) {
            console.error('Error sending data:', error);
        }
        console.log('Form submitted with data:', formData);
    };
  
    return (
        <form onSubmit={handleSubmit}>
        <Input
            type="text"
            name="alerta"
            placeholder="alerta"
            value={formData.alerta}
            onChange={handleChange}
        />
        <Input
            type="text"
            name="tipo"
            placeholder="tipo"
            value={formData.tipo}
            onChange={handleChange}
        />
        <Input
            type="text"
            name="longitude"
            placeholder="longitude"
            value={formData.longitude}
            onChange={handleChange}
        />
        <Input
            type="text"
            name="latitude"
            placeholder="latitude"
            value={formData.latitude}
            onChange={handleChange}
        />
        <Button type="submit">Submit Alerta</Button>
        </form>
    );
};

export default AlertForm;