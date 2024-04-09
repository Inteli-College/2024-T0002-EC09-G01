import React, { useState } from 'react';
import Input from './input';
import Button from './button';

interface FormData {
    sensor: string;
    tipo: string;
    longitude: string;
    latitude: string;
}

const Form: React.FC = () => {
  const [formData, setFormData] = useState<FormData>({ sensor: '',  tipo: '', longitude: '', latitude: '' });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const url:string = 'http://localhost:8000';

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await fetch(url+'/sensors', {
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
        name="sensor"
        placeholder="sensor"
        value={formData.sensor}
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
      <Button type="submit">Submit</Button>
    </form>
  );
};

export default Form;
