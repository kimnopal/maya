import React from "react";
interface IProps extends React.InputHTMLAttributes<HTMLInputElement> {
  labelText?: string;
  error?: string;
  children?: React.ReactNode;
}

const TextBox = React.forwardRef<HTMLInputElement, IProps>(
  (
    {
      className,
      children,
      labelText,
      type = "text",
      error,
      placeholder,
      ...props
    },
    ref
  ) => {
    return (
      <div className={"flex flex-col gap-1"}>
        {labelText && (
          <label
            className="block text-gray-600  mb-2 text-xs lg:text-sm xl:text-base"
            htmlFor="txt"
          >
            {labelText}
          </label>
        )}
        <div className="flex items-stretch">
          <input
            autoComplete="off"
            className={`border border-border disabled:border-slate-100 w-full block outline-none py-3 px-4 transition-all text-base  bg-white focus:shadow focus:shadow-blue-500
              ${error && "border-red-500 border  animate-shake"} ${
              children ? "rounded-r-md" : "rounded-xl"
            }`}
            {...props}
            placeholder={placeholder}
            ref={ref}
            type={type}
          ></input>

          <div className="flex">{children}</div>
        </div>
        {error && (
          <p className="text-red-600 text-right animate-shake">{error}</p>
        )}
      </div>
    );
  }
);

TextBox.displayName = "TextBox";
export default TextBox;
