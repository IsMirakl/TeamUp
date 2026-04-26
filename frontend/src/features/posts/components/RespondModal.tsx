import { useEffect, useRef, useState } from 'react';
import { Link } from 'react-router-dom';
import { postAPI } from '../../../api/endpoints/post';
import type { Post } from '../../../types/Post';

type RespondModalProps = {
  open: boolean;
  post: Post | null;
  onClose: () => void;
  onSent?: (postId: string) => void;
};

const MAX_MESSAGE_LENGTH = 1000;

const getErrorMessage = (err: unknown) => {
  if (typeof err === 'object' && err !== null && 'response' in err) {
    const response = (
      err as { response?: { data?: { message?: string; error?: string } } }
    ).response;
    return response?.data?.message || response?.data?.error;
  }
  return undefined;
};

const RespondModal = ({ open, post, onClose, onSent }: RespondModalProps) => {
  const [message, setMessage] = useState('');
  const [telegram, setTelegram] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [isSending, setIsSending] = useState(false);
  const textareaRef = useRef<HTMLTextAreaElement | null>(null);

  const isAuthed = typeof window !== 'undefined' && Boolean(localStorage.getItem('accessToken'));

  useEffect(() => {
    if (!open) return;
    setError(null);
    setIsSending(false);
    setTelegram('');
    window.setTimeout(() => textareaRef.current?.focus(), 0);
  }, [open]);

  useEffect(() => {
    if (!open) return;
    const onKeyDown = (e: KeyboardEvent) => {
      if (e.key === 'Escape') onClose();
    };
    window.addEventListener('keydown', onKeyDown);
    return () => window.removeEventListener('keydown', onKeyDown);
  }, [open, onClose]);

  if (!open || !post) return null;

  const remaining = MAX_MESSAGE_LENGTH - message.length;

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!isAuthed) return;

    const telegramTrimmed = telegram.trim();
    if (!telegramTrimmed) {
      setError('Оставьте Telegram для связи');
      return;
    }

    const trimmed = message.trim();
    if (!trimmed) {
      setError('Напишите сообщение для автора');
      return;
    }
    if (trimmed.length > MAX_MESSAGE_LENGTH) {
      setError('Сообщение слишком длинное');
      return;
    }

    setIsSending(true);
    setError(null);
    try {
      await postAPI.respond(post.id, trimmed, telegramTrimmed);
      setMessage('');
      onSent?.(post.id);
      onClose();
    } catch (err) {
      const apiMessage = getErrorMessage(err);
      setError(apiMessage || 'Не удалось отправить отклик');
    } finally {
      setIsSending(false);
    }
  };

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center px-4"
      role="dialog"
      aria-modal="true"
    >
      <button
        type="button"
        className="absolute inset-0 bg-slate-950/40 backdrop-blur-sm"
        onClick={onClose}
        aria-label="Close"
      />

      <div className="relative w-full max-w-xl overflow-hidden rounded-3xl border border-slate-200/80 bg-white shadow-2xl shadow-slate-900/25 ring-1 ring-slate-900/5">
        <div className="border-b border-slate-200/80 bg-gradient-to-r from-sky-50 via-white to-amber-50 px-6 py-5">
          <p className="text-xs font-semibold tracking-[0.25em] text-slate-500 uppercase">
            Отклик на объявление
          </p>
          <h3 className="mt-2 text-lg font-semibold text-slate-900">{post.title}</h3>
          <p className="mt-1 text-sm text-slate-600">
            {post.author ? `Автор: ${post.author}` : 'Автор: —'}
          </p>
        </div>

        {!isAuthed ? (
          <div className="px-6 py-6">
            <p className="text-sm text-slate-700">
              Чтобы отправить отклик, нужно войти в аккаунт.
            </p>
            <div className="mt-5 flex flex-wrap items-center gap-3">
              <Link
                to="/login"
                className="rounded-full bg-slate-900 px-5 py-2 text-sm font-semibold text-white shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
              >
                Войти
              </Link>
              <button
                type="button"
                className="rounded-full border border-slate-200 bg-white px-5 py-2 text-sm font-semibold text-slate-700 shadow-sm transition hover:bg-slate-50"
                onClick={onClose}
              >
                Отмена
              </button>
            </div>
          </div>
        ) : (
          <form onSubmit={handleSubmit} className="px-6 py-6">
            <label
              htmlFor="response-telegram"
              className="block text-xs font-semibold tracking-[0.2em] text-slate-500 uppercase"
            >
              Telegram
            </label>
            <input
              id="response-telegram"
              value={telegram}
              onChange={e => setTelegram(e.target.value)}
              maxLength={128}
              placeholder="@username или https://t.me/username"
              className="mt-2 w-full rounded-2xl border border-slate-200 bg-white/90 px-4 py-3 text-base text-slate-900 placeholder-slate-400 shadow-sm transition focus:border-sky-300 focus:ring-4 focus:ring-sky-100/70 focus:outline-none disabled:cursor-not-allowed disabled:bg-slate-100"
              disabled={isSending}
              required
            />

            <label
              htmlFor="response-message"
              className="mt-5 block text-xs font-semibold tracking-[0.2em] text-slate-500 uppercase"
            >
              Сообщение
            </label>
            <textarea
              ref={textareaRef}
              id="response-message"
              value={message}
              onChange={e => setMessage(e.target.value)}
              maxLength={MAX_MESSAGE_LENGTH}
              rows={7}
              placeholder="Коротко расскажите, чем вы можете помочь, и оставьте контакты"
              className="mt-2 w-full resize-y rounded-2xl border border-slate-200 bg-white/90 px-4 py-3 text-base text-slate-900 placeholder-slate-400 shadow-sm transition focus:border-sky-300 focus:ring-4 focus:ring-sky-100/70 focus:outline-none disabled:cursor-not-allowed disabled:bg-slate-100"
              disabled={isSending}
              required
            />

            <div className="mt-2 flex items-center justify-between gap-3">
              <p className={`text-xs ${remaining < 0 ? 'text-rose-600' : 'text-slate-500'}`}>
                {remaining} символов осталось
              </p>
              {error ? <p className="text-xs font-semibold text-rose-600">{error}</p> : null}
            </div>

            <div className="mt-6 flex flex-wrap items-center justify-end gap-3">
              <button
                type="button"
                className="rounded-full border border-slate-200 bg-white px-5 py-2 text-sm font-semibold text-slate-700 shadow-sm transition hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-70"
                onClick={onClose}
                disabled={isSending}
              >
                Отмена
              </button>
              <button
                type="submit"
                className="rounded-full bg-slate-900 px-5 py-2 text-sm font-semibold text-white shadow-lg shadow-slate-900/20 transition hover:bg-slate-800 disabled:cursor-not-allowed disabled:opacity-70"
                disabled={isSending}
              >
                {isSending ? 'Отправляем...' : 'Отправить отклик'}
              </button>
            </div>
          </form>
        )}
      </div>
    </div>
  );
};

export default RespondModal;
