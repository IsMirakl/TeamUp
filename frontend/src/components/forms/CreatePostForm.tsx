import { useState } from 'react';
import { usePost } from '../../hooks/usePost';
import type { PostCreate } from '../../types/Post';
import ButtonSubmit from '../ui/ButtonSubmit';
import InputField from '../ui/InputField';

type CreatePostFormProps = {
  onSubmit?: (data: PostCreate) => void;
};

const CreatePostForm = ({ onSubmit }: CreatePostFormProps) => {
  const { create, isLoading, error } = usePost();
  const [formData, setFormData] = useState<PostCreate>({
    title: '',
    description: '',
    tags: [],
    author: '',
  });
  const [tagsInput, setTagsInput] = useState('');

  const handleInputChange = (field: keyof PostCreate, value: string) => {
    setFormData(prev => ({ ...prev, [field]: value }));
  };

  const handleTagsChange = (value: string) => {
    setTagsInput(value);
    const tags = value
      .split(',')
      .map(tag => tag.trim())
      .filter(Boolean);
    setFormData(prev => ({ ...prev, tags }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const success = await create(formData);
    if (success) {
      onSubmit?.(formData);
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="space-y-5"
    >
      <InputField
        id="title"
        label="Заголовок"
        type="text"
        placeholder="Название поста"
        value={formData.title}
        onChange={e => handleInputChange('title', e.target.value)}
      />
      <InputField
        id="description"
        label="Описание"
        type="text"
        placeholder="Короткое описание"
        value={formData.description}
        onChange={e => handleInputChange('description', e.target.value)}
      />
      <InputField
        id="tags"
        label="Теги"
        type="text"
        placeholder="дизайн, фронтенд, ux"
        value={tagsInput}
        onChange={e => handleTagsChange(e.target.value)}
      />
      <InputField
        id="author"
        label="Автор"
        type="text"
        placeholder="Имя автора"
        value={formData.author}
        onChange={e => handleInputChange('author', e.target.value)}
      />

      {error ? <p className="text-sm text-rose-600">{error}</p> : null}
      {isLoading ? (
        <p className="text-xs text-slate-500">Сохраняем пост...</p>
      ) : null}

      <div className="pt-2">
        <ButtonSubmit
          id="create-post"
          type="submit"
          value="Опубликовать"
        />
      </div>
    </form>
  );
};

export default CreatePostForm;
