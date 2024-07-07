import React, { ReactNode } from "react";

interface ContainerProps extends React.HTMLAttributes<HTMLDivElement> {
  className?: string;
  children: ReactNode;
  isCenter?: boolean;
}

const centerStyle = "h-full items-center justify-center";

const Container: React.FC<ContainerProps> = ({
  children,
  className,
  isCenter,
}) => {
  return (
    <div className={`flex ${className} ${isCenter ? centerStyle : ""}`}>
      {children}
    </div>
  );
};

export default Container;
