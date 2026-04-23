import { usePostStore } from '../stores/postsStore';

export const usePost = () => {
  const {
    posts,
    fetchPosts,
    post,
    fetchPost,
    loading,
    error,
    createPost,
    updatePost,
    clearError,
  } = usePostStore();

  return {
    posts,
    fetchPosts,
    post,
    fetchPost,
    isLoading: loading,
    error,
    create: createPost,
    update: updatePost,
    clearError,
  };
};
