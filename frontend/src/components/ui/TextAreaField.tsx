interface TextAreaFieldProps {
  id: string;
  label: string;
  placeholder?: string;
  value?: string;
  onChange?: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
  required?: boolean;
  minLength?: number;
  maxLength?: number;
  rows?: number;
}

const TextAreaField: React.FC<TextAreaFieldProps> = ({
  id,
  label,
  placeholder,
  value,
  onChange,
  required,
  minLength,
  maxLength,
  rows = 6,
}) => (
  <div className="space-y-2">
    <label
      htmlFor={id}
      className="block text-xs font-semibold tracking-[0.2em] text-slate-500 uppercase"
    >
      {label}
    </label>
    <textarea
      id={id}
      required={required}
      minLength={minLength}
      maxLength={maxLength}
      placeholder={placeholder}
      value={value}
      onChange={onChange}
      rows={rows}
      className="w-full resize-y rounded-2xl border border-slate-200 bg-white/80 px-4 py-3 text-base text-slate-900 placeholder-slate-400 shadow-sm transition focus:border-sky-300 focus:ring-4 focus:ring-sky-100/70 focus:outline-none disabled:cursor-not-allowed disabled:bg-slate-100"
    />
  </div>
);

export default TextAreaField;
