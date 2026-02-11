/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react';
import { authAPI } from '../api/endpoints/auth';
import type { LoginData, RegisterData, User } from '../types/dto/User';

export const useAuth = () => {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  const login = async (data: LoginData): Promise<boolean> => {
    setIsLoading(true);
    setError('');

    try {
      const response = await authAPI.login(data);
      setUser(response.user);
      return true;
    } catch (err: any) {
      setError(err.response?.data?.message || 'Login error');
      return false;
    } finally {
      setIsLoading(false);
    }
  };

  const register = async (data: RegisterData): Promise<boolean> => {
    setIsLoading(true);
    setError('');

    try {
      if (
        !data.email ||
        !data.name ||
        !data.password ||
        !data.confirmPassword
      ) {
        setError('Fill in all fields');
        return false;
      }

      const response = await authAPI.register(data);
      setUser(response.user);
      return true;
    } catch (err: any) {
      setError(err.response?.data?.message || 'Register error');
      return false;
    } finally {
      setIsLoading(false);
    }
  };

  const clearError = () => {
    setError('');
  };

  return {
    user,
    isLoading,
    error,
    login,
    register,
    clearError,
  };
};
