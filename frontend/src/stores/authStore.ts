import { create } from 'zustand';
import { authAPI } from '../api/endpoints/auth';
import type { LoginData, RegisterData, User } from '../types/User';

type AuthState = {
  user: User | null;
  isLoading: boolean;
  error: string;
  isInitialized: boolean;
  initialize: () => Promise<void>;
  login: (data: LoginData) => Promise<boolean>;
  register: (data: RegisterData) => Promise<boolean>;
  checkAuth: () => Promise<boolean>;
  logout: () => void;
  clearError: () => void;
};

const decodeToken = (token: string): { user_id?: string } | null => {
  try {
    const payload = token.split('.')[1];
    if (!payload) return null;
    const base64 = payload.replace(/-/g, '+').replace(/_/g, '/');
    const padded = base64.padEnd(base64.length + ((4 - (base64.length % 4)) % 4), '=');
    const decoded = JSON.parse(atob(padded));
    return decoded as { user_id?: string };
  } catch {
    return null;
  }
};

const getErrorMessage = (err: unknown) => {
  if (typeof err === 'object' && err !== null && 'response' in err) {
    const response = (err as { response?: { data?: { message?: string } } }).response;
    return response?.data?.message;
  }
  return undefined;
};

export const useAuthStore = create<AuthState>((set, get) => ({
  user: null,
  isLoading: false,
  error: '',
  isInitialized: false,

  initialize: async () => {
    if (get().isInitialized) return;
    await get().checkAuth();
    set({ isInitialized: true });
  },

  login: async (data: LoginData) => {
    set({ isLoading: true, error: '' });
    try {
      await authAPI.login(data);
      const ok = await get().checkAuth();
      set({ isLoading: false });
      if (!ok) {
        set({ error: 'Failed to load user' });
        return false;
      }
      return true;
    } catch (err) {
      const message = getErrorMessage(err);
      set({ error: message || 'Login error', isLoading: false });
      return false;
    }
  },

  register: async (data: RegisterData) => {
    set({ isLoading: true, error: '' });
    try {
      if (!data.email || !data.name || !data.password || !data.confirmPassword) {
        set({ error: 'Fill in all fields', isLoading: false });
        return false;
      }
      await authAPI.register(data);
      const ok = await get().checkAuth();
      set({ isLoading: false });
      if (!ok) {
        set({ error: 'Failed to load user' });
        return false;
      }
      return true;
    } catch (err) {
      const message = getErrorMessage(err);
      set({ error: message || 'Register error', isLoading: false });
      return false;
    }
  },

  checkAuth: async () => {
    const token = localStorage.getItem('accessToken');
    if (!token) {
      set({ user: null });
      return false;
    }

    const tokenPayload = decodeToken(token);
    if (!tokenPayload?.user_id) {
      set({ user: null });
      return false;
    }

    set({ isLoading: true, error: '' });
    try {
      const user = await authAPI.getUserById(tokenPayload.user_id);
      set({ user, isLoading: false });
      return true;
    } catch {
      set({ user: null, isLoading: false });
      return false;
    }
  },

  logout: () => {
    localStorage.removeItem('accessToken');
    set({ user: null });
  },

  clearError: () => {
    set({ error: '' });
  },
}));
