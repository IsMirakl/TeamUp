interface InputFieldProps {
  id: string;
  label: string;
  type: string;
  placeholder?: string;
  value?: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
  required?: boolean;
}

const InputFIeld: React.FC<InputFieldProps> = ({
  id,
  label,
  type = 'text',
  placeholder,
  value,
  onChange,
  required,
}) => (
  <div className="m-3">
    <label
      htmlFor={id}
      className="mb-2 block text-sm font-medium tracking-wide text-gray-700"
    >
      {label}
    </label>
    <input
      id={id}
      type={type}
      required={required}
      placeholder={placeholder}
      value={value}
      onChange={onChange}
      className="h-12 w-80 rounded-xl border-2 border-gray-200 bg-white/80 px-4 py-3 text-base font-medium text-gray-800 placeholder-gray-400 transition-all duration-200 ease-in-out invalid:border-red-400 hover:border-blue-400 hover:shadow-md hover:shadow-blue-100 focus:border-blue-500 focus:shadow-xl focus:ring-4 focus:shadow-blue-200/80 focus:ring-blue-100/50 focus:outline-none disabled:cursor-not-allowed disabled:bg-gray-100"
    />
  </div>
);

export default InputFIeld;
