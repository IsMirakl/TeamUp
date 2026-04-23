export interface Post {
  id: string;
  title: string;
  description: string;
  tags: string[];
  author?: string;
}

export interface PostCreate {
  title: string;
  description: string;
  tags: string[];
}

export interface PostUpdate {
  title: string;
  description: string;
  tags: string[];
}
