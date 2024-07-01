import React from "react";

interface LabelProps {
  label: string;
  className?: string;
}
const Label: React.FC<LabelProps> = ({ label, className }) => {
  return <label className={className}> {label}</label>;
};

export default Label;
