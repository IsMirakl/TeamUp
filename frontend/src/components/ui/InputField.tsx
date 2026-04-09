interface InputFieldProps {
  id: string;
  label: string;
  type: string;
  placeholder?: string;
  value?: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
  required?: boolean;
}

const InputField: React.FC<InputFieldProps> = ({
  id,
  label,
  type = 'text',
  placeholder,
  value,
  onChange,
  required,
}) => (
  <div className="space-y-2">
    <label
      htmlFor={id}
      className="block text-xs font-semibold tracking-[0.2em] text-slate-500 uppercase"
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
      className="h-12 w-full rounded-2xl border border-slate-200 bg-white/80 px-4 text-base text-slate-900 placeholder-slate-400 shadow-sm transition focus:border-sky-300 focus:ring-4 focus:ring-sky-100/70 focus:outline-none disabled:cursor-not-allowed disabled:bg-slate-100"
    />
  </div>
);

export default InputField;
