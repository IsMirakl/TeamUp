import { create } from 'zustand';
import { postAPI } from '../api/endpoints/post';
import type { Post, PostCreate, PostUpdate } from '../types/Post';

interface PostState {
  posts: Post[];
  post: Post | null;
  loading: boolean;
  error: string | null;
  fetchPosts: (limit?: number, offset?: number) => Promise<void>;
  fetchPost: (id: string) => Promise<void>;
  createPost: (data: PostCreate) => Promise<Post | null>;
  updatePost: (id: string, data: PostUpdate) => Promise<Post | null>;
  clearError: () => void;
}

const getErrorMessage = (err: unknown) => {
  if (typeof err === 'object' && err !== null && 'response' in err) {
    const response = (
      err as { response?: { data?: { message?: string; error?: string } } }
    ).response;
    return response?.data?.message || response?.data?.error;
  }
  return undefined;
};

export const usePostStore = create<PostState>((set, get) => ({
  posts: [],
  post: null,
  loading: false,
  error: null,

  fetchPosts: async (limit = 50, offset = 0) => {
    set({ loading: true, error: null });
    try {
      const posts = await postAPI.list(limit, offset);
      set({ posts, loading: false });
    } catch (err) {
      const message = getErrorMessage(err);
      set({ error: message || 'Failed to load posts', loading: false });
    }
  },

  fetchPost: async (id: string) => {
    const { post } = get();
    if (post?.id === id) return;

    set({ loading: true, error: null });
    try {
      const data = await postAPI.getPostById(id);
      set({ post: data, loading: false });
    } catch (err) {
      const message = getErrorMessage(err);
      set({ error: message || 'Failed to load post', loading: false });
    }
  },

  createPost: async (data: PostCreate) => {
    set({ loading: true, error: null });
    try {
      const createdPost = await postAPI.create(data);
      const hydratedPost = await postAPI
        .getPostById(createdPost.id)
        .catch(() => createdPost);
      const { posts } = get();
      set({ post: hydratedPost, posts: [hydratedPost, ...posts], loading: false });
      return hydratedPost;
    } catch (err) {
      const message = getErrorMessage(err);
      set({ error: message || 'Failed to create post', loading: false });
      return null;
    }
  },

  updatePost: async (id: string, data: PostUpdate) => {
    set({ loading: true, error: null });
    try {
      const updatedPost = await postAPI.update(id, data);
      const hydratedPost = await postAPI
        .getPostById(updatedPost.id)
        .catch(() => updatedPost);
      const { posts } = get();
      set({
        post: hydratedPost,
        posts: posts.map(p => (p.id === hydratedPost.id ? hydratedPost : p)),
        loading: false,
      });
      return hydratedPost;
    } catch (err) {
      const message = getErrorMessage(err);
      set({ error: message || 'Failed to update post', loading: false });
      return null;
    }
  },

  clearError: () => {
    set({ error: null });
  },
}));
