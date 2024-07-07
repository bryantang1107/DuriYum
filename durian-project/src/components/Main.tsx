import { ReactNode } from "react";

interface MainProps {
  children: ReactNode;
}

const Main: React.FC<MainProps> = ({ children }) => {
  return (
    <div className="flex-grow items-center justify-center">{children}</div>
  );
};

export default Main;
