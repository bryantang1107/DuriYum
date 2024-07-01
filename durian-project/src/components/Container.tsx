import React, { ReactNode } from "react";

interface ContainerProps extends React.HTMLAttributes<HTMLDivElement> {
  className?: string;
  children: ReactNode;
}

const Container: React.FC<ContainerProps> = ({ children, className }) => {
  return <div className={`flex ${className}`}>{children}</div>;
};

export default Container;
