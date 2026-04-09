import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
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
      className="space-y-5"
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

      <label className="flex items-center gap-3 text-sm text-slate-600">
        <input
          type="checkbox"
          className="h-4 w-4 rounded border-slate-300 text-slate-900 focus:ring-2 focus:ring-sky-200"
        />
        Запомнить меня
      </label>

      <div className="pt-2">
        <ButtonSubmit
          id="login"
          type="submit"
          value="Войти"
        />
      </div>
    </form>
  );
};

export default LoginForm;
