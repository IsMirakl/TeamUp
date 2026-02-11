interface ButtonSubmitProps {
  id: string;
  type: string;
  value: string;
}

const ButtonSubmit: React.FC<ButtonSubmitProps> = ({ id, value }) => (
  <button
    id={id}
    type="submit"
    className="h-11 w-70 rounded-lg bg-blue-500 text-white transition-all duration-150 active:scale-95 active:bg-blue-600 active:shadow-md"
  >
    {value}
  </button>
);

export default ButtonSubmit;
