import { forwardRef, Ref } from "react";
import validationFunc from "../utils/validation";
import { checkNull } from "../utils/helper";

interface InputProps {
  type: "text" | "textarea" | "checkbox" | "file";
  placeholder?: string;
  validate?: string;
  className?: string;
  row?: number;
  col?: number;
}
const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ type, placeholder, validate, className, row, col }, ref) => {
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      if (checkNull(validate) !== "") {
        const func = validationFunc(validate);
        if (func !== undefined) {
          const result = func(e.target.value);
          console.log("result :: " + result);
          //setToast
        }
      }
    };

    if (type === "textarea") {
      return (
        <textarea
          ref={ref as Ref<HTMLTextAreaElement>}
          placeholder={placeholder}
          rows={row}
          cols={col}
        ></textarea>
      );
    }

    return (
      <input
        ref={ref as Ref<HTMLInputElement>}
        type={type}
        placeholder={placeholder}
        onChange={handleChange}
        className={className}
      ></input>
    );
  },
);

export default Input;
