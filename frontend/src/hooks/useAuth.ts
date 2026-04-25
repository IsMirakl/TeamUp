import { useAuthStore } from '../stores/authStore';

export const   useAuth = () => {
  const user = useAuthStore(state => state.user);
  const isLoading = useAuthStore(state => state.isLoading);
  const error = useAuthStore(state => state.error);
  const isInitialized = useAuthStore(state => state.isInitialized);
  const initialize = useAuthStore(state => state.initialize);
  const login = useAuthStore(state => state.login);
  const register = useAuthStore(state => state.register);
  const checkAuth = useAuthStore(state => state.checkAuth);
  const logout = useAuthStore(state => state.logout);
  const clearError = useAuthStore(state => state.clearError);

  return {
    user,
    isLoading,
    error,
    isInitialized,
    initialize,
    login,
    register,
    checkAuth,
    logout,
    clearError,
  };
};
