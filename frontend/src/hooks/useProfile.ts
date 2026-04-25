import { useProfileStore } from '../stores/profileStore';

export const useProfile = () => {
  const profile = useProfileStore(state => state.profile);
  const getMyProfile = useProfileStore(state => state.getMyProfile);
  const error = useProfileStore(state => state.error);
  const isLoading = useProfileStore(state => state.isLoading);
  const clearError = useProfileStore(state => state.clearError);
  const clearProfile = useProfileStore(state => state.clearProfile);

  return {
    profile,
    getMyProfile,
    error,
    isLoading,
    clearError,
    clearProfile,
  };
};
