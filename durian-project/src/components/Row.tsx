import React, { ReactNode } from "react";
import styled from "styled-components";

interface RowProps {
  children: ReactNode;
  className?: string;
}

const StyledRow = styled.div`
  display: flex;
  flex-wrap: wrap;
`;
const Row: React.FC<RowProps> = ({ children, className }) => {
  return <StyledRow className={className}>{children}</StyledRow>;
};

export default Row;
