import { create } from 'zustand';
import { profileAPI } from '../api/endpoints/profile';
import type { ProfileData } from '../types/User';

type ProfileState = {
  profile: ProfileData | null;
  isLoading: boolean;
  error: string | null;
  getMyProfile: () => Promise<boolean>;
  clearError: () => void;
  clearProfile: () => void;
};

const getErrorMessage = (err: unknown) => {
  if (typeof err === 'object' && err !== null && 'response' in err) {
    const response = (err as { response?: { data?: { message?: string } } }).response;
    return response?.data?.message;
  }
  return undefined;
};

export const useProfileStore = create<ProfileState>(set => ({
  profile: null,
  isLoading: false,
  error: null,

  getMyProfile: async () => {
    set({ isLoading: true, error: null });
    try {
      const response = await profileAPI.getMy();
      set({
        profile: {
          name: response.name,
          email: response.email,
          avatar: response.avatar ?? undefined,
          role: response.role,
          subscriptionPlan: response.subscriptionPlan,
        },
        isLoading: false,
      });
      return true;
    } catch (err) {
      const message = getErrorMessage(err);
      set({ error: message || 'Failed to fetch profile', isLoading: false });
      return false;
    }
  },

  clearError: () => {
    set({ error: null });
  },

  clearProfile: () => {
    set({ profile: null });
  },
}));
