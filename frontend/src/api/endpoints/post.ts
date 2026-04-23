import type { Post, PostCreate, PostUpdate } from '../../types/Post';
import { api } from '../axiosConfig';

const normalizePost = (raw: unknown): Post => {
  const obj: Record<string, unknown> =
    typeof raw === 'object' && raw !== null ? (raw as Record<string, unknown>) : {};

  const authorRaw =
    typeof obj.author === 'string'
      ? obj.author
      : typeof obj.author_name === 'string'
        ? obj.author_name
        : typeof obj.authorName === 'string'
          ? obj.authorName
          : undefined;

  return {
    id: String(obj.id ?? ''),
    title: String(obj.title ?? ''),
    description: String(obj.description ?? ''),
    tags: Array.isArray(obj.tags) ? (obj.tags as string[]) : [],
    author: authorRaw,
  };
};

export const postAPI = {
  list: async (limit = 50, offset = 0): Promise<Post[]> => {
    const response = await api.get('/api/v1/posts/post', {
      params: { limit, offset },
    });
    return Array.isArray(response.data)
      ? response.data.map(normalizePost)
      : [];
  },

  create: async (data: PostCreate): Promise<Post> => {
    const response = await api.post('/api/v1/posts/post', data);
    return normalizePost(response.data);
  },

  update: async (id: string, data: PostUpdate): Promise<Post> => {
    const response = await api.put(`/api/v1/posts/post/${id}`, data);
    return normalizePost(response.data);
  },

  getPostById: async (id: string): Promise<Post> => {
    const response = await api.get(`/api/v1/posts/post/${id}`);
    return normalizePost(response.data);
  },
};
