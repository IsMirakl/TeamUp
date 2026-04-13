import { create } from 'zustand';
import { postAPI } from '../api/endpoints/post';
import type { Post, PostCreate, PostUpdate } from '../types/Post';

interface PostState {
  post: Post | null;
  loading: boolean;
  error: string | null;
  fetchPost: (id: string) => Promise<void>;
  createPost: (data: PostCreate) => Promise<boolean>;
  updatePost: (id: string, data: PostUpdate) => Promise<boolean>;
  clearError: () => void;
}

export const usePostStore = create<PostState>((set, get) => ({
  post: null,
  loading: false,
  error: null,

  fetchPost: async (id: string) => {
    const { post } = get();
    if (post?.id === id) return;

    set({ loading: true, error: null });
    try {
      const data = await postAPI.getPostById(id);
      set({ post: data, loading: false });
    } catch (err) {
      const message =
        typeof err === 'object' && err !== null && 'response' in err
          ? (err as { response?: { data?: { message?: string } } }).response
              ?.data?.message
          : undefined;
      set({ error: message || 'Failed to load post', loading: false });
    }
  },

  createPost: async (data: PostCreate) => {
    set({ loading: true, error: null });
    try {
      const createdPost = await postAPI.create(data);
      set({ post: createdPost, loading: false });
      return true;
    } catch (err) {
      const message =
        typeof err === 'object' && err !== null && 'response' in err
          ? (err as { response?: { data?: { message?: string } } }).response
              ?.data?.message
          : undefined;
      set({ error: message || 'Failed to create post', loading: false });
      return false;
    }
  },

  updatePost: async (id: string, data: PostUpdate) => {
    set({ loading: true, error: null });
    try {
      const updatedPost = await postAPI.update(id, data);
      set({ post: updatedPost, loading: false });
      return true;
    } catch (err) {
      const message =
        typeof err === 'object' && err !== null && 'response' in err
          ? (err as { response?: { data?: { message?: string } } }).response
              ?.data?.message
          : undefined;
      set({ error: message || 'Failed to update post', loading: false });
      return false;
    }
  },

  clearError: () => {
    set({ error: null });
  },
}));
