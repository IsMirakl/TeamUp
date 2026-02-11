import { useState } from 'react';
import { Link } from 'react-router-dom';
import ButtonSubmit from '../ui/ButtonSubmit';
import InputField from '../ui/InputField';

interface LoginFormProps {
  onSubmit?: (data: { email: string; password: string }) => void;
}

const LoginForm: React.FC<LoginFormProps> = ({ onSubmit }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (onSubmit) {
      onSubmit({ email, password });
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="space-y-4"
    >
      <InputField
        id="email"
        label="Email"
        placeholder="Электронная почта"
        type="email"
        value={email}
        onChange={e => setEmail(e.target.value)}
      />
      <InputField
        id="password"
        label="Пароль"
        placeholder="Введите пароль"
        type="password"
        value={password}
        onChange={e => setPassword(e.target.value)}
      />

      <div className="mt-5 ml-5 flex items-center gap-2">
        <input
          type="checkbox"
          className="h-5 w-5"
        />
        <p>Запомнить меня</p>
      </div>

      <div className="mt-7 mb-6 flex justify-center">
        <ButtonSubmit
          id="login"
          type="submit"
          value="Войти"
        />
      </div>
      <Link
        className="flex justify-center text-blue-700"
        to={'/register'}
      >
        Нету учетной записи ?
      </Link>
    </form>
  );
};

export default LoginForm;
