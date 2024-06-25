import React from "react";
import validationFunc from "../utils/validation";
import { checkNull } from "../utils/helper";
interface TextInputProps {
  label: string;
  value: string;
  placeholder?: string;
  onChange: React.ChangeEventHandler<HTMLInputElement>;
  validate?: string;
}
const TextInput: React.FC<TextInputProps> = ({
  label,
  value,
  placeholder,
  onChange,
  validate,
}) => {
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (checkNull(validate) !== "") {
      const func = validationFunc(validate);
      if (func !== undefined) {
        const result = func(e.target.value);
        console.log("result :: " + result);
        //setToast
      }
    }
    onChange(e);
  };
  return (
    <div>
      <label>{label}</label>
      <input
        type="text"
        value={value}
        onChange={handleChange}
        placeholder={placeholder}
      />
    </div>
  );
};

export default TextInput;
