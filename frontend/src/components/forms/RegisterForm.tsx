import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';

import { useAuth } from '../../hooks/useAuth';
import ButtonSubmit from '../ui/ButtonSubmit';
import InputField from '../ui/InputField';

interface RegisterFormProps {
  onSubmit?: (data: { email: string; password: string }) => void;
}

const RegisterForm: React.FC<RegisterFormProps> = () => {
  const navigate = useNavigate();
  const { register } = useAuth();

  const [formData, setFormData] = useState({
    name: '',
    email: '',
    password: '',
    confirmPassword: '',
  });

  const handleInputChange = (field: string, value: string) => {
    setFormData(prev => ({ ...prev, [field]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (formData.password !== formData.confirmPassword) {
      return;
    }

    try {
      let success = false;
      // const {confirmPassword, ...registerData} = formData;
      success = await register(formData);

      if (success) {
        navigate('/home');
      }
    } catch (error) {
      console.error('Ошибка регистрации:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <InputField
        id="name"
        label="Имя"
        type="text"
        placeholder="Имя"
        value={formData.name}
        onChange={e => handleInputChange('name', e.target.value)}
      />

      <InputField
        id="email"
        label="Email"
        type="email"
        placeholder="Электронная почта"
        value={formData.email}
        onChange={e => handleInputChange('email', e.target.value)}
      />

      <InputField
        id="password"
        label="Пароль"
        type="password"
        placeholder="Введите пароль"
        value={formData.password}
        onChange={e => handleInputChange('password', e.target.value)}
      />

      <InputField
        id="confirmPassword"
        label="Повторите пароль"
        type="password"
        placeholder="Повторите пароль"
        value={formData.confirmPassword}
        onChange={e => handleInputChange('confirmPassword', e.target.value)}
      />

      <div className="mt-7 mb-6 flex justify-center">
        <ButtonSubmit
          id="login"
          type="submit"
          value="Зарегистрироваться"
        />
      </div>
      <Link
        className="flex justify-center text-blue-700"
        to={'/login'}
      >
        Уже зарегистрированы ?
      </Link>
    </form>
  );
};
export default RegisterForm;
