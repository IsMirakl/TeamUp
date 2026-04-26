import type { Post, PostCreate, PostUpdate } from '../../types/Post';
import type { PostResponse } from '../../types/PostResponse';
import { api } from '../axiosConfig';

const normalizePost = (raw: unknown): Post => {
  const obj: Record<string, unknown> =
    typeof raw === 'object' && raw !== null ? (raw as Record<string, unknown>) : {};

  const authorIdRaw =
    typeof obj.author_id === 'string'
      ? obj.author_id
      : typeof obj.authorId === 'string'
        ? obj.authorId
        : undefined;

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
    authorId: authorIdRaw,
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

  getResponses: async (postId: string): Promise<PostResponse[]> => {
    const response = await api.get(`/api/v1/posts/post/${postId}/responses`);
    return Array.isArray(response.data)
      ? response.data.map(normalizePostResponse)
      : [];
  },

  respond: async (postId: string, message: string, telegram: string): Promise<unknown> => {
    const response = await api.post(`/api/v1/posts/post/${postId}/responses`, {
      message,
      telegram,
    });
    return response.data;
  },
};

const normalizePostResponse = (raw: unknown): PostResponse => {
  const obj: Record<string, unknown> =
    typeof raw === 'object' && raw !== null ? (raw as Record<string, unknown>) : {};

  const readString = (...keys: string[]) => {
    for (const key of keys) {
      const v = obj[key];
      if (typeof v === 'string') return v;
      if (typeof v === 'number') return String(v);
      if (typeof v === 'object' && v !== null && 'String' in v) {
        const s = (v as { String?: unknown }).String;
        if (typeof s === 'string') return s;
      }
    }
    return '';
  };

  const readNullableString = (...keys: string[]) => {
    for (const key of keys) {
      const v = obj[key];
      if (typeof v === 'string') return v;
      if (v === null) return null;
      if (typeof v === 'object' && v !== null && 'Valid' in v) {
        const cast = v as { String?: unknown; Valid?: unknown };
        if (cast.Valid === false) return null;
        if (typeof cast.String === 'string') return cast.String;
      }
    }
    return null;
  };

  return {
    responseId: readString('response_id', 'responseId', 'ResponseID', 'ResponseId'),
    postId: readString('post_id', 'postId', 'PostID', 'PostId'),
    userId: readString('user_id', 'userId', 'UserID', 'UserId'),
    message: readString('message', 'Message'),
    telegram: readNullableString('telegram', 'Telegram'),
    status: readString('status', 'Status'),
    createdAt: readString('created_at', 'createdAt', 'CreatedAt'),
    updatedAt: readString('updated_at', 'updatedAt', 'UpdatedAt'),
    email: readString('email', 'Email'),
    name: readString('name', 'Name'),
    avatar: readNullableString('avatar', 'Avatar'),
  };
};
