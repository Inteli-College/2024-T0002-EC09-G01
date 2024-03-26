import React, { useState } from 'react';
import Input from './input';
import Button from './button';

interface FormData {
    sensor: string;
    location: string;
}

const Form: React.FC = () => {
  const [formData, setFormData] = useState<FormData>({ sensor: '', location: '' });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
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
        name="location"
        placeholder="location"
        value={formData.location}
        onChange={handleChange}
      />
      <Button type="submit">Submit</Button>
    </form>
  );
};

export default Form;
