import { useState } from 'react';
import { profileAPI } from '../api/endpoints/profile';
import type { ProfileData } from '../types/User';

export const useProfile = () => {
  const [profile, setProfile] = useState<ProfileData | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const getMyProfile = async () => {
    setIsLoading(true);
    setError(null);

    try {
      const response = await profileAPI.getMy();
      setProfile(response.profile ?? null);
      return true;
    } catch (err) {
      const message =
        typeof err === 'object' && err !== null && 'response' in err
          ? (err as { response?: { data?: { message?: string } } }).response
              ?.data?.message
          : undefined;
      setError(message || 'Failed to fetch profile');
      return false;
    } finally {
      setIsLoading(false);
    }
  };

  const clearError = () => {
    setError(null);
  };

  return {
    profile,
    getMyProfile,
    error,
    isLoading,
    clearError,
  };
};
