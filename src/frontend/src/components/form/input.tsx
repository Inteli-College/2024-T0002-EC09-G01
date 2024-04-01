import React from 'react';

interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string;
}

const Input: React.FC<InputProps> = ({ label, ...rest }) => {
  return (
    <div>
      {label && <label>{label}</label>}
      <input {...rest} />
    </div>
  );
};

export default Input;
