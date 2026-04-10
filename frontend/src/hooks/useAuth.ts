/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from 'react';
import { authAPI } from '../api/endpoints/auth';
import type { LoginData, RegisterData, User } from '../types/User';

export const useAuth = () => {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');
  const [isInitialized, setIsInitialized] = useState(false);

  const decodeToken = (token: string): { user_id?: string } | null => {
    try {
      const payload = token.split('.')[1];
      if (!payload) {
        return null;
      }
      const base64 = payload.replace(/-/g, '+').replace(/_/g, '/');
      const padded = base64.padEnd(
        base64.length + ((4 - (base64.length % 4)) % 4),
        '='
      );
      const decoded = JSON.parse(atob(padded));
      return decoded as { user_id?: string };
    } catch {
      return null;
    }
  };

  useEffect(() => {
    const initializeAuth = async () => {
      await checkAuth();
      setIsInitialized(true);
    };
    initializeAuth();
  }, []);

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

  const checkAuth = async (): Promise<boolean> => {
    const token = localStorage.getItem('accessToken');
    if (!token) {
      setUser(null);
      return false;
    }
    const tokenPayload = decodeToken(token);
    if (!tokenPayload?.user_id) {
      setUser(null);
      return false;
    }
    setIsLoading(true);

    try {
      const response = await authAPI.getUserById(tokenPayload.user_id);
      setUser(response);
      return true;
    } catch (error: any) {
      console.error('check auth error', error);
      setUser(null);
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
    checkAuth,
    clearError,
    isInitialized,
  };
};
