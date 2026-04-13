import { usePostStore } from '../stores/postsStore';
import { postAPI } from '../api/endpoints/post';
import type { PostCreate, PostUpdate } from '../types/Post';

export const usePost = () => {
  const { post, fetchPost, loading, error } = usePostStore();

  const create = async (data: PostCreate): Promise<boolean> => {
    usePostStore.setState({ loading: true, error: null });
    try {
      const createdPost = await postAPI.create(data);
      usePostStore.setState({ post: createdPost, loading: false });
      return true;
    } catch (err) {
      const message =
        typeof err === 'object' && err !== null && 'response' in err
          ? (err as { response?: { data?: { message?: string } } }).response
              ?.data?.message
          : undefined;
      usePostStore.setState({
        error: message || 'Failed to create post',
        loading: false,
      });
      return false;
    }
  };

  const update = async (id: string, data: PostUpdate): Promise<boolean> => {
    usePostStore.setState({ loading: true, error: null });
    try {
      const updatedPost = await postAPI.update(id, data);
      usePostStore.setState({ post: updatedPost, loading: false });
      return true;
    } catch (err) {
      const message =
        typeof err === 'object' && err !== null && 'response' in err
          ? (err as { response?: { data?: { message?: string } } }).response
              ?.data?.message
          : undefined;
      usePostStore.setState({
        error: message || 'Failed to update post',
        loading: false,
      });
      return false;
    }
  };

  const clearError = () => {
    usePostStore.setState({ error: null });
  };

  return {
    post,
    fetchPost,
    isLoading: loading,
    error,
    create,
    update,
    clearError,
  };
};
