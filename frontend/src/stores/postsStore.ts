import { create } from 'zustand';
import { postAPI } from '../api/endpoints/post';
import type { Post } from '../types/Post';

interface PostState {
  post: Post | null;
  loading: boolean;
  error: string | null;
  fetchPost: (id: string) => Promise<void>;
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
      const msg = err instanceof Error ? err.message : 'Ошибка загрузки поста';
      set({ error: msg, loading: false });
    }
  },
}));
