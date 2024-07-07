import React from "react";
import styled from "styled-components";

interface ColProps {
  children: React.ReactNode;
  size?: number;
  className?: string;
}

const StyledCol = styled.div<ColProps>`
  flex: ${({ size }) => (size ? `0 0 ${(size / 12) * 100}%` : "1")};
  max-width: ${({ size }) => (size ? `${(size / 12) * 100}%` : "100%")};
  padding-left: 15px;
  padding-right: 15px;
`;

const Col: React.FC<ColProps> = ({ children, size, className }) => {
  return (
    <StyledCol size={size} className={className}>
      {children}
    </StyledCol>
  );
};

export default Col;
