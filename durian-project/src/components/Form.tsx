import React, { ReactNode } from "react";
interface FormProps extends React.HTMLProps<HTMLFormElement> {
  children: ReactNode;
}
const Form: React.FC<FormProps> = ({ children, ...props }) => {
  return <form {...props}>{children}</form>;
};

export default Form;
