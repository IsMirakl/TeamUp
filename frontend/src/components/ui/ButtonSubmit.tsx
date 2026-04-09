interface ButtonSubmitProps {
  id: string;
  type: string;
  value: string;
}

const ButtonSubmit: React.FC<ButtonSubmitProps> = ({ id, value }) => (
  <button
    id={id}
    type="submit"
    className="h-12 w-full rounded-full bg-slate-900 text-sm font-semibold tracking-[0.2em] text-white uppercase shadow-lg shadow-slate-900/20 transition hover:bg-slate-800 active:scale-[0.98]"
  >
    {value}
  </button>
);

export default ButtonSubmit;
