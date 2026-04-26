import { useEffect, useMemo } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import Header from '../components/Header';
import { useAuth } from '../hooks/useAuth';
import { usePost } from '../hooks/usePost';

const MyPostsPage = () => {
  const navigate = useNavigate();
  const { user, isLoading: authLoading } = useAuth();
  const { posts, fetchPosts, isLoading, error } = usePost();

  const isAuthenticated = !!localStorage.getItem('accessToken');

  useEffect(() => {
    if (!isAuthenticated) return;
    void fetchPosts(100, 0);
  }, [fetchPosts, isAuthenticated]);

  const myPosts = useMemo(() => {
    const name = user?.name?.trim();
    if (!name) return [];
    return posts.filter(p => (p.author ?? '').trim() === name);
  }, [posts, user?.name]);

  return (
    <section className="relative min-h-screen overflow-hidden bg-gradient-to-br from-amber-50 via-white to-sky-50">
      <div className="pointer-events-none absolute -top-48 left-10 h-96 w-96 rounded-full bg-amber-200/50 blur-3xl" />
      <div className="pointer-events-none absolute right-0 -bottom-40 h-96 w-96 rounded-full bg-sky-200/40 blur-3xl" />

      <div className="relative">
        <Header />
      </div>

      <main className="relative mx-auto w-full max-w-6xl px-6 py-10">
        <div className="mb-8">
          <p className="text-xs font-semibold tracking-[0.25em] text-slate-500 uppercase">TeamUP</p>
          <h1 className="mt-3 text-3xl font-semibold text-slate-900">Мои посты</h1>
          <p className="mt-2 max-w-2xl text-sm text-slate-600">
            Здесь вы можете открыть страницу откликов для каждого своего объявления.
          </p>
        </div>

        {!isAuthenticated ? (
          <div className="w-full max-w-2xl rounded-3xl border border-slate-200/80 bg-white/80 p-6 shadow-xl ring-1 shadow-slate-900/10 ring-slate-900/5 backdrop-blur">
            <p className="text-sm font-semibold text-slate-900">Нужно войти</p>
            <p className="mt-2 text-sm text-slate-600">
              Перейдите на страницу входа, чтобы увидеть свои посты.
            </p>
            <div className="mt-5 flex flex-wrap items-center gap-3">
              <Link
                className="rounded-full bg-slate-900 px-5 py-2 text-sm font-semibold tracking-[0.15em] text-white uppercase shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
                to={'/login'}
              >
                Войти
              </Link>
              <button
                type="button"
                className="rounded-full border border-slate-200 bg-white/70 px-5 py-2 text-sm font-semibold tracking-[0.15em] text-slate-700 uppercase transition hover:bg-white"
                onClick={() => navigate('/home')}
              >
                На главную
              </button>
            </div>
          </div>
        ) : (
          <div className="grid grid-cols-1 gap-6">
            {error ? (
              <div className="rounded-2xl border border-rose-200/70 bg-rose-50/70 p-4 text-sm text-rose-800">
                {error}
              </div>
            ) : null}

            {isLoading ? (
              <div className="rounded-3xl border border-slate-200/80 bg-white/80 p-6 shadow-xl ring-1 shadow-slate-900/10 ring-slate-900/5 backdrop-blur">
                <p className="text-sm text-slate-700">Загружаем посты…</p>
              </div>
            ) : null}

            {!isLoading && authLoading ? (
              <div className="rounded-3xl border border-slate-200/80 bg-white/80 p-6 shadow-xl ring-1 shadow-slate-900/10 ring-slate-900/5 backdrop-blur">
                <p className="text-sm text-slate-700">Загружаем аккаунт…</p>
              </div>
            ) : null}

            {!isLoading && !authLoading && myPosts.length === 0 ? (
              <div className="rounded-3xl border border-slate-200/80 bg-white/80 p-6 shadow-xl ring-1 shadow-slate-900/10 ring-slate-900/5 backdrop-blur">
                <p className="text-sm font-semibold text-slate-900">Пока нет постов</p>
                <p className="mt-2 text-sm text-slate-600">
                  Создайте объявление, чтобы начать получать отклики.
                </p>
                <div className="mt-5">
                  <Link
                    className="inline-flex rounded-full bg-slate-900 px-5 py-2 text-sm font-semibold text-white shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
                    to={'/posts/new'}
                  >
                    Создать пост
                  </Link>
                </div>
              </div>
            ) : null}

            {myPosts.map(post => (
              <article
                key={post.id}
                className="rounded-3xl border border-slate-200/80 bg-white/85 p-6 shadow-xl shadow-slate-900/10 ring-1 ring-slate-900/5 backdrop-blur"
              >
                <div className="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
                  <div className="min-w-0">
                    <h2 className="truncate text-lg font-semibold text-slate-900">{post.title}</h2>
                    <p className="mt-2 text-sm text-slate-700">{post.description}</p>
                  </div>
                  <div className="flex shrink-0 flex-wrap items-center gap-3">
                    <Link
                      className="rounded-full bg-slate-900 px-5 py-2 text-sm font-semibold text-white shadow-lg shadow-slate-900/20 transition hover:bg-slate-800"
                      to={`/posts/${post.id}/responses`}
                    >
                      Отклики
                    </Link>
                  </div>
                </div>
              </article>
            ))}
          </div>
        )}
      </main>
    </section>
  );
};

export default MyPostsPage;
