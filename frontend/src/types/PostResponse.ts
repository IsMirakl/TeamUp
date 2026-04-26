export interface PostResponse {
  responseId: string;
  postId: string;
  userId: string;
  message: string;
  status: string;
  createdAt: string;
  updatedAt: string;
  email: string;
  name: string;
  avatar: string | null;
}
