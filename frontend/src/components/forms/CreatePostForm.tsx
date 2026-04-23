import { useState } from 'react';
import { usePost } from '../../hooks/usePost';
import type { Post, PostCreate } from '../../types/Post';
import ButtonSubmit from '../ui/ButtonSubmit';
import InputField from '../ui/InputField';
import TextAreaField from '../ui/TextAreaField';

type CreatePostFormProps = {
  onCreated?: (post: Post) => void;
};

const CreatePostForm = ({ onCreated }: CreatePostFormProps) => {
  const { create, isLoading, error } = usePost();
  const [formData, setFormData] = useState<PostCreate>({
    title: '',
    description: '',
    tags: [],
  });
  const [tagsInput, setTagsInput] = useState('');
  const [localError, setLocalError] = useState<string | null>(null);

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

    const trimmedTitle = formData.title.trim();
    const trimmedDescription = formData.description.trim();
    if (!trimmedTitle) {
      setLocalError('Укажите заголовок');
      return;
    }
    if (trimmedDescription.length < 100) {
      setLocalError('Описание должно быть не короче 100 символов');
      return;
    }
    if (trimmedDescription.length > 750) {
      setLocalError('Описание должно быть не длиннее 750 символов');
      return;
    }
    setLocalError(null);

    const created = await create({
      ...formData,
      title: trimmedTitle,
      description: trimmedDescription,
    });

    if (created) onCreated?.(created);
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
        required
      />
      <TextAreaField
        id="description"
        label="Описание"
        placeholder="Опишите задачу/идею (100-750 символов)"
        value={formData.description}
        onChange={e => handleInputChange('description', e.target.value)}
        required
        minLength={100}
        maxLength={750}
      />
      <InputField
        id="tags"
        label="Теги"
        type="text"
        placeholder="дизайн, фронтенд, ux"
        value={tagsInput}
        onChange={e => handleTagsChange(e.target.value)}
      />

      {localError ? <p className="text-sm text-rose-600">{localError}</p> : null}
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
