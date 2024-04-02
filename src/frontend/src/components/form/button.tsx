import React from 'react';

const Button: React.FC<React.ButtonHTMLAttributes<HTMLButtonElement>> = ({ children, ...rest }) => {
  return <button {...rest}>{children}</button>;
};

export default Button;
