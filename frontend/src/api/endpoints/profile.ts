import type { ProfileResponse } from '../../types/User';
import { api } from '../axiosConfig';

export const profileAPI = {
  getMy: async (): Promise<ProfileResponse> => {
    const response = await api.get('/api/v1/profile/me');
    return response.data;
  },
};
