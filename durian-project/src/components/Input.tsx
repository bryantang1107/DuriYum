import { forwardRef } from "react";
import validationFunc from "../utils/validation";
import { checkNull } from "../utils/helper";

interface InputProps {
  type: string;
  placeholder?: string;
  validate?: string;
  className?: string;
}
const Input = forwardRef<HTMLInputElement, InputProps>((props, ref) => {
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (checkNull(props.validate) !== "") {
      const func = validationFunc(props.validate);
      if (func !== undefined) {
        const result = func(e.target.value);
        console.log("result :: " + result);
        //setToast
      }
    }
  };

  return (
    <input
      ref={ref}
      type={props.type}
      placeholder={props.placeholder}
      onChange={handleChange}
      className={props.className && props.className}
    />
  );
});

export default Input;
