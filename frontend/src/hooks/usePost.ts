import { usePostStore } from '../stores/postsStore';

export const usePost = () => {
  const {
    post,
    fetchPost,
    loading,
    error,
    createPost,
    updatePost,
    clearError,
  } = usePostStore();

  return {
    post,
    fetchPost,
    isLoading: loading,
    error,
    create: createPost,
    update: updatePost,
    clearError,
  };
};
