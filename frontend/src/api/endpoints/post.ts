import type { Post, PostCreate, PostUpdate } from '../../types/Post';
import { api } from '../axiosConfig';

export const postAPI = {
  create: async (data: PostCreate): Promise<Post> => {
    const response = await api.post('/api/v1/posts/post', data);
    return response.data;
  },

  update: async (id: string, data: PostUpdate): Promise<Post> => {
    const response = await api.patch(`/api/v1/posts/post/${id}`, data);
    return response.data;
  },

  getPostById: async (id: string): Promise<Post> => {
    const response = await api.get(`/api/v1/posts/post/${id}`);
    return response.data;
  },
};
