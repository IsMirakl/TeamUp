import { useEffect, useMemo, useState } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import Header from '../components/Header';
import { postAPI } from '../api/endpoints/post';
import { useAuth } from '../hooks/useAuth';
import { usePost } from '../hooks/usePost';
import type { PostResponse } from '../types/PostResponse';

const formatDateTime = (value: string) => {
  const d = new Date(value);
  if (Number.isNaN(d.getTime())) return value;
  return d.toLocaleString();
};

const PostResponsesPage = () => {
  const navigate = useNavigate();
  const { id } = useParams();
  const { user, isLoading: authLoading } = useAuth();
  const { post, fetchPost, isLoading: postLoading, error: postError } = usePost();

  const [responses, setResponses] = useState<PostResponse[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const isAuthenticated = !!localStorage.getItem('accessToken');

  useEffect(() => {
    if (!id) return;
    void fetchPost(id);
  }, [fetchPost, id]);

  const isOwner = useMemo(() => {
    const me = user?.name?.trim();
    const author = post?.author?.trim();
    if (!me || !author) return false;
    return me === author;
  }, [post?.author, user?.name]);

  useEffect(() => {
    if (!id) return;
    if (!isAuthenticated) return;
    if (!post) return;
    if (!isOwner) return;

    setLoading(true);
    setError(null);
    void postAPI
      .getResponses(post.id)
      .then(data => setResponses(data))
      .catch(() => setError('Не удалось загрузить отклики'))
      .finally(() => setLoading(false));
  }, [id, isAuthenticated, isOwner, post]);

  return (
    <section className="relative min-h-screen overflow-hidden bg-gradient-to-br from-amber-50 via-white to-sky-50">
      <div className="pointer-events-none absolute -top-48 left-10 h-96 w-96 rounded-full bg-amber-200/50 blur-3xl" />
      <div className="pointer-events-none absolute right-0 -bottom-40 h-96 w-96 rounded-full bg-sky-200/40 blur-3xl" />

      <div className="relative">
        <Header />
      </div>

      <main className="relative mx-auto w-full max-w-6xl px-6 py-10">
        <div className="mb-8 flex flex-wrap items-center justify-between gap-4">
          <div>
            <p className="text-xs font-semibold tracking-[0.25em] text-slate-500 uppercase">TeamUP</p>
            <h1 className="mt-3 text-3xl font-semibold text-slate-900">Отклики</h1>
            <p className="mt-2 max-w-2xl text-sm text-slate-600">
              {post?.title ? `Пост: ${post.title}` : 'Загружаем пост…'}
            </p>
          </div>

          <div className="flex flex-wrap items-center gap-3">
            <Link
              to={'/my/posts'}
              className="rounded-full border border-slate-200 bg-white/70 px-5 py-2 text-sm font-semibold text-slate-700 shadow-sm transition hover:bg-white"
            >
              Мои посты
            </Link>
            <button
              type="button"
              onClick={() => navigate(-1)}
              className="rounded-full bg-slate-900 px-5 py-2 text-sm font-semibold text-white shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
            >
              Назад
            </button>
          </div>
        </div>

        {!id ? (
          <div className="rounded-3xl border border-slate-200/80 bg-white/80 p-6 shadow-xl ring-1 shadow-slate-900/10 ring-slate-900/5 backdrop-blur">
            <p className="text-sm font-semibold text-slate-900">Некорректный пост</p>
          </div>
        ) : !isAuthenticated ? (
          <div className="w-full max-w-2xl rounded-3xl border border-slate-200/80 bg-white/80 p-6 shadow-xl ring-1 shadow-slate-900/10 ring-slate-900/5 backdrop-blur">
            <p className="text-sm font-semibold text-slate-900">Нужно войти</p>
            <p className="mt-2 text-sm text-slate-600">
              Отклики доступны только автору поста.
            </p>
            <div className="mt-5 flex flex-wrap items-center gap-3">
              <Link
                className="rounded-full bg-slate-900 px-5 py-2 text-sm font-semibold tracking-[0.15em] text-white uppercase shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
                to={'/login'}
              >
                Войти
              </Link>
              <Link
                className="rounded-full border border-slate-200 bg-white/70 px-5 py-2 text-sm font-semibold tracking-[0.15em] text-slate-700 uppercase transition hover:bg-white"
                to={'/home'}
              >
                На главную
              </Link>
            </div>
          </div>
        ) : postError ? (
          <div className="rounded-2xl border border-rose-200/70 bg-rose-50/70 p-4 text-sm text-rose-800">
            {postError}
          </div>
        ) : authLoading || !user ? (
          <div className="rounded-3xl border border-slate-200/80 bg-white/80 p-6 shadow-xl ring-1 shadow-slate-900/10 ring-slate-900/5 backdrop-blur">
            <p className="text-sm text-slate-700">Загружаем аккаунт…</p>
          </div>
        ) : postLoading || !post ? (
          <div className="rounded-3xl border border-slate-200/80 bg-white/80 p-6 shadow-xl ring-1 shadow-slate-900/10 ring-slate-900/5 backdrop-blur">
            <p className="text-sm text-slate-700">Загружаем…</p>
          </div>
        ) : !isOwner ? (
          <div className="w-full max-w-3xl rounded-3xl border border-amber-200/70 bg-amber-50/70 p-6 text-sm text-amber-900 shadow-xl ring-1 ring-slate-900/5 backdrop-blur">
            <p className="font-semibold">Отклики видит только автор поста.</p>
            <p className="mt-2">
              Если это ваш пост, убедитесь, что вы вошли под правильным аккаунтом.
            </p>
          </div>
        ) : (
          <div className="space-y-6">
            {error ? (
              <div className="rounded-2xl border border-rose-200/70 bg-rose-50/70 p-4 text-sm text-rose-800">
                {error}
              </div>
            ) : null}

            {loading ? (
              <div className="rounded-3xl border border-slate-200/80 bg-white/80 p-6 shadow-xl ring-1 shadow-slate-900/10 ring-slate-900/5 backdrop-blur">
                <p className="text-sm text-slate-700">Загружаем отклики…</p>
              </div>
            ) : null}

            {!loading && responses.length === 0 ? (
              <div className="rounded-3xl border border-slate-200/80 bg-white/80 p-6 shadow-xl ring-1 shadow-slate-900/10 ring-slate-900/5 backdrop-blur">
                <p className="text-sm font-semibold text-slate-900">Пока нет откликов</p>
                <p className="mt-2 text-sm text-slate-600">
                  Когда кто-то откликнется, он появится здесь.
                </p>
              </div>
            ) : null}

            <div className="grid grid-cols-1 gap-4">
              {responses.map(r => (
                <article
                  key={r.responseId || `${r.email}-${r.createdAt}-${r.message}`}
                  className="rounded-3xl border border-slate-200/80 bg-white/85 p-6 shadow-xl shadow-slate-900/10 ring-1 ring-slate-900/5 backdrop-blur"
                >
                  <div className="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
                    <div className="flex min-w-0 items-start gap-4">
                      <div className="h-12 w-12 shrink-0 overflow-hidden rounded-2xl border border-slate-200/80 bg-gradient-to-br from-slate-900 to-slate-700 shadow-sm">
                        {r.avatar ? (
                          <img src={r.avatar} alt="Avatar" className="h-full w-full object-cover" />
                        ) : (
                          <div className="grid h-full w-full place-items-center text-sm font-semibold text-white">
                            {(r.name?.trim()?.[0] ?? 'U').toUpperCase()}
                          </div>
                        )}
                      </div>

                      <div className="min-w-0">
                        <p className="truncate text-sm font-semibold text-slate-900">
                          {r.name || 'Пользователь'}
                        </p>
                        <p className="truncate text-xs text-slate-600">{r.email || '—'}</p>
                        <p className="mt-2 whitespace-pre-wrap text-sm leading-6 text-slate-700">
                          {r.message}
                        </p>
                      </div>
                    </div>

                    <div className="flex shrink-0 flex-col items-start gap-2 sm:items-end">
                      <span className="rounded-full border border-slate-200 bg-white/70 px-3 py-1 text-xs font-semibold text-slate-700">
                        {r.status || '—'}
                      </span>
                      <span className="text-xs text-slate-500">
                        {r.createdAt ? formatDateTime(r.createdAt) : '—'}
                      </span>
                    </div>
                  </div>
                </article>
              ))}
            </div>
          </div>
        )}
      </main>
    </section>
  );
};

export default PostResponsesPage;
