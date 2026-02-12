import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../../hooks/useAuth';
import ButtonSubmit from '../ui/ButtonSubmit';
import InputField from '../ui/InputField';

interface LoginFormProps {
  onSubmit?: (data: { email: string; password: string }) => void;
}

const LoginForm: React.FC<LoginFormProps> = () => {
  const navigate = useNavigate();
  const { login } = useAuth();

  const [formData, setFormData] = useState({
    email: '',
    password: '',
  });

  const handleInputChange = (field: string, value: string) => {
    setFormData(prev => ({ ...prev, [field]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      let success = false;
      success = await login(formData);

      if (success) {
        navigate('/home');
      }
    } catch (error) {
      console.error('Ошибка входа: ', error);
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
        value={formData.email}
        onChange={e => handleInputChange('email', e.target.value)}
      />
      <InputField
        id="password"
        label="Пароль"
        placeholder="Введите пароль"
        type="password"
        value={formData.password}
        onChange={e => handleInputChange('password', e.target.value)}
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
