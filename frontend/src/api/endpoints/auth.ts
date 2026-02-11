import type {
  AuthResponse,
  LoginData,
  RegisterData,
} from '../../types/dto/User';
import { api } from '../axiosConfig';

export const authAPI = {
  login: async (data: LoginData): Promise<AuthResponse> => {
    const response = await api.post('/api/v1/auth/login', data);
    localStorage.setItem('accessToken', response.data.accessToken);
    return response.data;
  },

  register: async (data: RegisterData): Promise<AuthResponse> => {
    const response = await api.post('/api/v1/auth/register', data);
    localStorage.setItem('accessToken', response.data.accessToken);
    return response.data;
  },
};

export const getAuthHeader = () => {
  const token = localStorage.getItem('accessToken');
  return token ? { Authorization: `Bearer ${token}` } : {};
};
