import type {
  AuthResponse,
  LoginData,
  RegisterData,
  User,
} from '../../types/User';
import { api } from '../axiosConfig';

type UserApiResponse = {
  user_id: string;
  name: string;
  email: string;
  avatar?: string | null;
  role: string;
  subscriptionPlan: string;
};

const mapUserResponse = (data: UserApiResponse): User => ({
  email: data.email,
  role: data.role as User['role'],
  name: data.name,
  avatarUrl: data.avatar ?? undefined,
  subscriptionPlan: data.subscriptionPlan as User['subscriptionPlan'],
});

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

  getUserById: async (userId: string): Promise<User> => {
    const response = await api.get(`/api/v1/user/${userId}`);
    return mapUserResponse(response.data);
  },
};

export const getAuthHeader = () => {
  const token = localStorage.getItem('accessToken');
  return token ? { Authorization: `Bearer ${token}` } : {};
};
